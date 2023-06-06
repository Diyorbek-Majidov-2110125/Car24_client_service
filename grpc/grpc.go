package grpc

import (
	"CAR24/Car24_client_service/config"
	"CAR24/Car24_client_service/genproto/client_service"
	"CAR24/Car24_client_service/grpc/client"
	"CAR24/Car24_client_service/grpc/service"
	"CAR24/Car24_client_service/pkg/logger"
	"CAR24/Car24_client_service/storage"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func SetUpServer(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvc client.ServiceManagerI) (grpcServer *grpc.Server) {

	grpcServer = grpc.NewServer()

	client_service.RegisterClientServiceServer(grpcServer, service.NewClientService(cfg, log, strg, srvc))

	reflection.Register(grpcServer)
	return
}
