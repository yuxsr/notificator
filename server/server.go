package server

import (
	"github.com/yuxsr/notificator/server/service"
	goproto "github.com/yuxsr/yuxsr-dev-pb/gencode/go_proto"
	"google.golang.org/grpc"
)

func RegisterNewGRPCServer(config service.NotificatorServiceConfig) *grpc.Server {
	notificatorService := service.NewNotificatorService(config)
	s := grpc.NewServer()
	goproto.RegisterNotificatorServiceServer(s, notificatorService)
	return s
}
