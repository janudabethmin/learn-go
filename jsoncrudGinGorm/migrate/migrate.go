package main

import (
	"github.com/janudabethmin/learn-go/jsoncrudGinGorm/initializers"
	"github.com/janudabethmin/learn-go/jsoncrudGinGorm/models"
	"log"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDatabase()
}

func main() {
	// Migrate the database schema using models
	// This will create the posts table in the database
	err := initializers.DB.AutoMigrate(&models.Post{})
	if err != nil {
		log.Fatal("Failed to migrate database")
		return
	}
}
