package config

import "os"

type Config struct {
	Host string
	Port string
}

func New() (*Config, error) {
	return &Config{
		Host: getEnv("HOST", "localhost"),
		Port: getEnv("PORT", "8080"),
	}, nil
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
