// Package config is application config
package config

import (
	"strings"

	"github.com/spf13/viper"
)

const (
	LineChannelSecret      = "LINE_CHANNEL_SECRET"
	LineChannelAccessToken = "LINE_CHANNEL_ACCESS_TOKEN"
	LineUserID             = "LINE_USER_ID"
)

var serverViper = viper.New()

func init() {
	serverViper.AutomaticEnv()
	serverViper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
}

type ServerConfig struct {
	LineChannelSecret      string
	LineChannelAccessToken string
	LineUserID             string
}

func GetServerConfig() ServerConfig {
	config := ServerConfig{
		LineChannelSecret:      serverViper.GetString(LineChannelSecret),
		LineChannelAccessToken: serverViper.GetString(LineChannelAccessToken),
		LineUserID:             serverViper.GetString(LineUserID),
	}
	return config
}
