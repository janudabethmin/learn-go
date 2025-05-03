package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/janudabethmin/learn-go/jsoncrudGinGorm/initializers"
	"github.com/janudabethmin/learn-go/jsoncrudGinGorm/models"
)

func CreatePost(c *gin.Context) {

	// Create a struct that matches the request body
	var body struct {
		Message string `json:"message"`
		Title   string `json:"title"`
	}
	// Bind the request body to the struct
	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	// Create a new post using the data got from the request body
	post := models.Post{
		Title:   body.Title,
		Message: body.Message,
	}
	// Save the post to the database
	result := initializers.DB.Create(&post)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to create post"})
		return
	}

	// Return the created post as a JSON response
	c.JSON(200, post)
}

func GetAllPosts(c *gin.Context) {

	// Get all posts from the database
	var posts []models.Post // Array of posts
	result := initializers.DB.Find(&posts)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to get posts"})
		return
	}

	// Return the posts as a JSON response
	c.JSON(200, posts)
}
