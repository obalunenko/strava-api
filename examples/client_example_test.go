package examples

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	strava "github.com/obalunenko/strava-api/gen/strava-api-go"
)

func ExampleGetLoggedInAthlete() {
	token, ok := os.LookupEnv("STRAVA_ACCESS_TOKEN")
	if !ok {
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
}
