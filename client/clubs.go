package client

import (
	"context"

	"github.com/go-openapi/runtime"

	strava "github.com/obalunenko/strava-api/internal/gen/strava-api-go/client"
	"github.com/obalunenko/strava-api/internal/gen/strava-api-go/client/clubs"
	"github.com/obalunenko/strava-api/internal/gen/strava-api-go/models"
)

// ClubsAPI is an interface for interacting with clubs endpoints of Strava API
type ClubsAPI interface {
	// GetClubActivitiesById returns a list of the recent activities performed by members of a given club
	GetClubActivitiesById(ctx context.Context, id int64, opts ...GetClubActivitiesByIdOpts) ([]models.ClubActivity, error)
	// GetClubAdminsById returns a list of the administrators of a given club
	GetClubAdminsById(ctx context.Context, id int64, opts ...GetClubAdminsByIdOpts) ([]models.SummaryAthlete, error)
	// GetClubById returns a given club using its identifier
	GetClubById(ctx context.Context, id int64) (models.DetailedClub, error)
	// GetClubMembersById returns a list of the athletes who are members of a given club
	GetClubMembersById(ctx context.Context, id int64, opts ...GetClubMembersByIdOpts) ([]models.ClubAthlete, error)
	// GetLoggedInAthleteClubs returns a list of the clubs whose membership includes the authenticated athlete
	GetLoggedInAthleteClubs(ctx context.Context, opts ...GetLoggedInAthleteClubsOpts) ([]models.SummaryClub, error)
}

// GetClubActivitiesByIdOpts is an option for GetClubActivitiesById.
type GetClubActivitiesByIdOpts struct {
	Page    *int32
	PerPage *int32
}

// GetClubAdminsByIdOpts is an option for GetClubAdminsById.
type GetClubAdminsByIdOpts struct {
	Page    *int32
	PerPage *int32
}

// GetClubMembersByIdOpts is an option for GetClubMembersById.
type GetClubMembersByIdOpts struct {
	Page    *int32
	PerPage *int32
}

// GetLoggedInAthleteClubsOpts is an option for GetLoggedInAthleteClubs.
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

	return convertToListClubActivity(activities.GetPayload()), nil
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

	return convertToListSummaryAthlete(admins.GetPayload()), nil
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

	return convertToDetailedClub(club.GetPayload()), nil
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

	return convertToListClubAthlete(members.GetPayload()), nil
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

	return convertToListSummaryClub(list.GetPayload()), nil
}

func convertToClubActivity(activity *models.ClubActivity) models.ClubActivity {
	if activity == nil {
		return models.ClubActivity{}
	}

	return *activity
}

func convertToSummaryAthlete(athlete *models.SummaryAthlete) models.SummaryAthlete {
	if athlete == nil {
		return models.SummaryAthlete{}
	}

	return *athlete
}

func convertToClubAthlete(athlete *models.ClubAthlete) models.ClubAthlete {
	if athlete == nil {
		return models.ClubAthlete{}
	}

	return *athlete
}

func convertToSummaryClub(club *models.SummaryClub) models.SummaryClub {
	if club == nil {
		return models.SummaryClub{}
	}

	return *club
}

func convertToDetailedClub(club *models.DetailedClub) models.DetailedClub {
	if club == nil {
		return models.DetailedClub{}
	}

	return *club
}

func convertToListSummaryClub(clubs []*models.SummaryClub) []models.SummaryClub {
	if clubs == nil {
		return nil
	}

	list := make([]models.SummaryClub, len(clubs))

	for i, club := range clubs {
		list[i] = convertToSummaryClub(club)
	}

	return list
}

func convertToListClubActivity(activities []*models.ClubActivity) []models.ClubActivity {
	if activities == nil {
		return nil
	}

	list := make([]models.ClubActivity, len(activities))

	for i, activity := range activities {
		list[i] = convertToClubActivity(activity)
	}

	return list
}

func convertToListClubAthlete(athletes []*models.ClubAthlete) []models.ClubAthlete {
	if athletes == nil {
		return nil
	}

	list := make([]models.ClubAthlete, len(athletes))

	for i, athlete := range athletes {
		list[i] = convertToClubAthlete(athlete)
	}

	return list
}

func convertToListSummaryAthlete(athletes []*models.SummaryAthlete) []models.SummaryAthlete {
	if athletes == nil {
		return nil
	}

	list := make([]models.SummaryAthlete, len(athletes))

	for i, athlete := range athletes {
		list[i] = convertToSummaryAthlete(athlete)
	}

	return list
}
