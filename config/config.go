package config

import (
	"github.com/spf13/viper"
)

type MongoConfiguration struct {
	Server     string
	Database   string
	Collection string
}

type Configuration struct {
	Environment string
	Mongo       MongoConfiguration
}

func GetConfig() Configuration {
	conf := Configuration{}

	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")

	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&conf)
	if err != nil {
		panic(err)
	}

	return conf
}
