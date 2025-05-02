package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

// LoadEnvVariables Need to start the function with a capital letter to make it public (can be used in other packages / files)
func LoadEnvVariables() {
	// Loading environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
