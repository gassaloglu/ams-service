package config

import (
	"errors"
	"os"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

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
		log.Error().Err(err).Msg("Error reading secret config file")
		return nil, err
	}

	var secretConfig SecretConfig
	if err := viper.Unmarshal(&secretConfig); err != nil {
		log.Error().Err(err).Msg("Error unmarshalling secret config file")
		return nil, err
	}

	JWTSecretKey = secretConfig.JWTSecretKey

	log.Info().Msg("JWT Secret Key loaded successfully")

	return &secretConfig, nil
}
