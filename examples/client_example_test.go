package examples

import (
	"context"
	"encoding/json"
	"log"
	"testing"

	"github.com/obalunenko/getenv"
	"github.com/stretchr/testify/require"

	"github.com/obalunenko/strava-api/client"
)

// helper function to print JSON
func printJSON(t testing.TB, v any) {
	indent, err := json.MarshalIndent(v, "", "  ")
	require.NoError(t, err)

	t.Log(string(indent))
}

// helper to get a valid API client
func getToken(t testing.TB) string {
	token := getenv.EnvOrDefault("STRAVA_ACCESS_TOKEN", "")
	if token == "" {
		t.Skip("STRAVA_ACCESS_TOKEN not set")
	}

	return token
}

func TestGetLoggedInAthlete(t *testing.T) {
	token := getToken(t)

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
	printJSON(t, athlete)
}

// Test Activites API
func TestGetLoggedInAthleteActivities(t *testing.T) {
	token := getToken(t)

	apiClient, err := client.NewAPIClient(token)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	activities, err := apiClient.Activities.GetLoggedInAthleteActivities(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Indent activities
	printJSON(t, activities)
}

// Test Gear API
func TestGetLoggedInAthleteGear(t *testing.T) {
	token := getToken(t)

	apiClient, err := client.NewAPIClient(token)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	athlete, err := apiClient.Athletes.GetLoggedInAthlete(ctx)
	require.NoError(t, err)

	require.NotNil(t, athlete.Bikes)

	gear, err := apiClient.Gears.GetGearById(ctx, athlete.Bikes[0].ID)
	require.NoError(t, err)

	// Indent gear
	printJSON(t, gear)
}
