/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.notificator.yaml)")
	rootCmd.AddCommand(serveCmd)
}

var (
	cfgFile          string
	notificatorViper = viper.New()

	// rootCmd represents the base command when called without any subcommands
	rootCmd = &cobra.Command{
		Use:   "notificator",
		Short: "notificator notify user of the specified message.",
		Long: `Notificator is a server that
	notify user of the specified message.`,
	}
)

func Execute() error {
	return rootCmd.Execute()
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		notificatorViper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".notificator" (without extension).
		notificatorViper.AddConfigPath(home)
		notificatorViper.SetConfigType("yaml")
		notificatorViper.SetConfigName(".notificator")
	}

	notificatorViper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := notificatorViper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

const (
	LineChannelSecret      = "LINE_CHANNEL_SECRET"
	LineChannelAccessToken = "LINE_CHANNEL_ACCESS_TOKEN"
	LineUserID             = "LINE_USER_ID"
	MetricsAddr            = "METRICS_ADDR"
)

type NotificatorConfig struct {
	LineChannelSecret      string
	LineChannelAccessToken string
	LineUserID             string
	ObservabilityConfig
}

func getNotificatorConfig() NotificatorConfig {
	config := NotificatorConfig{
		LineChannelSecret:      notificatorViper.GetString(LineChannelSecret),
		LineChannelAccessToken: notificatorViper.GetString(LineChannelAccessToken),
		LineUserID:             notificatorViper.GetString(LineUserID),
		ObservabilityConfig:    getObservabilityConfig(),
	}
	return config
}

type ObservabilityConfig struct {
	MetricsAddr string
}

func getObservabilityConfig() ObservabilityConfig {
	return ObservabilityConfig{
		MetricsAddr: notificatorViper.GetString(MetricsAddr),
	}
}
