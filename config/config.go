package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	LogLevel    string
	DatabaseURL string
}

func LoadConfig() (*Config, error) {

	err := godotenv.Load()
	if err != nil && !os.IsNotExist(err) {
		log.Printf("Error loading .env file: %v", err)
	}

	dsn := fmt.Sprintf(
		"host=db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=UTC",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	config := &Config{
		LogLevel:    os.Getenv("LOG_LEVEL"),
		DatabaseURL: dsn,
	}

	if config.LogLevel == "" {
		config.LogLevel = "info"
	}

	return config, nil
}
