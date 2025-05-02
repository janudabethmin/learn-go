package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.GET("/hello", hello)
	router.Run("localhost:8080")
}

func hello(c *gin.Context) {
	c.IndentedJSON(200, gin.H{"message": "Hello World"})
}