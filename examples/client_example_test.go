package examples

import (
	"encoding/json"
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

func makeClient(t testing.TB) *client.APIClient {
	t.Helper()

	token := getToken(t)

	apiClient, err := client.NewAPIClient(token)
	require.NoError(t, err)

	return apiClient
}

func TestGetLoggedInAthlete(t *testing.T) {
	apiClient := makeClient(t)

	athlete, err := apiClient.Athletes.GetLoggedInAthlete(t.Context())
	require.NoError(t, err)

	// Indent athlete
	printJSON(t, athlete)
}

// Test Activites API
func TestGetLoggedInAthleteActivities(t *testing.T) {
	apiClient := makeClient(t)

	activities, err := apiClient.Activities.GetLoggedInAthleteActivities(t.Context())
	require.NoError(t, err)

	// Indent activities
	printJSON(t, activities)
}

// Test Gear API
func TestGetLoggedInAthleteGear(t *testing.T) {
	apiClient := makeClient(t)

	ctx := t.Context()

	athlete, err := apiClient.Athletes.GetLoggedInAthlete(ctx)
	require.NoError(t, err)

	require.NotNil(t, athlete.Bikes)

	gear, err := apiClient.Gears.GetGearById(ctx, athlete.Bikes[0].ID)
	require.NoError(t, err)

	// Indent gear
	printJSON(t, gear)
}

func TestGetRoutesByAthleteID(t *testing.T) {
	apiClient := makeClient(t)

	ctx := t.Context()

	athlete, err := apiClient.Athletes.GetLoggedInAthlete(ctx)
	require.NoError(t, err)

	id := athlete.ID
	require.NotZero(t, id)

	routes, err := apiClient.Routes.GetRoutesByAthleteId(ctx, id)
	require.NoError(t, err)

	printJSON(t, routes)
}
