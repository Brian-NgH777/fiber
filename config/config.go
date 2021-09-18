package config

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"strings"
)

const (
	defaultEnv = "dev"
	configType = "ini"
	configFilePath = "."
	production = "production"
	staging = "staging"
	development = "development"
)

func init() {
	Load()
}

func Load() error {
	env := strings.ToLower(os.Getenv("env"))
	if env == "" || (env != production && env != staging && env != development) {
		env = defaultEnv
	}

	viper.AddConfigPath(configFilePath)
	viper.SetConfigName(fmt.Sprintf("%s.%s","config", env))
	viper.SetConfigType(configType)

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Error(fmt.Sprintf("Config file not found; ignore error if desired, err: %v", err))
		} else {
			log.Error(fmt.Sprintf("Config file was found but another error was produced, err: %v", err))
		}
		return err
	}

	return err
}

func GetInt64(key string) int64 {
	v := viper.GetViper()
	value := v.GetInt64(key)
	return value
}

func GetInt(key string) int {
	v := viper.GetViper()
	value := v.GetInt(key)
	return value
}

func GetString(key string) string {
	v := viper.GetViper()
	value := v.GetString(key)
	return value
}
