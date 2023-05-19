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

func (activitiesService) GetActivityById(ctx context.Context, id int64, opts ...GetActivityByIdOpts) (models.DetailedActivity, error) {
	//TODO implement me
	panic("implement me")
}

func (activitiesService) GetCommentsByActivityId(ctx context.Context, id int64, opts ...GetCommentsByActivityIdOpts) ([]models.Comment, error) {
	//TODO implement me
	panic("implement me")
}

func (activitiesService) GetKudoersByActivityId(ctx context.Context, id int64, opts ...GetKudoersByActivityIdOpts) ([]models.SummaryAthlete, error) {
	//TODO implement me
	panic("implement me")
}

func (activitiesService) GetLapsByActivityId(ctx context.Context, id int64) ([]models.Lap, error) {
	//TODO implement me
	panic("implement me")
}

func (activitiesService) GetLoggedInAthleteActivities(ctx context.Context, opts ...GetLoggedInAthleteActivitiesOpts) ([]models.SummaryActivity, error) {
	//TODO implement me
	panic("implement me")
}

func (activitiesService) GetZonesByActivityId(ctx context.Context, id int64) ([]models.ActivityZone, error) {
	//TODO implement me
	panic("implement me")
}

func (activitiesService) UpdateActivityById(ctx context.Context, id int64, opts ...UpdateActivityByIdOpts) (models.DetailedActivity, error) {
	//TODO implement me
	panic("implement me")
}
