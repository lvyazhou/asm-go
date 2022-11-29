package config

import (
	"fmt"
	"log"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var config *viper.Viper

func Init(env string) {
	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		log.Panicf("error when finding home directory: %s\n", err)
	}
	homeConfigPath := fmt.Sprintf("%s/.360-osint-web/", home)
	currentConfigPath := "infrastructure/config/"

	config = viper.New()
	config.SetConfigType("yaml")
	config.SetConfigName(env)
	config.AddConfigPath(currentConfigPath)
	config.AddConfigPath(homeConfigPath)
	if err := config.ReadInConfig(); err != nil {
		log.Panicf("error on parsing configuration file: %s\n", err)
	}
}

func GetConfig() *viper.Viper {
	return config
}
