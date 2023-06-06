package storage

import (
	"CAR24/Car24_client_service/genproto/client_service"
	"CAR24/Car24_client_service/models"
	"context"
)

type StorageI interface {
	CloseDB()
	Client() ClientRepoI
}

type ClientRepoI interface {
	Create(ctx context.Context, req *client_service.CreateClientRequest) (resp *client_service.ClientPrimaryKey, err error)
	GetById(ctx context.Context, req *client_service.ClientPrimaryKey) (client *client_service.Client, err error)
	Update(ctx context.Context, req *client_service.UpdateClient) (rowsaffected int64, err error)
	GetList(ctx context.Context, req *client_service.GetListClientRequest) (clients *client_service.GetListClientResponse, err error)
	Delete(ctx context.Context, req *client_service.ClientPrimaryKey) (err error)
	UpdatePatch(ctx context.Context, req models.UpdatePatchRequest) (rowsaffected int64, err error)
}
