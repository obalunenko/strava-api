package client

import (
	"context"
	"github.com/go-openapi/runtime"
	strava "github.com/obalunenko/strava-api/internal/gen/strava-api-go/client"
	"github.com/obalunenko/strava-api/internal/gen/strava-api-go/models"
	"io"
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
	//TODO implement me
	panic("implement me")
}

func (u uploadsService) GetUploadById(ctx context.Context, uploadId int64) (models.Upload, error) {
	//TODO implement me
	panic("implement me")
}
