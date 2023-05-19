package client

import (
	"context"
	"io"

	"github.com/go-openapi/runtime"

	strava "github.com/obalunenko/strava-api/internal/gen/strava-api-go/client"
	"github.com/obalunenko/strava-api/internal/gen/strava-api-go/client/uploads"
	"github.com/obalunenko/strava-api/internal/gen/strava-api-go/models"
)

type UploadsAPI interface {
	CreateUpload(ctx context.Context, opts ...CreateUploadParams) (models.Upload, error)
	GetUploadById(ctx context.Context, uploadId int64) (models.Upload, error)
}

type CreateUploadParams struct {
	File        io.Reader
	Name        *string
	Description *string
	Trainer     *string
	Commute     *string
	DataType    *string
	ExternalId  *string
}

func newUploadsApiService(client *strava.StravaAPIV3, auth runtime.ClientAuthInfoWriter) UploadsAPI {
	return uploadsService{
		client: client,
		auth:   auth,
	}
}

type uploadsService struct {
	client *strava.StravaAPIV3
	auth   runtime.ClientAuthInfoWriter
}

func (u uploadsService) CreateUpload(ctx context.Context, opts ...CreateUploadParams) (models.Upload, error) {
	params := uploads.NewCreateUploadParams()

	params.SetDefaults()
	params.SetContext(ctx)

	if len(opts) > 1 {
		return models.Upload{}, ErrTooManyOptions
	}

	if len(opts) == 1 {
		opt := opts[0]
		if opt.File != nil {
			params.SetFile(runtime.NamedReader("file", opt.File))
		}
		if opt.Name != nil {
			params.SetName(opt.Name)
		}
		if opt.Description != nil {
			params.SetDescription(opt.Description)
		}
		if opt.Trainer != nil {
			params.SetTrainer(opt.Trainer)
		}
		if opt.Commute != nil {
			params.SetCommute(opt.Commute)
		}
		if opt.DataType != nil {
			params.SetDataType(opt.DataType)
		}
		if opt.ExternalId != nil {
			params.SetExternalID(opt.ExternalId)
		}
	}

	resp, err := u.client.Uploads.CreateUpload(params, u.auth)
	if err != nil {
		return models.Upload{}, err
	}

	return *resp.Payload, nil
}

func (u uploadsService) GetUploadById(ctx context.Context, uploadId int64) (models.Upload, error) {
	params := uploads.NewGetUploadByIDParams()

	params.SetDefaults()
	params.SetContext(ctx)

	params.SetUploadID(uploadId)

	resp, err := u.client.Uploads.GetUploadByID(params, u.auth)
	if err != nil {
		return models.Upload{}, err
	}

	return *resp.Payload, nil
}
