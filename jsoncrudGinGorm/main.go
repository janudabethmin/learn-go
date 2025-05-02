package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// init() runs before main() and is used to initialize the application
func init(){
	// Loading environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	router := gin.Default()
	router.GET("/hello", hello)
	// Router default port is 8080. 
	// If the PORT ENV variable is set, it will be used instead
	router.Run()
}

func hello(c *gin.Context) {
	c.IndentedJSON(200, gin.H{"message": "Hello World"})
}