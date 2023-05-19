package client

import (
	"context"
	"time"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	strava "github.com/obalunenko/strava-api/internal/gen/strava-api-go/client"
	apiactivities "github.com/obalunenko/strava-api/internal/gen/strava-api-go/client/activities"
	"github.com/obalunenko/strava-api/internal/gen/strava-api-go/models"
)

type ActivitiesAPI interface {
	CreateActivity(ctx context.Context, name string, activityType string, sportType string, startDateLocal time.Time, elapsedTime int32, description string, distance float32, trainer int32, commute int32) (models.DetailedActivity, error)
	GetActivityById(ctx context.Context, id int64, opts ...GetActivityByIdOpts) (models.DetailedActivity, error)
	GetCommentsByActivityId(ctx context.Context, id int64, opts ...GetCommentsByActivityIdOpts) ([]models.Comment, error)
	GetKudoersByActivityId(ctx context.Context, id int64, opts ...GetKudoersByActivityIdOpts) ([]models.SummaryAthlete, error)
	GetLapsByActivityId(ctx context.Context, id int64) ([]models.Lap, error)
	GetLoggedInAthleteActivities(ctx context.Context, opts ...GetLoggedInAthleteActivitiesOpts) ([]models.SummaryActivity, error)
	GetZonesByActivityId(ctx context.Context, id int64) ([]models.ActivityZone, error)
	UpdateActivityById(ctx context.Context, id int64, opts ...UpdateActivityByIdOpts) (models.DetailedActivity, error)
}

type GetActivityByIdOpts struct {
	IncludeAllEfforts *bool `json:"include_all_efforts,omitempty"`
}

type GetCommentsByActivityIdOpts struct {
	Page        *int32  `json:"page,omitempty"`
	PerPage     *int32  `json:"per_page,omitempty"`
	PageSize    *int32  `json:"page_size,omitempty"`
	AfterCursor *string `json:"after,omitempty"`
}

type GetKudoersByActivityIdOpts struct {
	Page    *int32 `json:"page,omitempty"`
	PerPage *int32 `json:"per_page,omitempty"`
}

type GetLoggedInAthleteActivitiesOpts struct {
	Before  *int64 `json:"before,omitempty"`
	After   *int64 `json:"after,omitempty"`
	Page    *int32 `json:"page,omitempty"`
	PerPage *int32 `json:"per_page,omitempty"`
}

type UpdateActivityByIdOpts struct {
	Body *models.UpdatableActivity
}

type activitiesService struct {
	client *strava.StravaAPIV3
	auth   runtime.ClientAuthInfoWriter
}

func newActivitiesApiService(client *strava.StravaAPIV3, auth runtime.ClientAuthInfoWriter) ActivitiesAPI {
	return &activitiesService{
		client: client,
		auth:   auth,
	}
}

func (a activitiesService) CreateActivity(ctx context.Context, name string, activityType string, sportType string, startDateLocal time.Time, elapsedTime int32, description string, distance float32, trainer int32, commute int32) (models.DetailedActivity, error) {
	params := apiactivities.NewCreateActivityParams()
	params.SetDefaults()

	params.SetContext(ctx)
	params.SetName(name)

	if activityType != "" {
		params.SetType(&activityType)
	}

	if sportType != "" {
		params.SetType(&sportType)
	}

	params.SetStartDateLocal(strfmt.DateTime(startDateLocal))
	params.SetElapsedTime(int64(elapsedTime))

	if description != "" {
		params.SetDescription(&description)
	}

	if distance != 0 {
		params.SetDistance(&distance)
	}

	t := int64(trainer)
	if trainer != 0 {
		params.SetTrainer(&t)
	}

	c := int64(commute)
	if commute != 0 {
		params.SetCommute(&c)
	}

	resp, err := a.client.Activities.CreateActivity(params, a.auth)
	if err != nil {
		return models.DetailedActivity{}, err
	}

	return *resp.Payload, nil
}

func (a activitiesService) GetActivityById(ctx context.Context, id int64, opts ...GetActivityByIdOpts) (models.DetailedActivity, error) {
	params := apiactivities.NewGetActivityByIDParams()
	params.SetDefaults()

	params.SetContext(ctx)
	params.SetID(id)

	for _, o := range opts {
		if o.IncludeAllEfforts != nil {
			params.SetIncludeAllEfforts(o.IncludeAllEfforts)
		}
	}

	resp, err := a.client.Activities.GetActivityByID(params, a.auth)
	if err != nil {
		return models.DetailedActivity{}, err
	}

	return *resp.Payload, nil
}

