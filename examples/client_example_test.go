package examples

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"testing"

	"github.com/obalunenko/getenv"
	"github.com/stretchr/testify/assert"

	strava "github.com/obalunenko/strava-api/gen/strava-api-go"
)

func TestGetLoggedInAthlete(t *testing.T) {
	if getenv.EnvOrDefault("CI", false) {
		t.Skip("Do not run on CI")
	}

	token := getenv.EnvOrDefault("STRAVA_ACCESS_TOKEN", "")
	if token == "" {
		log.Fatal("STRAVA_ACCESS_TOKEN not set")
	}

	// Authentication is provided via context values.
	ctx := context.WithValue(context.Background(), strava.ContextAccessToken, token)

	client := strava.NewAPIClient(strava.NewConfiguration())

	athlete, resp, err := client.AthletesApi.GetLoggedInAthlete(ctx)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatal(fmt.Errorf("not ok: %s", http.StatusText(resp.StatusCode)))
	}

	fmt.Println(athlete)

	// not empty
	assert.NotEqual(t, strava.DetailedAthlete{}, athlete)
}
