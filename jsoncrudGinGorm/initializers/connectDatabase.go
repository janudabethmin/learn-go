package initializers

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

// DB Database Connection Global Variable
var DB *gorm.DB

func init() {
	// No need initializers.LoadEnvVariables because it's in the same package
	LoadEnvVariables()
}

// ConnectDatabase Connecting to the PostgreSQL database
func ConnectDatabase() {
	var err error
	dsn := os.Getenv("DB_STRING")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
		return
	}
}
