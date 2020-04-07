package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	// read config from [Project]/config/application.yml
	viper.AddConfigPath("./config")
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	if err := viper.ReadInConfig(); err != nil {
		logrus.Fatal(err)
	}
}
