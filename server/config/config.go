package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUser     string
	DBPassword string
	DBName     string
	DBHost     string
	DBPort     string
}

func LoadConfig() (*Config, error) {
	// .envファイルを読み込みます
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	config := &Config{
		DBUser:     getEnv("DB_USER", "user"),
		DBPassword: getEnv("DB_PASSWORD", "password"),
		DBName:     getEnv("DB_NAME", "todos"),
		DBHost:     getEnv("DB_HOST", "db"),
		DBPort:     getEnv("DB_PORT", "5432"),
	}

	return config, nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
