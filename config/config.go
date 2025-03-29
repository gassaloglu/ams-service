package config

import (
	"ams-service/middlewares"
	"errors"
	"fmt"
	"os"

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

	middlewares.LogInfo(fmt.Sprintf("%s - Loading configuration from: %s", CONFIG_LOG_PREFIX, configFile))

	viper.SetConfigFile(configFile)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	middlewares.LogInfo(fmt.Sprintf("%s - Configuration loaded successfully: %s", CONFIG_LOG_PREFIX, configFile))

	return &cfg, nil
}
