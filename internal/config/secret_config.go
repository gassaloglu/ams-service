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

	log.Info().Str("file", secretConfigFile).Msg("Secret config loaded successfully")

	return &secretConfig, nil
}
