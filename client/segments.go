package client

import (
	"context"

	"github.com/go-openapi/runtime"

	strava "github.com/obalunenko/strava-api/internal/gen/strava-api-go/client"
	"github.com/obalunenko/strava-api/internal/gen/strava-api-go/client/segments"
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
	params := segments.NewExploreSegmentsParams()

	params.SetDefaults()
	params.SetContext(ctx)
	params.SetBounds(bounds)

	if len(opts) > 1 {
		return models.ExplorerResponse{}, ErrTooManyOptions
	}

	if len(opts) == 1 {
		for _, opt := range opts {
			if opt.MinCat != nil {
				p := int64(*opt.MinCat)
				params.SetMinCat(&p)
			}
			if opt.MaxCat != nil {
				p := int64(*opt.MaxCat)
				params.SetMaxCat(&p)
			}
			if opt.ActivityType != nil {
				params.SetActivityType(opt.ActivityType)
			}
		}
	}

	resp, err := s.client.Segments.ExploreSegments(params, s.auth)
	if err != nil {
		return models.ExplorerResponse{}, err
	}

	return *resp.Payload, nil
}

func (s segmentsService) GetLoggedInAthleteStarredSegments(ctx context.Context, opts ...GetLoggedInAthleteStarredSegmentsOpts) ([]models.SummarySegment, error) {
	params := segments.NewGetLoggedInAthleteStarredSegmentsParams()

	params.SetDefaults()
	params.SetContext(ctx)

	if len(opts) > 1 {
		return nil, ErrTooManyOptions
	}

	if len(opts) == 1 {
		for _, opt := range opts {
			if opt.Page != nil {
				p := int64(*opt.Page)
				params.SetPage(&p)
			}
			if opt.PerPage != nil {
				p := int64(*opt.PerPage)
				params.SetPerPage(&p)
			}
		}
	}

	resp, err := s.client.Segments.GetLoggedInAthleteStarredSegments(params, s.auth)
	if err != nil {
		return nil, err
	}

	list := make([]models.SummarySegment, len(resp.Payload))

	for i, v := range resp.Payload {
		list[i] = *v
	}

	return list, nil
}

func (s segmentsService) GetSegmentById(ctx context.Context, id int64) (models.DetailedSegment, error) {
	params := segments.NewGetSegmentByIDParams()

	params.SetDefaults()
	params.SetContext(ctx)
	params.SetID(id)

	resp, err := s.client.Segments.GetSegmentByID(params, s.auth)
	if err != nil {
		return models.DetailedSegment{}, err
	}

	return *resp.Payload, nil
}

func (s segmentsService) StarSegment(ctx context.Context, starred bool, id int64) (models.DetailedSegment, error) {
	params := segments.NewStarSegmentParams()

	params.SetDefaults()
	params.SetContext(ctx)
	params.SetStarred(starred)
	params.SetID(id)

	resp, err := s.client.Segments.StarSegment(params, s.auth)
	if err != nil {
		return models.DetailedSegment{}, err
	}

	return *resp.Payload, nil
}
