package client

import (
	"context"

	"github.com/go-openapi/runtime"

	strava "github.com/obalunenko/strava-api/internal/gen/strava-api-go/client"
	"github.com/obalunenko/strava-api/internal/gen/strava-api-go/client/clubs"
	"github.com/obalunenko/strava-api/internal/gen/strava-api-go/models"
)

type ClubsAPI interface {
	GetClubActivitiesById(ctx context.Context, id int64, opts ...GetClubActivitiesByIdOpts) ([]models.ClubActivity, error)
	GetClubAdminsById(ctx context.Context, id int64, opts ...GetClubAdminsByIdOpts) ([]models.SummaryAthlete, error)
	GetClubById(ctx context.Context, id int64) (models.DetailedClub, error)
	GetClubMembersById(ctx context.Context, id int64, opts ...GetClubMembersByIdOpts) ([]models.ClubAthlete, error)
	GetLoggedInAthleteClubs(ctx context.Context, opts ...GetLoggedInAthleteClubsOpts) ([]models.SummaryClub, error)
}

type GetClubActivitiesByIdOpts struct {
	Page    *int32
	PerPage *int32
}

type GetClubAdminsByIdOpts struct {
	Page    *int32
	PerPage *int32
}

type GetClubMembersByIdOpts struct {
	Page    *int32
	PerPage *int32
}

type GetLoggedInAthleteClubsOpts struct {
	Page    *int32
	PerPage *int32
}

type clubsService struct {
	client *strava.StravaAPIV3
	auth   runtime.ClientAuthInfoWriter
}

func newClubsApiService(client *strava.StravaAPIV3, auth runtime.ClientAuthInfoWriter) ClubsAPI {
	return clubsService{
		client: client,
		auth:   auth,
	}
}

func (c clubsService) GetClubActivitiesById(ctx context.Context, id int64, opts ...GetClubActivitiesByIdOpts) ([]models.ClubActivity, error) {
	params := clubs.NewGetClubActivitiesByIDParams()

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

	activities, err := c.client.Clubs.GetClubActivitiesByID(params, c.auth)
	if err != nil {
		return nil, err
	}

	resp := make([]models.ClubActivity, 0, len(activities.GetPayload()))

	for _, a := range activities.GetPayload() {
		resp = append(resp, *a)
	}

	return resp, nil
}

func (c clubsService) GetClubAdminsById(ctx context.Context, id int64, opts ...GetClubAdminsByIdOpts) ([]models.SummaryAthlete, error) {
	params := clubs.NewGetClubAdminsByIDParams()

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

	admins, err := c.client.Clubs.GetClubAdminsByID(params, c.auth)
	if err != nil {
		return nil, err
	}

	resp := make([]models.SummaryAthlete, 0, len(admins.GetPayload()))

	for _, a := range admins.GetPayload() {
		resp = append(resp, *a)
	}

	return resp, nil
}

func (c clubsService) GetClubById(ctx context.Context, id int64) (models.DetailedClub, error) {
	params := clubs.NewGetClubByIDParams()

	params.SetDefaults()
	params.SetContext(ctx)
	params.SetID(id)

	club, err := c.client.Clubs.GetClubByID(params, c.auth)
	if err != nil {
		return models.DetailedClub{}, err
	}

	return *club.GetPayload(), nil
}

func (c clubsService) GetClubMembersById(ctx context.Context, id int64, opts ...GetClubMembersByIdOpts) ([]models.ClubAthlete, error) {
	params := clubs.NewGetClubMembersByIDParams()

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

	members, err := c.client.Clubs.GetClubMembersByID(params, c.auth)
	if err != nil {
		return nil, err
	}

	resp := make([]models.ClubAthlete, 0, len(members.GetPayload()))

	for _, m := range members.GetPayload() {
		resp = append(resp, *m)
	}

	return resp, nil
}

func (c clubsService) GetLoggedInAthleteClubs(ctx context.Context, opts ...GetLoggedInAthleteClubsOpts) ([]models.SummaryClub, error) {
	params := clubs.NewGetLoggedInAthleteClubsParams()

	params.SetDefaults()
	params.SetContext(ctx)

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

	list, err := c.client.Clubs.GetLoggedInAthleteClubs(params, c.auth)
	if err != nil {
		return nil, err
	}

	resp := make([]models.SummaryClub, 0, len(list.GetPayload()))

	for _, cl := range list.GetPayload() {
		resp = append(resp, *cl)
	}

	return resp, nil
}
