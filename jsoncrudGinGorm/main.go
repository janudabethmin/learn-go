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
	// Create a new Gin router
	router := gin.Default()
	// Define the routes for the application
	router.POST("/post", controllers.CreatePost)
	router.GET("/posts", controllers.GetAllPosts)
	router.GET("/post", controllers.GetPostById)
	router.PUT("/post", controllers.UpdatePost)
	router.DELETE("/post", controllers.DeletePost)
	// Router default port is 8080.
	// If the PORT ENV variable is set, it will be used instead
	err := router.Run()
	if err != nil {
		log.Fatal("Router is not working as expected.")
		return
	}
}
