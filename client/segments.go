package client

import (
	"context"
	"github.com/go-openapi/runtime"
	strava "github.com/obalunenko/strava-api/internal/gen/strava-api-go/client"
	"github.com/obalunenko/strava-api/internal/gen/strava-api-go/models"
)

type SegmentsAPI interface {
	ExploreSegments(ctx context.Context, bounds []float32, opts ...ExploreSegmentsOpts) (models.ExplorerResponse, error)
	GetLoggedInAthleteStarredSegments(ctx context.Context, opts ...GetLoggedInAthleteStarredSegmentsOpts) ([]models.SummarySegment, error)
	GetSegmentById(ctx context.Context, id int64) (models.DetailedSegment, error)
	StarSegment(ctx context.Context, starred bool, id int64) (models.DetailedSegment, error)
}

type ExploreSegmentsOpts struct {
	MinCat       *int32
	MaxCat       *int32
	ActivityType *string
}

type GetLoggedInAthleteStarredSegmentsOpts struct {
	Page    *int32
	PerPage *int32
}

func newSegmentsService(client *strava.StravaAPIV3, auth runtime.ClientAuthInfoWriter) SegmentsAPI {
	return segmentsService{
		client: client,
		auth:   auth,
	}
}

type segmentsService struct {
	client *strava.StravaAPIV3
	auth   runtime.ClientAuthInfoWriter
}

func (s segmentsService) ExploreSegments(ctx context.Context, bounds []float32, opts ...ExploreSegmentsOpts) (models.ExplorerResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s segmentsService) GetLoggedInAthleteStarredSegments(ctx context.Context, opts ...GetLoggedInAthleteStarredSegmentsOpts) ([]models.SummarySegment, error) {
	//TODO implement me
	panic("implement me")
}

func (s segmentsService) GetSegmentById(ctx context.Context, id int64) (models.DetailedSegment, error) {
	//TODO implement me
	panic("implement me")
}

func (s segmentsService) StarSegment(ctx context.Context, starred bool, id int64) (models.DetailedSegment, error) {
	//TODO implement me
	panic("implement me")
}
