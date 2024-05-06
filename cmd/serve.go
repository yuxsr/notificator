/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"net"

	"github.com/spf13/cobra"
	"github.com/yuxsr/notificator/internal/client"
	"github.com/yuxsr/notificator/internal/observability"
	"github.com/yuxsr/notificator/internal/server"
	"github.com/yuxsr/notificator/internal/server/service"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "serve notificator server",
	Long:  `Serve notificator server.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		config := getNotificatorConfig()
		return Serve(config)
	},
}

// Serve is run server
func Serve(config NotificatorConfig) error {
	ctx := context.Background()
	clinet, err := client.NewLineClient(ctx, client.LineClientConfig{
		ChannelSecret:      config.LineChannelSecret,
		ChannelAccessToken: config.LineChannelAccessToken,
		UserID:             config.LineUserID,
	})
	if err != nil {
		return err
	}

	server := server.RegisterNewGRPCServer(service.NotificatorServiceConfig{
		Client: clinet,
	})

	shutdown := observability.InitMetrics(config.MetricsAddr)
	defer func() {
		_ = shutdown()
	}()

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		return err
	}
	if err := server.Serve(lis); err != nil {
		return err
	}
	return nil
}
