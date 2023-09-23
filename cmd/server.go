package cmd

import (
	"context"
	"net"

	"github.com/spf13/cobra"
	"github.com/yuxsr/notificator/client"
	"github.com/yuxsr/notificator/config"
	"github.com/yuxsr/notificator/server"
	"github.com/yuxsr/notificator/server/service"
)

// NewServeCmd is create sub command `server` instance.
func NewServeCmd() *cobra.Command {
	return &cobra.Command{
		Use:           "serve",
		Short:         "Serve notificator",
		Long:          ``,
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			config := config.GetServerConfig()
			return Serve(config)
		},
	}
}

// Serve is run server
func Serve(config config.ServerConfig) error {
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

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		return err
	}
	if err := server.Serve(lis); err != nil {
		return err
	}
	return nil
}
