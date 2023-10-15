package server

import (
	"github.com/yuxsr/notificator/internal/server/service"
	protov1 "github.com/yuxsr/yuxsr-dev-pb/gencode/go/proto/v1"
	"google.golang.org/grpc"
)

func RegisterNewGRPCServer(config service.NotificatorServiceConfig) *grpc.Server {
	notificatorService := service.NewNotificatorService(config)
	s := grpc.NewServer()
	protov1.RegisterNotificatorServiceServer(s, notificatorService)
	return s
}
