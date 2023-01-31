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
	homeConfigPath := fmt.Sprintf("%s/.asm-web/", home)
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

// S3Config s3 配置信息
type S3Config struct {
	S3Bucket    string `ini:"s3_bucket"`
	S3AccessKey string `ini:"s3_access_key"`
	S3SecretKey string `ini:"s3_secret_key"`
	S3Region    string `ini:"s3_region"`
	S3EndPoint  string `ini:"s3_end_point"`
}
