package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	ServerPort string
	Database   DatabaseConfig
	Firebase   FirebaseConfig
}

type DatabaseConfig struct {
	Type     string
	Host     string
	Port     int
	User     string
	Password string
	Name     string
	URI      string
}

type FirebaseConfig struct {
	CredentialsFile string
}

func LoadConfig() (*Config, error) {
	viper.SetConfigFile("C:/DEV/db-config/local-config.yaml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
