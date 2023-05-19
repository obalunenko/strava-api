package client

import (
	"context"

	"github.com/go-openapi/runtime"

	strava "github.com/obalunenko/strava-api/internal/gen/strava-api-go/client"
	"github.com/obalunenko/strava-api/internal/gen/strava-api-go/client/gears"
	"github.com/obalunenko/strava-api/internal/gen/strava-api-go/models"
)

type GearsAPI interface {
	GetGearById(ctx context.Context, id string) (models.DetailedGear, error)
}

type gearsService struct {
	client *strava.StravaAPIV3
	auth   runtime.ClientAuthInfoWriter
}

func newGearsApiService(client *strava.StravaAPIV3, auth runtime.ClientAuthInfoWriter) GearsAPI {
	return gearsService{
		client: client,
		auth:   auth,
	}
}

func (g gearsService) GetGearById(ctx context.Context, id string) (models.DetailedGear, error) {
	params := gears.NewGetGearByIDParams()
	params.SetDefaults()
	params.SetContext(ctx)
	params.SetID(id)

	gear, err := g.client.Gears.GetGearByID(params, g.auth)
	if err != nil {
		return models.DetailedGear{}, err
	}

	return *gear.GetPayload(), nil
}
