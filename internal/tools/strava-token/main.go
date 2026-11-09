// Command strava-token prints a Strava access token for the integration tests.
//
// The credentials are kept in the user config directory and the token is refreshed
// when it expires. The first run asks for the application credentials and either
// takes an existing refresh token or authorizes in the browser.
package main

import (
	"bufio"
	"context"
	"crypto/rand"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"net"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"runtime"
	"slices"
	"strings"
	"time"

	"github.com/obalunenko/getenv"
)

const (
	authorizeURL = "https://www.strava.com/oauth/authorize"
	tokenURL     = "https://www.strava.com/oauth/token"

	// scope is read-only on purpose: the integration tests never change the account.
	scope = "read,read_all,profile:read_all,activity:read_all"

	callbackPath     = "/exchange_token"
	authorizeTimeout = 5 * time.Minute
	expiryMargin     = 2 * time.Minute
	requestTimeout   = 30 * time.Second
)

var errNoTerminal = errors.New("no terminal to ask for credentials: run it once from a terminal")

type credentials struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresAt    int64  `json:"expires_at"`
}

func (c credentials) fresh(now time.Time) bool {
	return c.AccessToken != "" && time.Unix(c.ExpiresAt, 0).After(now.Add(expiryMargin))
}

func (c credentials) canRefresh() bool {
	return c.ClientID != "" && c.ClientSecret != "" && c.RefreshToken != ""
}

func (c *credentials) fillFromEnv() {
	fill := func(field *string, key string) {
		if *field == "" {
			*field = getenv.EnvOrDefault(key, "")
		}
	}

	fill(&c.ClientID, "STRAVA_CLIENT_ID")
	fill(&c.ClientSecret, "STRAVA_CLIENT_SECRET")
	fill(&c.RefreshToken, "STRAVA_REFRESH_TOKEN")
}

type terminal interface {
	ask(prompt string) (string, error)
	askSecret(prompt string) (string, error)
}

type helper struct {
	authorizeURL string
	tokenURL     string
	cachePath    string
	httpClient   *http.Client
	term         terminal
	openBrowser  func(target string) error
	log          io.Writer
	reauthorize  bool
}

func (h *helper) token(ctx context.Context) (string, error) {
	creds, err := h.load()
	if err != nil {
		return "", err
	}

	creds.fillFromEnv()

	if !h.reauthorize {
		if creds.fresh(time.Now()) {
			return creds.AccessToken, nil
		}

		if creds.canRefresh() {
			err = h.refresh(ctx, &creds)
			if err == nil {
				return creds.AccessToken, nil
			}

			_, _ = fmt.Fprintf(h.log, "refresh failed: %v\n", err)
		}
	}

	if h.term == nil {
		return "", errNoTerminal
	}

	if err = h.askApplication(&creds); err != nil {
		return "", err
	}

	if !h.reauthorize {
		creds.RefreshToken, err = h.term.askSecret("Refresh token (press Enter to authorize in the browser): ")
		if err != nil {
			return "", err
		}

		if creds.RefreshToken != "" {
			if err = h.refresh(ctx, &creds); err != nil {
				return "", err
			}

			return creds.AccessToken, nil
		}
	}

	code, err := h.authorize(ctx, creds.ClientID)
	if err != nil {
		return "", err
	}

	grant := url.Values{"grant_type": {"authorization_code"}, "code": {code}}
	if err = h.exchange(ctx, &creds, grant); err != nil {
		return "", err
	}

	return creds.AccessToken, nil
}

func (h *helper) askApplication(creds *credentials) error {
	var err error

	if creds.ClientID == "" {
		if creds.ClientID, err = h.term.ask("Client ID: "); err != nil {
			return err
		}
	}

	if creds.ClientSecret == "" {
		if creds.ClientSecret, err = h.term.askSecret("Client Secret: "); err != nil {
			return err
		}
	}

	if creds.ClientID == "" || creds.ClientSecret == "" {
		return errors.New("client ID and client secret are required")
	}

	return nil
}

