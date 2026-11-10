package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

type fakeTerminal struct {
	answers map[string]string
	asked   []string
}

func (f *fakeTerminal) ask(prompt string) (string, error) {
	f.asked = append(f.asked, prompt)

	return f.answers[prompt], nil
}

func (f *fakeTerminal) askSecret(prompt string) (string, error) {
	return f.ask(prompt)
}

type fakeStrava struct {
	mu    sync.Mutex
	forms []url.Values
}

func (s *fakeStrava) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mu.Lock()
	defer s.mu.Unlock()

	_ = r.ParseForm()
	s.forms = append(s.forms, r.PostForm)

	if r.PostForm.Get("refresh_token") == "revoked" {
		http.Error(w, `{"message":"Bad Request"}`, http.StatusBadRequest)

		return
	}

	n := len(s.forms)

	_ = json.NewEncoder(w).Encode(map[string]any{
		"access_token":  fmt.Sprintf("access-%d", n),
		"refresh_token": fmt.Sprintf("refresh-%d", n),
		"expires_at":    time.Now().Add(6 * time.Hour).Unix(),
	})
}

func (s *fakeStrava) calls() int {
	s.mu.Lock()
	defer s.mu.Unlock()

	return len(s.forms)
}

func (s *fakeStrava) last() url.Values {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.forms[len(s.forms)-1]
}

// browser answers the authorization request the way Strava redirects back.
func browser(t *testing.T, reply func(state string) url.Values) func(string) error {
	t.Helper()

	return func(target string) error {
		parsed, err := url.Parse(target)
		if err != nil {
			return err
		}

		query := parsed.Query()
		redirect := strings.Replace(query.Get("redirect_uri"), "localhost", "127.0.0.1", 1)

		go func() {
			resp, err := http.Get(redirect + "?" + reply(query.Get("state")).Encode())
			if err == nil {
				_ = resp.Body.Close()
			}
		}()

		return nil
	}
}

func approved(state string) url.Values {
	return url.Values{"state": {state}, "code": {"the-code"}, "scope": {"read,activity:read_all"}}
}

func newTestHelper(t *testing.T) (*helper, *fakeStrava, *fakeTerminal) {
	t.Helper()

	t.Setenv("STRAVA_CLIENT_ID", "")
	t.Setenv("STRAVA_CLIENT_SECRET", "")
	t.Setenv("STRAVA_REFRESH_TOKEN", "")

	strava := &fakeStrava{}
	server := httptest.NewServer(strava)
	t.Cleanup(server.Close)

	term := &fakeTerminal{answers: map[string]string{
		"Client ID: ":     "id",
		"Client Secret: ": "secret",
		"Refresh token (press Enter to authorize in the browser): ": "",
	}}

	return &helper{
		authorizeURL: server.URL + "/oauth/authorize",
		tokenURL:     server.URL + "/oauth/token",
		cachePath:    filepath.Join(t.TempDir(), "strava-api", "token.json"),
		httpClient:   server.Client(),
		term:         term,
		openBrowser:  browser(t, approved),
		log:          io.Discard,
	}, strava, term
}

func TestFirstRunAuthorizesInBrowser(t *testing.T) {
	h, strava, _ := newTestHelper(t)

	token, err := h.token(t.Context())
	require.NoError(t, err)
	require.Equal(t, "access-1", token)

	form := strava.last()
	require.Equal(t, "authorization_code", form.Get("grant_type"))
	require.Equal(t, "the-code", form.Get("code"))
	require.Equal(t, "id", form.Get("client_id"))
	require.Equal(t, "secret", form.Get("client_secret"))

	info, err := os.Stat(h.cachePath)
	require.NoError(t, err)
	require.Equal(t, os.FileMode(0o600), info.Mode().Perm())

	dir, err := os.Stat(filepath.Dir(h.cachePath))
	require.NoError(t, err)
	require.Equal(t, os.FileMode(0o700), dir.Mode().Perm())
}

func TestExistingRefreshTokenSkipsBrowser(t *testing.T) {
	h, strava, term := newTestHelper(t)
	term.answers["Refresh token (press Enter to authorize in the browser): "] = "pasted"
	h.openBrowser = func(string) error {
		t.Error("browser must not be opened")

		return nil
	}

	token, err := h.token(t.Context())
	require.NoError(t, err)
	require.Equal(t, "access-1", token)
	require.Equal(t, "refresh_token", strava.last().Get("grant_type"))
	require.Equal(t, "pasted", strava.last().Get("refresh_token"))
}

