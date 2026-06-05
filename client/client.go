// Package client provides a wrapper around the generated Strava API client.
// It provides a more convenient interface for interacting with the API.
package client

import (
	openapiruntime "github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"

	apiclient "github.com/obalunenko/strava-api/internal/gen/strava-api-go/client"
)

// APIClient is a wrapper around the generated Strava API client.
// It provides a more convenient interface for interacting with the API.
type APIClient struct {
	Activities     ActivitiesAPI
	Athletes       AthletesAPI
	Clubs          ClubsAPI
	Routes         RoutesAPI
	Gears          GearsAPI
	SegmentEfforts SegmentEffortsAPI
	Segments       SegmentsAPI
	Streams        StreamsAPI
	Uploads        UploadsAPI
}

// NewAPIClient creates a new APIClient. Requires a Strava API token.
func NewAPIClient(token string) (*APIClient, error) {
	client := apiclient.NewHTTPClient(nil)
	registerRouteExportConsumers(client.Transport)

	// Set the bearer token for auth.
	auth := httptransport.BearerToken(token)

	return &APIClient{
		Activities:     newActivitiesApiService(client, auth),
		Athletes:       newAthletesApiService(client, auth),
		Clubs:          newClubsApiService(client, auth),
		Routes:         newRoutesApiService(client, auth),
		Gears:          newGearsApiService(client, auth),
		SegmentEfforts: newSegmentEffortsApiService(client, auth),
		Segments:       newSegmentsService(client, auth),
		Streams:        newStreamsApiService(client, auth),
		Uploads:        newUploadsApiService(client, auth),
	}, nil
}

func registerRouteExportConsumers(transport openapiruntime.ClientTransport) {
	runtimeTransport, ok := transport.(*httptransport.Runtime)
	if !ok {
		return
	}

	consumer := openapiruntime.ByteStreamConsumer()
	runtimeTransport.Consumers["application/gpx+xml"] = consumer
	runtimeTransport.Consumers["application/tcx+xml"] = consumer
	runtimeTransport.Consumers["application/vnd.garmin.tcx+xml"] = consumer
	runtimeTransport.Consumers[openapiruntime.XMLMime] = consumer
	runtimeTransport.Consumers[openapiruntime.TextMime] = consumer
}
