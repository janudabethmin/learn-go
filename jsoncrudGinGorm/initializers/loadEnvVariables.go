package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

// LoadEnvVariables Loading environment variables from the .env file
func LoadEnvVariables() {
	// Loading environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}
}
