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
	JWTKey	 	string
}

func LoadConfig() (*Config, error) {

   err := godotenv.Load(".env")
    if err != nil {
        log.Fatal("Error loading .env file")
    }

   err = godotenv.Load(".env.secrets")
    if err != nil {
        log.Fatal("Error loading .env.secrets file")
    }

	dsn := fmt.Sprintf("dbname=%s", getEnv("DB_NAME", "/private/var/ContasCorrentes/sqlite_dev.db"))

	cfg := &Config{
		Env:        getEnv("APP_ENV", "development"),
		Port:       getEnv("APP_PORT", "8080"),
		DSN:        dsn,
		CSRFSecret: getEnv("CSRF_SECRET", "default-secret-must-be-32-chars-long"),
		SessionKey: getEnv("SESSION_KEY", "default-secret-must-be-32-chars-long"),
		JWTKey:     getEnv("JWT_KEY",     "default-secret-must-be-32-chars-long"),
	}

	return cfg, nil
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		os.Setenv(key, value)
		return value
	}
	os.Setenv(key, fallback)
	return fallback
}