// Package service is application logic
package service

import (
	"context"

	protov1 "github.com/yuxsr/yuxsr-dev-pb/gencode/go/proto/v1"
)

type NotificatorServiceConfig struct {
	Client Client
}

type notificatorService struct {
	client Client
}

func NewNotificatorService(config NotificatorServiceConfig) protov1.NotificatorServiceServer {
	return &notificatorService{
		client: config.Client,
	}
}

func (n *notificatorService) Notify(ctx context.Context, req *protov1.NotifyRequest) (*protov1.NotifyResponse, error) {
	if err := n.client.Notify(ctx, req.Message); err != nil {
		return nil, err
	}
	return &protov1.NotifyResponse{}, nil
}
