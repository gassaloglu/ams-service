package config

import (
	"errors"
	"os"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

var CONFIG_LOG_PREFIX = "config.go"

type Config struct {
	ServerPort string
	Database   DatabaseConfig
}

type DatabaseConfig struct {
	Type     string
	Host     string
	Port     int
	User     string
	Password string
	Name     string
	SSLMode  string
}

func LoadConfig() (*Config, error) {
	configFile, found := os.LookupEnv("CONFIG_FILE")

	if !found {
		return nil, errors.New("could not read environment variable CONFIG_FILE")
	}

	log.Info().Str("file", configFile).Msg("Loading configuration")

	viper.SetConfigFile(configFile)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	log.Info().Str("file", configFile).Msg("Configuration loaded successfully")

	return &cfg, nil
}
