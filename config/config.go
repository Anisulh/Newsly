package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config stores all configuration values required by the application.
type Config struct {
	Port               string
	Environment        string
	ClientAddress      string
	DBConnectionString string
	JWTSecret          string
	NewsAPIKey         string
}

// loadEnv loads environment variables from a .env file and system environment.
func loadEnv() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: No .env file found")
	}

	var config Config
	var err error

	if config.Port, err = getEnv("PORT"); err != nil {
		return nil, err
	}
	if config.Environment, err = getEnv("ENVIRONMENT"); err != nil {
		return nil, err
	}
	if config.ClientAddress, err = getEnv("CLIENT_ADDRESS"); err != nil {
		return nil, err
	}
	if config.DBConnectionString, err = getEnv("DATABASE_URL"); err != nil {
		return nil, err
	}
	if config.JWTSecret, err = getEnv("JWT_SECRET_KEY"); err != nil {
		return nil, err
	}
	if config.NewsAPIKey, err = getEnv("NEWS_API_KEY"); err != nil {
		return nil, err
	}

	return &config, nil
}

// getEnv retrieves the value of the environment variable named by the key.
// It returns an error if the variable is not set.
func getEnv(key string) (string, error) {
	value, exists := os.LookupEnv(key)
	if !exists {
		return "", fmt.Errorf("environment variable not found: %s", key)
	}
	return value, nil
}

// MustLoadConfig loads the application configuration and panics if any required
// environment variables are missing.
func MustLoadConfig() *Config {
	cfg, err := loadEnv()
	if err != nil {
		panic(fmt.Sprintf("failed to load configuration: %v", err))
	}
	return cfg
}
