package service

import (
	"CAR24/Car24_client_service/config"
	"CAR24/Car24_client_service/genproto/client_service"
	"CAR24/Car24_client_service/grpc/client"
	"CAR24/Car24_client_service/models"
	"CAR24/Car24_client_service/pkg/logger"
	"CAR24/Car24_client_service/storage"
	"context"
	"fmt"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ClientService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	*client_service.UnimplementedClientServiceServer
}

func NewClientService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *ClientService {
	return &ClientService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (i *ClientService) Create(ctx context.Context, req *client_service.CreateClientRequest) (resp *client_service.Client, err error) {

	i.log.Info("---CreateClient------>", logger.Any("req", req))

	pKey, err := i.strg.Client().Create(ctx, req)
	if err != nil {
		i.log.Error("!!!CreateClient->Client->Create--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp, err = i.strg.Client().GetById(ctx, &client_service.ClientPrimaryKey{Id: pKey.Id})
	if err != nil {
		i.log.Error("!!!GetByPKeyClient->Client->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *ClientService) GetById(ctx context.Context, req *client_service.ClientPrimaryKey) (resp *client_service.Client, err error) {

	i.log.Info("---GetClientByID------>", logger.Any("req", req))

	fmt.Println("I am here")
	resp, err = i.strg.Client().GetById(ctx, req)
	if err != nil {
		i.log.Error("!!!GetClientByID->Client->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *ClientService) GetList(ctx context.Context, req *client_service.GetListClientRequest) (resp *client_service.GetListClientResponse, err error) {

	i.log.Info("---GetClients------>", logger.Any("req", req))

	resp, err = i.strg.Client().GetList(ctx, req)
	if err != nil {
		i.log.Error("!!!GetClients->Client->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *ClientService) Update(ctx context.Context, req *client_service.UpdateClient) (resp *client_service.Client, err error) {

	i.log.Info("---UpdateClient------>", logger.Any("req", req))

	rowsAffected, err := i.strg.Client().Update(ctx, req)

	if err != nil {
		i.log.Error("!!!UpdateClient--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = i.strg.Client().GetById(ctx, &client_service.ClientPrimaryKey{Id: req.Id})
	if err != nil {
		i.log.Error("!!!GetClient->Client->Get--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *ClientService) UpdatePatch(ctx context.Context, req *client_service.UpdatePatchClient) (resp *client_service.Client, err error) {

	i.log.Info("---UpdatePatchClient------>", logger.Any("req", req))

	updatePatchModel := models.UpdatePatchRequest{
		Id:     req.GetId(),
		Fields: req.GetFields().AsMap(),
	}

	rowsAffected, err := i.strg.Client().UpdatePatch(ctx, updatePatchModel)

	if err != nil {
		i.log.Error("!!!UpdatePatchUser--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = i.strg.Client().GetById(ctx, &client_service.ClientPrimaryKey{Id: req.Id})
	if err != nil {
		i.log.Error("!!!GetClient->Client->Get--->", logger.Error(err))

		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *ClientService) Delete(ctx context.Context, req *client_service.ClientPrimaryKey) (resp *empty.Empty, err error) {

	i.log.Info("---DeleteUser------>", logger.Any("req", req))

	err = i.strg.Client().Delete(ctx, req)
	if err != nil {
		i.log.Error("!!!DeleteClient->Client->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &empty.Empty{}, nil
}
