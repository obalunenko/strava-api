package client

import (
	"context"
	"github.com/go-openapi/runtime"
	strava "github.com/obalunenko/strava-api/internal/gen/strava-api-go/client"
	"github.com/obalunenko/strava-api/internal/gen/strava-api-go/models"
)

type RoutesAPI interface {
	GetRouteAsGPX(ctx context.Context, id int64) error
	GetRouteAsTCX(ctx context.Context, id int64) error
	GetRouteById(ctx context.Context, id int64) (models.Route, error)
	GetRoutesByAthleteId(ctx context.Context, opts ...GetRoutesByAthleteIdOpts) ([]models.Route, error)
}

type GetRoutesByAthleteIdOpts struct {
	Page    *int32
	PerPage *int32
}

func newRoutesApiService(client *strava.StravaAPIV3, auth runtime.ClientAuthInfoWriter) RoutesAPI {
	return routesService{
		client: client,
		auth:   auth,
	}
}

type routesService struct {
	client *strava.StravaAPIV3
	auth   runtime.ClientAuthInfoWriter
}

func (routesService) GetRouteAsGPX(ctx context.Context, id int64) error {
	//TODO implement me
	panic("implement me")
}

func (routesService) GetRouteAsTCX(ctx context.Context, id int64) error {
	//TODO implement me
	panic("implement me")
}

func (routesService) GetRouteById(ctx context.Context, id int64) (models.Route, error) {
	//TODO implement me
	panic("implement me")
}

func (routesService) GetRoutesByAthleteId(ctx context.Context, opts ...GetRoutesByAthleteIdOpts) ([]models.Route, error) {
	//TODO implement me
	panic("implement me")
}
