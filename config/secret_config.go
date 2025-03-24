package config

import (
	"ams-service/middlewares"
	"fmt"
	"runtime"

	"github.com/spf13/viper"
)

var SECRET_CONFIG_LOG_PREFIX = "secret_config.go"

type SecretConfig struct {
	JWTSecretKey string `mapstructure:"jwt_secret_key"`
}

var JWTSecretKey string

func init() {
	var secretConfigFile string
	if runtime.GOOS == "windows" {
		secretConfigFile = "C:/DEV/secret-config/secret-config.yaml"
	} else {
		secretConfigFile = "/path/to/db-config/local-config.yaml"
	}

	viper.SetConfigFile(secretConfigFile)

	if err := viper.ReadInConfig(); err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error reading secret config file: %v", SECRET_CONFIG_LOG_PREFIX, err))
		return
	}

	var secretConfig SecretConfig
	if err := viper.Unmarshal(&secretConfig); err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error unmarshalling secret config file: %v", SECRET_CONFIG_LOG_PREFIX, err))
		return
	}

	JWTSecretKey = secretConfig.JWTSecretKey
	middlewares.LogInfo(fmt.Sprintf("%s - JWT Secret Key loaded successfully", SECRET_CONFIG_LOG_PREFIX))
}
