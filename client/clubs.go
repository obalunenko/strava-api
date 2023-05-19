package client

import (
	"context"
	"github.com/go-openapi/runtime"
	strava "github.com/obalunenko/strava-api/internal/gen/strava-api-go/client"
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

func (clubsService) GetClubActivitiesById(ctx context.Context, id int64, opts ...GetClubActivitiesByIdOpts) ([]models.ClubActivity, error) {
	//TODO implement me
	panic("implement me")
}

func (clubsService) GetClubAdminsById(ctx context.Context, id int64, opts ...GetClubAdminsByIdOpts) ([]models.SummaryAthlete, error) {
	//TODO implement me
	panic("implement me")
}

func (clubsService) GetClubById(ctx context.Context, id int64) (models.DetailedClub, error) {
	//TODO implement me
	panic("implement me")
}

func (clubsService) GetClubMembersById(ctx context.Context, id int64, opts ...GetClubMembersByIdOpts) ([]models.ClubAthlete, error) {
	//TODO implement me
	panic("implement me")
}

func (clubsService) GetLoggedInAthleteClubs(ctx context.Context, opts ...GetLoggedInAthleteClubsOpts) ([]models.SummaryClub, error) {
	//TODO implement me
	panic("implement me")
}
