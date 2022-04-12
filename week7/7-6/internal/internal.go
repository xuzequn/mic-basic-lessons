package internal

import (
	"github.com/spf13/viper"
	"mic-basic-lessons/week7/7-6/conf"
)

var ServerConfig config.ServerConfig

func init() {
	v := viper.New()
	v.SetConfigFile("week7/7-6/conf/dev-config.yaml")
	v.ReadInConfig()
	v.Unmarshal(&ServerConfig)
}
