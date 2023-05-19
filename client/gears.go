package client

import (
	"context"

	"github.com/go-openapi/runtime"

	strava "github.com/obalunenko/strava-api/internal/gen/strava-api-go/client"
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
	// TODO implement me
	panic("implement me")
}
