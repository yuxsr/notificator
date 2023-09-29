package main

import (
	"github.com/yuxsr/notificator/cmd"
	"github.com/yuxsr/notificator/logging"
)

func main() {
	logger := logging.NewLogger()
	if err := cmd.NewRootCmd().Execute(); err != nil {
		logger.Error("command execution error", "error", err)
	}
}