func (a activitiesService) GetCommentsByActivityId(ctx context.Context, id int64, opts ...GetCommentsByActivityIdOpts) ([]models.Comment, error) {
	params := apiactivities.NewGetCommentsByActivityIDParams()
	params.SetDefaults()

	params.SetContext(ctx)
	params.SetID(id)

	if len(opts) > 1 {
		return nil, ErrTooManyOptions
	}

	if len(opts) == 1 {
		if opts[0].Page != nil {
			p := int64(*opts[0].Page)
			params.SetPage(&p)
		}

		if opts[0].PerPage != nil {
			p := int64(*opts[0].PerPage)
			params.SetPerPage(&p)
		}

		if opts[0].PageSize != nil {
			p := int64(*opts[0].PageSize)
			params.SetPageSize(&p)
		}

		if opts[0].AfterCursor != nil {
			params.SetAfterCursor(opts[0].AfterCursor)
		}
	}

	resp, err := a.client.Activities.GetCommentsByActivityID(params, a.auth)
	if err != nil {
		return nil, err
	}

	list := make([]models.Comment, len(resp.Payload))

	for i, c := range resp.Payload {
		list[i] = *c
	}

	return list, nil
}

func (a activitiesService) GetKudoersByActivityId(ctx context.Context, id int64, opts ...GetKudoersByActivityIdOpts) ([]models.SummaryAthlete, error) {
	params := apiactivities.NewGetKudoersByActivityIDParams()
	params.SetDefaults()

	params.SetContext(ctx)
	params.SetID(id)

	if len(opts) > 1 {
		return nil, ErrTooManyOptions
	}

	if len(opts) == 1 {
		if opts[0].Page != nil {
			p := int64(*opts[0].Page)
			params.SetPage(&p)
		}

		if opts[0].PerPage != nil {
			p := int64(*opts[0].PerPage)
			params.SetPerPage(&p)
		}
	}

	resp, err := a.client.Activities.GetKudoersByActivityID(params, a.auth)
	if err != nil {
		return nil, err
	}

	list := make([]models.SummaryAthlete, len(resp.Payload))

	for i, a := range resp.Payload {
		list[i] = *a
	}

	return list, nil
}

func (a activitiesService) GetLapsByActivityId(ctx context.Context, id int64) ([]models.Lap, error) {
	params := apiactivities.NewGetLapsByActivityIDParams()
	params.SetDefaults()

	params.SetContext(ctx)
	params.SetID(id)

	resp, err := a.client.Activities.GetLapsByActivityID(params, a.auth)
	if err != nil {
		return nil, err
	}

	list := make([]models.Lap, len(resp.Payload))

	for i, l := range resp.Payload {
		list[i] = *l
	}

	return list, nil
}

func (a activitiesService) GetLoggedInAthleteActivities(ctx context.Context, opts ...GetLoggedInAthleteActivitiesOpts) ([]models.SummaryActivity, error) {
	params := apiactivities.NewGetLoggedInAthleteActivitiesParams()
	params.SetDefaults()

	params.SetContext(ctx)

	if len(opts) > 1 {
		return nil, ErrTooManyOptions
	}

	if len(opts) == 1 {
		if opts[0].Before != nil {
			params.SetBefore(opts[0].Before)
		}

		if opts[0].After != nil {
			params.SetAfter(opts[0].After)
		}

		if opts[0].Page != nil {
			p := int64(*opts[0].Page)
			params.SetPage(&p)
		}

		if opts[0].PerPage != nil {
			p := int64(*opts[0].PerPage)
			params.SetPerPage(&p)
		}
	}

	resp, err := a.client.Activities.GetLoggedInAthleteActivities(params, a.auth)
	if err != nil {
		return nil, err
	}

	list := make([]models.SummaryActivity, len(resp.Payload))

	for i, a := range resp.Payload {
		list[i] = *a
	}

	return list, nil
}

func (a activitiesService) GetZonesByActivityId(ctx context.Context, id int64) ([]models.ActivityZone, error) {
	params := apiactivities.NewGetZonesByActivityIDParams()
	params.SetDefaults()

	params.SetContext(ctx)
	params.SetID(id)

	resp, err := a.client.Activities.GetZonesByActivityID(params, a.auth)
	if err != nil {
		return nil, err
	}

	list := make([]models.ActivityZone, len(resp.Payload))

	for i, z := range resp.Payload {
		list[i] = *z
	}

	return list, nil
}

func (a activitiesService) UpdateActivityById(ctx context.Context, id int64, opts ...UpdateActivityByIdOpts) (models.DetailedActivity, error) {
	params := apiactivities.NewUpdateActivityByIDParams()
	params.SetDefaults()

	params.SetContext(ctx)
	params.SetID(id)

	if len(opts) > 1 {
		return models.DetailedActivity{}, ErrTooManyOptions
	}

	if len(opts) == 1 {
		if opts[0].Body != nil {
			params.SetBody(opts[0].Body)
		}
	}

	resp, err := a.client.Activities.UpdateActivityByID(params, a.auth)
	if err != nil {
		return models.DetailedActivity{}, err
	}

	return *resp.Payload, nil
}
