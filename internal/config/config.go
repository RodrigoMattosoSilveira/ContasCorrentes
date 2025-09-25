package config

import (
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
)

type Config struct {
	Env        string
	Port       string
	DSN        string
	CSRFSecret string
	SessionKey string
}

func LoadConfig() (*Config, error) {
	if os.Getenv("APP_ENV") != "production" {
		if err := godotenv.Load(); err != nil {
			log.Println("Error loading .env file")
		}
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		getEnv("DB_HOST", "localhost"),
		getEnv("DB_USER", "sqlite"),
		getEnv("DB_PASSWORD", "password"),
		getEnv("DB_NAME", "/private/var/ContasCorrentes/sqlite_dev.db"),
		getEnv("DB_PORT", "5432"),
		getEnv("DB_SSLMODE", "disable"),
	)

	cfg := &Config{
		Env:        getEnv("APP_ENV", "development"),
		Port:       getEnv("APP_PORT", "8080"),
		DSN:        dsn,
		CSRFSecret: getEnv("CSRF_SECRET", "default-secret-must-be-32-chars-long"),
		SessionKey: getEnv("SESSION_KEY", "default-session-key-change-me"),
	}

	return cfg, nil
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
