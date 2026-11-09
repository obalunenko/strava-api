# strava-api

Go Strava API client generated from [API spec](https://developers.strava.com/swagger/swagger.json)

## Examples:

[Usage examples](examples/client_example_test.go)

The examples are live integration tests and require a Strava access token:

```shell
STRAVA_ACCESS_TOKEN="$(make strava-token)" go test -tags=integration ./examples
```

`make strava-token` refreshes the token when it expires. The first run asks for the Client ID and Client Secret
from [API settings](https://www.strava.com/settings/api), and for an existing refresh token. Press Enter instead of
the refresh token to authorize in the browser, this requires `localhost` as the Authorization Callback Domain.
The token needs the `activity:read` scope, which the refresh token from API settings does not have. To grant it, run
`go run ./internal/tools/strava-token -reauthorize`.
The credentials are stored unencrypted in `strava-api/token.json` in the user config directory, readable only by the
current user.

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
