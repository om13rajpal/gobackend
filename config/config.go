package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/om13rajpal/gobackend/utils"
)

var (
	PORT       string
	MONGO_URI  string
	JWT_SECRET string
)

func LoadConfig() {
	envError := godotenv.Load(".env")
	utils.HandleError(envError, ".env not loaded")

	PORT = getEnv("PORT", "3000")
	MONGO_URI = getEnv("MONGO_URI", "mongodb://localhost:27017/")
	JWT_SECRET = getEnv("JWT_SECRET", "golang")
}

func getEnv(key string, fallback string) string {
	value, exists := os.LookupEnv(key)

	if exists {
		return value
	}

	return fallback
}