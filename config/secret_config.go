package config

import (
	"ams-service/middlewares"
	"errors"
	"fmt"
	"os"

	"github.com/spf13/viper"
)

var SECRET_CONFIG_LOG_PREFIX = "secret_config.go"

type SecretConfig struct {
	JWTSecretKey string `mapstructure:"jwt_secret_key"`
}

var JWTSecretKey string

func LoadSecretConfig() (*SecretConfig, error) {
	secretConfigFile, found := os.LookupEnv("SECRET_FILE")

	if !found {
		return nil, errors.New("could not read environment variable SECRET_FILE")
	}

	viper.SetConfigFile(secretConfigFile)

	if err := viper.ReadInConfig(); err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error reading secret config file: %v", SECRET_CONFIG_LOG_PREFIX, err))
		return nil, err
	}

	var secretConfig SecretConfig
	if err := viper.Unmarshal(&secretConfig); err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error unmarshalling secret config file: %v", SECRET_CONFIG_LOG_PREFIX, err))
		return nil, err
	}

	JWTSecretKey = secretConfig.JWTSecretKey
	middlewares.LogInfo(fmt.Sprintf("%s - JWT Secret Key loaded successfully", SECRET_CONFIG_LOG_PREFIX))

	return &secretConfig, nil
}
