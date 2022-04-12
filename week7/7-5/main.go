package main

import (
	"fmt"
	"github.com/spf13/viper"
)

type RedisConfig struct {
	RedisHost string `mapstructure:"redisHost"`
	RedisPort int32  `mapstructure:"redisPort"`
}

type ServerConfig struct {
	ServerName  string      `mapstructure:"serverName"`
	RedisConfig RedisConfig `mapstructuure:"redisConfig"`
}

func GetEnv(s string) int {
	viper.AutomaticEnv() // 预先加载环境变量
	return viper.GetInt(s)

}

func main() {
	v := viper.New()
	dev := GetEnv("DEV")
	// export DEV=1
	configFilePath := "week7/7-5/dev-config.yaml"
	if dev == 0 {
		configFilePath = "week7/7-5/pro-config.yaml"
	}
	v.SetConfigFile(configFilePath)
	v.ReadInConfig()
	var serverConfig ServerConfig
	err := v.Unmarshal(&serverConfig)
	if err != nil {
		panic(err)
	}
	fmt.Println(serverConfig.ServerName)
	fmt.Println(serverConfig.RedisConfig.RedisHost)
	fmt.Println(serverConfig.RedisConfig.RedisPort)

}
