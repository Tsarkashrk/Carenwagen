package config

import (
	"log"
	"os"
)

type Config struct {
	DBUser     string
	DBPassword string
	DBName     string
	DBHost     string
	DBPort     string
	NatsURL    string
}

func LoadConfig() Config {
	cfg := Config{
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		NatsURL:    os.Getenv("NATS_URL"),
	}

	// Log loaded configuration for debugging
	log.Printf("Loaded configuration: %+v", cfg)

	return cfg
}
