package client

import (
	"context"

	"github.com/go-openapi/runtime"

	strava "github.com/obalunenko/strava-api/internal/gen/strava-api-go/client"
	"github.com/obalunenko/strava-api/internal/gen/strava-api-go/client/segments"
	"github.com/obalunenko/strava-api/internal/gen/strava-api-go/models"
)

// SegmentsAPI is an interface for interacting with segments endpoints of Strava API
type SegmentsAPI interface {
	// ExploreSegments returns the top segments matching a specified query
	ExploreSegments(ctx context.Context, bounds []float32, opts ...ExploreSegmentsOpts) (models.ExplorerResponse, error)
	// GetLoggedInAthleteStarredSegments returns the segments starred by the authenticated user
	GetLoggedInAthleteStarredSegments(ctx context.Context, opts ...GetLoggedInAthleteStarredSegmentsOpts) ([]models.SummarySegment, error)
	// GetSegmentById returns the specified segment
	GetSegmentById(ctx context.Context, id int64) (models.DetailedSegment, error)
	// StarSegment stars/ unstars the given segment for the authenticated athlete
	StarSegment(ctx context.Context, starred bool, id int64) (models.DetailedSegment, error)
}

// ExploreSegmentsOpts is a set of optional parameters for ExploreSegments method
type ExploreSegmentsOpts struct {
	MinCat       *int32
	MaxCat       *int32
	ActivityType *string
}

// GetLoggedInAthleteStarredSegmentsOpts is a set of optional parameters for GetLoggedInAthleteStarredSegments method
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

	return convertToExplorerResponse(resp.GetPayload()), nil
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

	return convertToListSummarySegment(resp.GetPayload()), nil
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

	return convertToDetailedSegment(resp.GetPayload()), nil
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

	return convertToDetailedSegment(resp.GetPayload()), nil
}

func convertToDetailedSegment(segment *models.DetailedSegment) models.DetailedSegment {
	if segment == nil {
		return models.DetailedSegment{}
	}

	return *segment
}

func convertToSummarySegment(segment *models.SummarySegment) models.SummarySegment {
	if segment == nil {
		return models.SummarySegment{}
	}

	return *segment
}

func convertToListSummarySegment(segments []*models.SummarySegment) []models.SummarySegment {
	if segments == nil {
		return nil
	}

	list := make([]models.SummarySegment, len(segments))

	for i, segment := range segments {
		list[i] = convertToSummarySegment(segment)
	}

	return list
}

func convertToExplorerResponse(response *models.ExplorerResponse) models.ExplorerResponse {
	if response == nil {
		return models.ExplorerResponse{}
	}

	return *response
}
