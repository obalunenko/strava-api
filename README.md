# strava-api

Go Strava API client generated from [API spec](https://developers.strava.com/swagger/swagger.json)

## Examples:

[Usage examples](examples/client_example_test.go)

The examples are live integration tests and require a Strava access token:

```shell
make test-integration
```

`make strava-token` refreshes the token when it expires. The first run asks for the Client ID and Client Secret
from [API settings](https://www.strava.com/settings/api), and for an existing refresh token. Press Enter instead of
the refresh token to authorize in the browser, this requires `localhost` as the Authorization Callback Domain.
The token needs the `activity:read` scope, which the refresh token from API settings does not have. To grant it, run
`go run ./internal/tools/strava-token -reauthorize`.
The credentials are stored unencrypted in `strava-api/token.json` in the user config directory, readable only by the
current user.

## Updating the API spec

The official spec does not describe everything the API returns, so `docs/swagger.json` has local changes on top of it.
`docs/swagger.upstream.json` is an untouched copy of the official spec, do not edit it.

```shell
make update-swagger-spec   # applies what Strava changed, keeping the local changes
make codegen               # regenerates the client
make test-integration      # compares live API responses with the spec
```

A conflict means that Strava changed the same place as a local change. The command stops and leaves the markers in
`docs/swagger.json`, resolve them and run `make codegen`. `make swagger-local-changes` lists all the local changes.

`internal/specaudit` fails when a live response has a field missing from the spec or a value of a different type.
Fix such a field in `docs/swagger.json`.

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
