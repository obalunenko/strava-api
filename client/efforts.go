package client

import (
	"context"
	"time"

	"github.com/go-openapi/runtime"

	strava "github.com/obalunenko/strava-api/internal/gen/strava-api-go/client"
	"github.com/obalunenko/strava-api/internal/gen/strava-api-go/models"
)

type SegmentEffortsAPI interface {
	GetEffortsBySegmentId(ctx context.Context, segmentId int32, opts ...GetEffortsBySegmentIdOpts) ([]models.DetailedSegmentEffort, error)
	GetSegmentEffortById(ctx context.Context, id int64) (models.DetailedSegmentEffort, error)
}

type GetEffortsBySegmentIdOpts struct {
	StartDateLocal *time.Time
	EndDateLocal   *time.Time
	PerPage        *int32
}

func newSegmentEffortsApiService(client *strava.StravaAPIV3, auth runtime.ClientAuthInfoWriter) SegmentEffortsAPI {
	return segmentsEffortsService{
		client: client,
		auth:   auth,
	}
}

type segmentsEffortsService struct {
	client *strava.StravaAPIV3
	auth   runtime.ClientAuthInfoWriter
}

func (s segmentsEffortsService) GetEffortsBySegmentId(ctx context.Context, segmentId int32, opts ...GetEffortsBySegmentIdOpts) ([]models.DetailedSegmentEffort, error) {
	// TODO implement me
	panic("implement me")
}

func (s segmentsEffortsService) GetSegmentEffortById(ctx context.Context, id int64) (models.DetailedSegmentEffort, error) {
	// TODO implement me
	panic("implement me")
}
