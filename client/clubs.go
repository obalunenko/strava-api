package client

import (
	"context"

	"github.com/go-openapi/runtime"

	strava "github.com/obalunenko/strava-api/internal/gen/strava-api-go/client"
	"github.com/obalunenko/strava-api/internal/gen/strava-api-go/client/clubs"
	"github.com/obalunenko/strava-api/models"
)

// ClubsAPI is an interface for interacting with clubs endpoints of Strava API
type ClubsAPI interface {
	// GetClubById returns a given club using its identifier
	GetClubById(ctx context.Context, id int64) (models.DetailedClub, error)
	// GetLoggedInAthleteClubs returns a list of the clubs whose membership includes the authenticated athlete
	GetLoggedInAthleteClubs(ctx context.Context, opts ...GetLoggedInAthleteClubsOpts) ([]models.SummaryClub, error)
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
