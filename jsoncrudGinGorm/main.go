package main

import (
	"github.com/gin-gonic/gin"
	"github.com/janudabethmin/learn-go/jsoncrudGinGorm/initializers"
)

// init() runs before main() and is used to initialize the application
func init(){
	// Running the LoadEnvVariables function from the initializers package
	initializers.LoadEnvVariables()
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