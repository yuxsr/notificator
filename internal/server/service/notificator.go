// Package service is application logic
package service

import (
	"context"

	goproto "github.com/yuxsr/yuxsr-dev-pb/gencode/go_proto"
)

type NotificatorServiceConfig struct {
	Client Client
}

type notificatorService struct {
	client Client
}

func NewNotificatorService(config NotificatorServiceConfig) goproto.NotificatorServiceServer {
	return &notificatorService{
		client: config.Client,
	}
}

func (n *notificatorService) Notify(ctx context.Context, req *goproto.NotifyRequest) (*goproto.NotifyResponse, error) {
	if err := n.client.Notify(ctx, req.Message); err != nil {
		return nil, err
	}
	return &goproto.NotifyResponse{}, nil
}