func (h *helper) refresh(ctx context.Context, creds *credentials) error {
	return h.exchange(ctx, creds, url.Values{"grant_type": {"refresh_token"}, "refresh_token": {creds.RefreshToken}})
}

func (h *helper) exchange(ctx context.Context, creds *credentials, form url.Values) error {
	form.Set("client_id", creds.ClientID)
	form.Set("client_secret", creds.ClientSecret)

	ctx, cancel := context.WithTimeout(ctx, requestTimeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, h.tokenURL, strings.NewReader(form.Encode()))
	if err != nil {
		return fmt.Errorf("build token request: %w", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := h.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("request token: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(io.LimitReader(resp.Body, 300))

		return fmt.Errorf("strava responded %d: %s", resp.StatusCode, body)
	}

	var token struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
		ExpiresAt    int64  `json:"expires_at"`
	}

	if err = json.NewDecoder(resp.Body).Decode(&token); err != nil {
		return fmt.Errorf("decode token response: %w", err)
	}

	if token.AccessToken == "" || token.RefreshToken == "" {
		return errors.New("token response has no tokens")
	}

	creds.AccessToken = token.AccessToken
	creds.RefreshToken = token.RefreshToken
	creds.ExpiresAt = token.ExpiresAt

	return h.save(*creds)
}

func (h *helper) authorize(ctx context.Context, clientID string) (string, error) {
	var lc net.ListenConfig

	listener, err := lc.Listen(ctx, "tcp", "127.0.0.1:0")
	if err != nil {
		return "", fmt.Errorf("listen for the callback: %w", err)
	}

	callbacks := make(chan url.Values, 1)

	mux := http.NewServeMux()
	mux.HandleFunc(callbackPath, func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintln(w, "You can close this tab and return to the terminal.")

		select {
		case callbacks <- r.URL.Query():
		default:
		}
	})

	server := &http.Server{Handler: mux, ReadHeaderTimeout: 5 * time.Second}

	go func() { _ = server.Serve(listener) }()

	defer func() {
		// Graceful, so that the browser still gets the response.
		shutdownCtx, cancel := context.WithTimeout(context.WithoutCancel(ctx), 2*time.Second)
		defer cancel()

		_ = server.Shutdown(shutdownCtx)
	}()

	port := listener.Addr().(*net.TCPAddr).Port
	state := rand.Text()

	target := h.authorizeURL + "?" + url.Values{
		"client_id":       {clientID},
		"response_type":   {"code"},
		"redirect_uri":    {fmt.Sprintf("http://localhost:%d%s", port, callbackPath)},
		"approval_prompt": {"force"},
		"scope":           {scope},
		"state":           {state},
	}.Encode()

	_, _ = fmt.Fprintf(h.log, "Confirm access in the browser. If it did not open, visit:\n%s\n", target)

	if err = h.openBrowser(target); err != nil {
		_, _ = fmt.Fprintf(h.log, "open browser: %v\n", err)
	}

	ctx, cancel := context.WithTimeout(ctx, authorizeTimeout)
	defer cancel()

	select {
	case query := <-callbacks:
		return checkCallback(query, state)
	case <-ctx.Done():
		return "", fmt.Errorf("wait for the browser confirmation: %w", ctx.Err())
	}
}

func checkCallback(query url.Values, state string) (string, error) {
	if query.Get("state") != state {
		return "", errors.New("state mismatch, callback rejected")
	}

	if reason := query.Get("error"); reason != "" {
		return "", fmt.Errorf("access was not granted: %s", reason)
	}

	code := query.Get("code")
	if code == "" {
		return "", errors.New("callback has no authorization code")
	}

	granted := strings.Split(query.Get("scope"), ",")
	if !slices.Contains(granted, "activity:read") && !slices.Contains(granted, "activity:read_all") {
		return "", errors.New("reading activities was not granted, keep that checkbox enabled")
	}

	return code, nil
}

