package utils

import "github.com/spf13/viper"

func InitViper(confPath string) error {

	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(confPath)

	return viper.ReadInConfig()
}
