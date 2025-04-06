package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	ServerPort string
}

func LoadConfig() *Config {
	_ = godotenv.Load(".env")
	return &Config{
		ServerPort: getEnv("SERVER_PORT", "8080"),
	} 
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}