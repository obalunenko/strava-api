package examples

import (
	"context"
	"fmt"
	strava "github.com/obalunenko/strava-api/gen/strava-api-go"
)

func ExampleNewClient() {
	cfg := strava.NewConfiguration()
	client := strava.NewAPIClient(cfg)

	ctx := context.Background()

	athlete, _, err := client.AthletesApi.GetLoggedInAthlete(ctx)
	if err != nil {
		return
	}

	fmt.Println(athlete)
}
