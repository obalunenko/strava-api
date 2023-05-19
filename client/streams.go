package client

import (
	"context"

	"github.com/go-openapi/runtime"

	strava "github.com/obalunenko/strava-api/internal/gen/strava-api-go/client"
	"github.com/obalunenko/strava-api/internal/gen/strava-api-go/models"
)

type StreamsAPI interface {
	GetActivityStreams(ctx context.Context, id int64, keys []string, keyByType bool) (models.StreamSet, error)
	GetRouteStreams(ctx context.Context, id int64) (models.StreamSet, error)
	GetSegmentEffortStreams(ctx context.Context, id int64, keys []string, keyByType bool) (models.StreamSet, error)
	GetSegmentStreams(ctx context.Context, id int64, keys []string, keyByType bool) (models.StreamSet, error)
}

func newStreamsApiService(client *strava.StravaAPIV3, auth runtime.ClientAuthInfoWriter) StreamsAPI {
	return streamsService{
		client: client,
		auth:   auth,
	}
}

type streamsService struct {
	client *strava.StravaAPIV3
	auth   runtime.ClientAuthInfoWriter
}

func (s streamsService) GetActivityStreams(ctx context.Context, id int64, keys []string, keyByType bool) (models.StreamSet, error) {
	// TODO implement me
	panic("implement me")
}

func (s streamsService) GetRouteStreams(ctx context.Context, id int64) (models.StreamSet, error) {
	// TODO implement me
	panic("implement me")
}

func (s streamsService) GetSegmentEffortStreams(ctx context.Context, id int64, keys []string, keyByType bool) (models.StreamSet, error) {
	// TODO implement me
	panic("implement me")
}

func (s streamsService) GetSegmentStreams(ctx context.Context, id int64, keys []string, keyByType bool) (models.StreamSet, error) {
	// TODO implement me
	panic("implement me")
}