func (h *helper) load() (credentials, error) {
	var creds credentials

	data, err := os.ReadFile(h.cachePath)
	if errors.Is(err, fs.ErrNotExist) {
		return creds, nil
	}

	if err != nil {
		return creds, fmt.Errorf("read credentials: %w", err)
	}

	if err = json.Unmarshal(data, &creds); err != nil {
		return creds, fmt.Errorf("parse %s: %w", h.cachePath, err)
	}

	return creds, nil
}

func (h *helper) save(creds credentials) error {
	if err := os.MkdirAll(filepath.Dir(h.cachePath), 0o700); err != nil {
		return fmt.Errorf("create credentials directory: %w", err)
	}

	data, err := json.Marshal(creds)
	if err != nil {
		return fmt.Errorf("encode credentials: %w", err)
	}

	tmp := h.cachePath + ".tmp"

	if err = os.WriteFile(tmp, data, 0o600); err != nil {
		return fmt.Errorf("write credentials: %w", err)
	}

	if err = os.Rename(tmp, h.cachePath); err != nil {
		return fmt.Errorf("replace credentials: %w", err)
	}

	return nil
}

type tty struct {
	file   *os.File
	reader *bufio.Reader
}

// openTerminal uses the controlling terminal rather than stdin and stdout,
// so that the token can be captured with $(...).
func openTerminal() terminal {
	file, err := os.OpenFile("/dev/tty", os.O_RDWR, 0)
	if err != nil {
		return nil
	}

	return &tty{file: file, reader: bufio.NewReader(file)}
}

func (t *tty) ask(prompt string) (string, error) {
	if _, err := fmt.Fprint(t.file, prompt); err != nil {
		return "", fmt.Errorf("write prompt: %w", err)
	}

	line, err := t.reader.ReadString('\n')
	if err != nil && !errors.Is(err, io.EOF) {
		return "", fmt.Errorf("read answer: %w", err)
	}

	return strings.TrimSpace(line), nil
}

func (t *tty) askSecret(prompt string) (string, error) {
	if err := t.stty("-echo"); err != nil {
		return "", fmt.Errorf("hide input: %w", err)
	}

	interrupted := make(chan os.Signal, 1)
	signal.Notify(interrupted, os.Interrupt)

	go func() {
		if _, ok := <-interrupted; ok {
			_ = t.stty("echo")

			os.Exit(130)
		}
	}()

	defer func() {
		signal.Stop(interrupted)
		close(interrupted)

		_ = t.stty("echo")
		_, _ = fmt.Fprintln(t.file)
	}()

	return t.ask(prompt)
}

func (t *tty) stty(arg string) error {
	cmd := exec.Command("stty", arg)
	cmd.Stdin = t.file

	return cmd.Run()
}

func openBrowser(target string) error {
	name := "xdg-open"
	if runtime.GOOS == "darwin" {
		name = "open"
	}

	return exec.Command(name, target).Start()
}

func run(ctx context.Context) error {
	reauthorize := flag.Bool("reauthorize", false, "ignore stored tokens and authorize in the browser, e.g. to grant missing scopes")
	flag.Parse()

	configDir, err := os.UserConfigDir()
	if err != nil {
		return fmt.Errorf("find config directory: %w", err)
	}

	h := &helper{
		authorizeURL: authorizeURL,
		tokenURL:     tokenURL,
		cachePath:    filepath.Join(configDir, "strava-api", "token.json"),
		httpClient:   http.DefaultClient,
		term:         openTerminal(),
		openBrowser:  openBrowser,
		log:          os.Stderr,
		reauthorize:  *reauthorize,
	}

	token, err := h.token(ctx)
	if err != nil {
		return err
	}

	fmt.Println(token)

	return nil
}

func main() {
	if err := run(context.Background()); err != nil {
		fmt.Fprintln(os.Stderr, "strava-token:", err)

		os.Exit(1)
	}
}
