package client

import (
	"context"

	"github.com/go-openapi/runtime"

	strava "github.com/obalunenko/strava-api/internal/gen/strava-api-go/client"
	"github.com/obalunenko/strava-api/internal/gen/strava-api-go/client/routes"
	"github.com/obalunenko/strava-api/internal/gen/strava-api-go/models"
)

// RoutesAPI is an interface for interacting with routes endpoints of Strava API
type RoutesAPI interface {
	// GetRouteAsGPX returns a GPX file of the route
	GetRouteAsGPX(ctx context.Context, id int64) error
	// GetRouteAsTCX returns a TCX file of the route
	GetRouteAsTCX(ctx context.Context, id int64) error
	// GetRouteById returns a route with given id
	GetRouteById(ctx context.Context, id int64) (models.Route, error)
	// GetRoutesByAthleteId returns a list of routes for the given athlete
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

func (r routesService) GetRouteAsGPX(ctx context.Context, id int64) error {
	params := routes.NewGetRouteAsGPXParams()

	params.SetDefaults()
	params.SetContext(ctx)
	params.SetID(id)

	_, err := r.client.Routes.GetRouteAsGPX(params, r.auth)
	if err != nil {
		return err
	}

	return nil
}

func (r routesService) GetRouteAsTCX(ctx context.Context, id int64) error {
	params := routes.NewGetRouteAsTCXParams()

	params.SetDefaults()
	params.SetContext(ctx)
	params.SetID(id)

	_, err := r.client.Routes.GetRouteAsTCX(params, r.auth)
	if err != nil {
		return err
	}

	return nil
}

func (r routesService) GetRouteById(ctx context.Context, id int64) (models.Route, error) {
	params := routes.NewGetRouteByIDParams()

	params.SetDefaults()
	params.SetContext(ctx)
	params.SetID(id)

	route, err := r.client.Routes.GetRouteByID(params, r.auth)
	if err != nil {
		return models.Route{}, err
	}

	return convertToModelsRoute(route.GetPayload()), nil
}

func (r routesService) GetRoutesByAthleteId(ctx context.Context, opts ...GetRoutesByAthleteIdOpts) ([]models.Route, error) {
	params := routes.NewGetRoutesByAthleteIDParams()

	params.SetDefaults()
	params.SetContext(ctx)

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

	resp, err := r.client.Routes.GetRoutesByAthleteID(params, r.auth)
	if err != nil {
		return nil, err
	}

	list := make([]models.Route, 0, len(resp.GetPayload()))

	for _, route := range resp.GetPayload() {
		list = append(list, *route)
	}

	return list, nil
}

// convertToModelsRoute converts routes.GetRouteByIDOKBody to models.Route
func convertToModelsRoute(route *models.Route) models.Route {
	if route == nil {
		return models.Route{}
	}

	return *route
}
