package client

import (
	"context"

	"github.com/go-openapi/runtime"

	strava "github.com/obalunenko/strava-api/internal/gen/strava-api-go/client"
	"github.com/obalunenko/strava-api/internal/gen/strava-api-go/client/streams"
	"github.com/obalunenko/strava-api/internal/gen/strava-api-go/models"
)

// StreamsAPI is an interface for interacting with streams endpoints of Strava API
type StreamsAPI interface {
	// GetActivityStreams returns a set of streams for an activity identified by its id
	GetActivityStreams(ctx context.Context, id int64, keys []string, keyByType bool) (models.StreamSet, error)
	// GetRouteStreams returns a set of streams for a route identified by its id
	GetRouteStreams(ctx context.Context, id int64) (models.StreamSet, error)
	// GetSegmentEffortStreams returns a set of streams for a segment effort identified by its id
	GetSegmentEffortStreams(ctx context.Context, id int64, keys []string, keyByType bool) (models.StreamSet, error)
	// GetSegmentStreams returns a set of streams for a segment identified by its id
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
	params := streams.NewGetActivityStreamsParams()

	params.SetDefaults()
	params.SetContext(ctx)
	params.SetID(id)
	params.SetKeys(keys)
	params.SetKeyByType(keyByType)

	strm, err := s.client.Streams.GetActivityStreams(params, s.auth)
	if err != nil {
		return models.StreamSet{}, err
	}

	if strm.Payload == nil {
		return models.StreamSet{}, nil
	}

	return *strm.Payload, nil
}

func (s streamsService) GetRouteStreams(ctx context.Context, id int64) (models.StreamSet, error) {
	params := streams.NewGetRouteStreamsParams()

	params.SetDefaults()
	params.SetContext(ctx)
	params.SetID(id)

	strm, err := s.client.Streams.GetRouteStreams(params, s.auth)
	if err != nil {
		return models.StreamSet{}, err
	}

	if strm.Payload == nil {
		return models.StreamSet{}, nil
	}

	return *strm.Payload, nil
}

func (s streamsService) GetSegmentEffortStreams(ctx context.Context, id int64, keys []string, keyByType bool) (models.StreamSet, error) {
	params := streams.NewGetSegmentEffortStreamsParams()

	params.SetDefaults()
	params.SetContext(ctx)
	params.SetID(id)
	params.SetKeys(keys)
	params.SetKeyByType(keyByType)

	strm, err := s.client.Streams.GetSegmentEffortStreams(params, s.auth)
	if err != nil {
		return models.StreamSet{}, err
	}

	if strm.Payload == nil {
		return models.StreamSet{}, nil
	}

	return *strm.Payload, nil
}

func (s streamsService) GetSegmentStreams(ctx context.Context, id int64, keys []string, keyByType bool) (models.StreamSet, error) {
	params := streams.NewGetSegmentStreamsParams()

	params.SetDefaults()
	params.SetContext(ctx)
	params.SetID(id)
	params.SetKeys(keys)
	params.SetKeyByType(keyByType)

	strm, err := s.client.Streams.GetSegmentStreams(params, s.auth)
	if err != nil {
		return models.StreamSet{}, err
	}

	if strm.Payload == nil {
		return models.StreamSet{}, nil
	}

	return *strm.Payload, nil
}