func TestCredentialsFromEnvNeedNoTerminal(t *testing.T) {
	h, strava, _ := newTestHelper(t)
	h.term = nil

	t.Setenv("STRAVA_CLIENT_ID", "env-id")
	t.Setenv("STRAVA_CLIENT_SECRET", "env-secret")
	t.Setenv("STRAVA_REFRESH_TOKEN", "env-refresh")

	token, err := h.token(t.Context())
	require.NoError(t, err)
	require.Equal(t, "access-1", token)
	require.Equal(t, "env-id", strava.last().Get("client_id"))
	require.Equal(t, "env-refresh", strava.last().Get("refresh_token"))
}

func TestFreshTokenIsReusedWithoutNetwork(t *testing.T) {
	h, strava, _ := newTestHelper(t)

	_, err := h.token(t.Context())
	require.NoError(t, err)

	calls := strava.calls()

	token, err := h.token(t.Context())
	require.NoError(t, err)
	require.Equal(t, "access-1", token)
	require.Equal(t, calls, strava.calls())
}

func expire(t *testing.T, h *helper, refreshToken string) {
	t.Helper()

	creds, err := h.load()
	require.NoError(t, err)

	creds.ExpiresAt = 0
	creds.RefreshToken = refreshToken

	require.NoError(t, h.save(creds))
}

func TestExpiredTokenIsRefreshedWithoutTerminal(t *testing.T) {
	h, strava, _ := newTestHelper(t)

	_, err := h.token(t.Context())
	require.NoError(t, err)

	expire(t, h, "refresh-1")

	h.term = nil

	token, err := h.token(t.Context())
	require.NoError(t, err)
	require.Equal(t, "access-2", token)
	require.Equal(t, "refresh_token", strava.last().Get("grant_type"))
	require.Equal(t, "refresh-1", strava.last().Get("refresh_token"))

	creds, err := h.load()
	require.NoError(t, err)
	require.Equal(t, "refresh-2", creds.RefreshToken)
}

func TestReauthorizeIgnoresStoredTokens(t *testing.T) {
	h, strava, term := newTestHelper(t)

	_, err := h.token(t.Context())
	require.NoError(t, err)

	t.Setenv("STRAVA_REFRESH_TOKEN", "stale")

	h.reauthorize = true
	term.asked = nil

	token, err := h.token(t.Context())
	require.NoError(t, err)
	require.Equal(t, "access-2", token)
	require.Equal(t, "authorization_code", strava.last().Get("grant_type"))
	require.Equal(t, "id", strava.last().Get("client_id"))
	require.Empty(t, term.asked)

	creds, err := h.load()
	require.NoError(t, err)
	require.Equal(t, "refresh-2", creds.RefreshToken)
}

func TestRevokedRefreshTokenFallsBackToBrowser(t *testing.T) {
	h, strava, term := newTestHelper(t)

	_, err := h.token(t.Context())
	require.NoError(t, err)

	expire(t, h, "revoked")

	term.asked = nil

	_, err = h.token(t.Context())
	require.NoError(t, err)
	require.Equal(t, "authorization_code", strava.last().Get("grant_type"))
	require.Equal(t, []string{"Refresh token (press Enter to authorize in the browser): "}, term.asked)
}

func TestRevokedRefreshTokenWithoutTerminalFails(t *testing.T) {
	h, _, _ := newTestHelper(t)

	_, err := h.token(t.Context())
	require.NoError(t, err)

	expire(t, h, "revoked")

	h.term = nil

	_, err = h.token(t.Context())
	require.ErrorIs(t, err, errNoTerminal)
}

func TestNoTerminalAndNoCredentialsFails(t *testing.T) {
	h, strava, _ := newTestHelper(t)
	h.term = nil

	_, err := h.token(t.Context())
	require.ErrorIs(t, err, errNoTerminal)
	require.Zero(t, strava.calls())
}

func TestRejectedCallbacksRequestNoToken(t *testing.T) {
	tests := map[string]struct {
		reply func(state string) url.Values
		want  string
	}{
		"forged state": {
			reply: func(string) url.Values {
				return url.Values{"state": {"forged"}, "code": {"x"}, "scope": {"activity:read"}}
			},
			want: "state mismatch",
		},
		"access denied": {
			reply: func(state string) url.Values {
				return url.Values{"state": {state}, "error": {"access_denied"}}
			},
			want: "access_denied",
		},
		"no code": {
			reply: func(state string) url.Values {
				return url.Values{"state": {state}, "scope": {"activity:read"}}
			},
			want: "no authorization code",
		},
		"activities not granted": {
			reply: func(state string) url.Values {
				return url.Values{"state": {state}, "code": {"x"}, "scope": {"read"}}
			},
			want: "reading activities",
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			h, strava, _ := newTestHelper(t)
			h.openBrowser = browser(t, tt.reply)

			_, err := h.token(t.Context())
			require.ErrorContains(t, err, tt.want)
			require.Zero(t, strava.calls())
			require.NoFileExists(t, h.cachePath)
		})
	}
}
