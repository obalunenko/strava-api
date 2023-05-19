package client

import (
	"context"
	"time"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	strava "github.com/obalunenko/strava-api/internal/gen/strava-api-go/client"
	"github.com/obalunenko/strava-api/internal/gen/strava-api-go/client/segment_efforts"
	"github.com/obalunenko/strava-api/internal/gen/strava-api-go/models"
)

// SegmentEffortsAPI is an interface for interacting with segment_efforts endpoints of Strava API
type SegmentEffortsAPI interface {
	// GetEffortsBySegmentId returns a set of the authenticated athlete's segment efforts for a given segment
	GetEffortsBySegmentId(ctx context.Context, segmentId int32, opts ...GetEffortsBySegmentIdOpts) ([]models.DetailedSegmentEffort, error)
	// GetSegmentEffortById returns a segment effort from an activity that is owned by the authenticated athlete
	GetSegmentEffortById(ctx context.Context, id int64) (models.DetailedSegmentEffort, error)
}

// GetEffortsBySegmentIdOpts is an options struct for GetEffortsBySegmentId method
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
	params := segment_efforts.NewGetEffortsBySegmentIDParams()

	params.SetDefaults()
	params.SetContext(ctx)

	params.SetSegmentID(int64(segmentId))

	if len(opts) > 1 {
		return nil, ErrTooManyOptions
	}

	if len(opts) == 1 {
		opt := opts[0]
		if opt.StartDateLocal != nil {
			d := strfmt.DateTime(*opt.StartDateLocal)
			params.SetStartDateLocal(&d)
		}

		if opt.EndDateLocal != nil {
			d := strfmt.DateTime(*opt.EndDateLocal)
			params.SetEndDateLocal(&d)
		}

		if opt.PerPage != nil {
			p := int64(*opt.PerPage)
			params.SetPerPage(&p)
		}
	}

	resp, err := s.client.SegmentEfforts.GetEffortsBySegmentID(params, s.auth)
	if err != nil {
		return nil, err
	}

	list := make([]models.DetailedSegmentEffort, len(resp.Payload))

	for i, effort := range resp.Payload {
		list[i] = *effort
	}

	return list, nil
}

func (s segmentsEffortsService) GetSegmentEffortById(ctx context.Context, id int64) (models.DetailedSegmentEffort, error) {
	params := segment_efforts.NewGetSegmentEffortByIDParams()

	params.SetDefaults()
	params.SetContext(ctx)

	params.SetID(id)

	resp, err := s.client.SegmentEfforts.GetSegmentEffortByID(params, s.auth)
	if err != nil {
		return models.DetailedSegmentEffort{}, err
	}

	return convertToSegmentEffort(resp.GetPayload()), nil
}

// convert to models.DetailedSegmentEffort
func convertToSegmentEffort(effort *segment_efforts.GetSegmentEffortByIDOKBody) models.DetailedSegmentEffort {
	return models.DetailedSegmentEffort{
		SummarySegmentEffort: effort.SummarySegmentEffort,
		Activity:             effort.Activity,
		Athlete:              effort.Athlete,
		AverageCadence:       effort.AverageCadence,
		AverageHeartrate:     effort.AverageHeartrate,
		AverageWatts:         effort.AverageWatts,
		DeviceWatts:          effort.DeviceWatts,
		EndIndex:             effort.EndIndex,
		Hidden:               effort.Hidden,
		KomRank:              effort.KomRank,
		MaxHeartrate:         effort.MaxHeartrate,
		MovingTime:           effort.MovingTime,
		Name:                 effort.Name,
		PrRank:               effort.PrRank,
		Segment:              effort.Segment,
		StartIndex:           effort.StartIndex,
	}
}
