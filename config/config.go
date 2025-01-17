package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type Config struct {
	ServerPort string `yaml:"server_port"`
	Database   struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
	} `yaml:"database"`
}

func LoadConfig() (*Config, error) {
	data, err := ioutil.ReadFile("config/config.yaml")
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
