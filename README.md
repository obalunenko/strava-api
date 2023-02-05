# strava-api

Go Strava API client generated from [API spec](https://developers.strava.com/swagger/swagger.json)

## Examples:

[Usage examples](examples/client_example_test.go)

### Usage

```go
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	strava "github.com/obalunenko/strava-api/gen/strava-api-go"
)

func main() {
	token := os.Getenv("STRAVA_ACCESS_TOKEN")
	if token == "" {
		log.Fatal("STRAVA_ACCESS_TOKEN not set")
	}

	// Authentication is provided via context values.
	ctx := context.WithValue(context.Background(), strava.ContextAccessToken, token)

	client := strava.NewAPIClient(strava.NewConfiguration())

	athlete, _, err := client.AthletesApi.GetLoggedInAthlete(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(athlete)
}

```

