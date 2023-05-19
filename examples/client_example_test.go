package examples

import (
	"context"
	"encoding/json"
	"log"
	"testing"

	"github.com/obalunenko/getenv"

	"github.com/obalunenko/strava-api/client"
)

func TestWrapper(t *testing.T) {
	if getenv.EnvOrDefault("CI", false) {
		t.Skip("Do not run on CI")
	}

	token := getenv.EnvOrDefault("STRAVA_ACCESS_TOKEN", "")
	if token == "" {
		log.Fatal("STRAVA_ACCESS_TOKEN not set")
	}

	apiClient, err := client.NewAPIClient(token)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	athlete, err := apiClient.Athletes.GetLoggedInAthlete(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Indent athlete
	indent, err := json.MarshalIndent(athlete, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(indent))
}
