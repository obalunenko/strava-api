package client

import (
	"context"

	"github.com/go-openapi/runtime"

	strava "github.com/obalunenko/strava-api/internal/gen/strava-api-go/client"
	apiathletes "github.com/obalunenko/strava-api/internal/gen/strava-api-go/client/athletes"
	"github.com/obalunenko/strava-api/internal/gen/strava-api-go/models"
)

type AthletesAPI interface {
	GetLoggedInAthlete(ctx context.Context) (models.DetailedAthlete, error)
	GetLoggedInAthleteZones(ctx context.Context) (models.Zones, error)
	GetStats(ctx context.Context, id int64) (models.ActivityStats, error)
	UpdateLoggedInAthlete(ctx context.Context, weight float32) (models.DetailedAthlete, error)
}

func newAthletesApiService(client *strava.StravaAPIV3, auth runtime.ClientAuthInfoWriter) AthletesAPI {
	return athletesAPI{
		client: client,
		auth:   auth,
	}
}

type athletesAPI struct {
	client *strava.StravaAPIV3
	auth   runtime.ClientAuthInfoWriter
}

func (a athletesAPI) GetLoggedInAthlete(ctx context.Context) (models.DetailedAthlete, error) {
	params := apiathletes.NewGetLoggedInAthleteParams()
	params.SetDefaults()
	params.SetContext(ctx)

	athlete, err := a.client.Athletes.GetLoggedInAthlete(params, a.auth)
	if err != nil {
		return models.DetailedAthlete{}, err
	}

	return *athlete.Payload, nil
}

func (a athletesAPI) GetLoggedInAthleteZones(ctx context.Context) (models.Zones, error) {
	params := apiathletes.NewGetLoggedInAthleteZonesParams()
	params.SetDefaults()
	params.SetContext(ctx)

	athlete, err := a.client.Athletes.GetLoggedInAthleteZones(params, a.auth)
	if err != nil {
		return models.Zones{}, err
	}

	return *athlete.Payload, nil
}

func (a athletesAPI) GetStats(ctx context.Context, id int64) (models.ActivityStats, error) {
	params := apiathletes.NewGetStatsParams()
	params.SetDefaults()
	params.SetContext(ctx)
	params.SetID(id)

	stats, err := a.client.Athletes.GetStats(params, a.auth)
	if err != nil {
		return models.ActivityStats{}, err
	}

	return *stats.Payload, nil
}

func (a athletesAPI) UpdateLoggedInAthlete(ctx context.Context, weight float32) (models.DetailedAthlete, error) {
	params := apiathletes.NewUpdateLoggedInAthleteParams()
	params.SetDefaults()
	params.SetContext(ctx)
	params.SetWeight(weight)

	athlete, err := a.client.Athletes.UpdateLoggedInAthlete(params, a.auth)
	if err != nil {
		return models.DetailedAthlete{}, err
	}

	return *athlete.Payload, nil
}
