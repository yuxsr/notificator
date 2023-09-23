package main

import (
	"log/slog"
	"os"

	"github.com/yuxsr/notificator/cmd"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	rootCmd := cmd.NewRootCmd()
	if err := rootCmd.Execute(); err != nil {
		logger.Error("command execution error", "error", err)
	}
}
