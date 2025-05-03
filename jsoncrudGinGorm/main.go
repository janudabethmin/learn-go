package main

import (
	"github.com/gin-gonic/gin"
	"github.com/janudabethmin/learn-go/jsoncrudGinGorm/controllers"
	"github.com/janudabethmin/learn-go/jsoncrudGinGorm/initializers"
	"log"
)

// init() runs before main() and is used to initialize the application
func init() {
	// Running the LoadEnvVariables and ConnectDatabase functions from the initializers package
	// Need to start these functions with a capital letter to make it public (can be used in other packages/files)
	initializers.LoadEnvVariables()
	initializers.ConnectDatabase()
}

func main() {
	router := gin.Default()
	router.POST("/posts", controllers.CreatePost)
	router.GET("/posts", controllers.GetAllPosts)
	// Router default port is 8080.
	// If the PORT ENV variable is set, it will be used instead
	err := router.Run()
	if err != nil {
		log.Fatal("Router is not working as expected.")
		return
	}
}

func hello(c *gin.Context) {
	c.IndentedJSON(200, gin.H{"message": "Hello World"})
}
