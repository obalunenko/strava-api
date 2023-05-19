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
	"time"

	strava "github.com/obalunenko/strava-api/client"
)

func main() {
	key := "STRAVA_ACCESS_TOKEN"
	
	token := os.Getenv(key)
	if token == "" {
		log.Fatalf("%q not set", key)
	}

	apiClient, err := strava.NewAPIClient(token)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	athlete, err := apiClient.Athletes.GetLoggedInAthlete(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(athlete)
}

```

