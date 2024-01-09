package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	ClientAddress      string
	DBConnectionString string
	JWTSecret          string
	NewsAPIKey         string
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	ClientAddress = getEnv("CLIENT_ADDRESS", "default-client-address")
	DBConnectionString = getEnv("DATABASE_URL", "")
	JWTSecret = getEnv("JWT_SECRET_KEY", "")
	NewsAPIKey = getEnv("NEWS_API_KEY", "")

	// Example of mandatory env variable checking
	if DBConnectionString == "" {
		log.Fatal("DATABASE_URL is a required environment variable")
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
