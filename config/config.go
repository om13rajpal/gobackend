package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	PORT       string
	MONGO_URI  string
	JWT_SECRET string
	EMAIL      string
	PASSWORD   string
	POSTGRES_URI string
)

func LoadConfig() {
	envError := godotenv.Load(".env")
	if envError != nil {
		fmt.Println("could not read .env")
	}

	PORT = getEnv("PORT", "3000")
	MONGO_URI = getEnv("MONGO_URI", "mongodb://localhost:27017/")
	JWT_SECRET = getEnv("JWT_SECRET", "golang")
	EMAIL = getEnv("EMAIL", "")
	PASSWORD = getEnv("PASSWORD", "")
	POSTGRES_URI = getEnv("POSTGRES_URI", "")
}

func getEnv(key string, fallback string) string {
	value, exists := os.LookupEnv(key)

	if exists {
		return value
	}

	return fallback
}
