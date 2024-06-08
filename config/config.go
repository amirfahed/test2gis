package config

import (
	"applicationDesignTest/internal/utils/logger"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

var AppConfig Config

type ServerConfig struct{
	Host string
	Port string

}
type Config struct {
	Server ServerConfig
	DB DBConfig
}

type DBConfig struct {
    Host     string
    Port     int
    User     string
    Password string
    DBName   string
}

func New() (*Config, error) {
	ex, err := os.Executable()
    if err != nil {
        panic(err)
    }
    exPath := filepath.Dir(ex)
    fmt.Println(exPath)
	
	viper.SetConfigFile("../config/config.yaml")
	if err := viper.ReadInConfig(); err != nil {
	}

	err = viper.Unmarshal(&AppConfig)
    if err != nil {
        logger.LogErrorf("Unable to decode into struct %v", err)
    }

	return &AppConfig, nil
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
