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

func GetPostById(c *gin.Context) {

	// Get the post ID from the query parameters
	id, ok := c.GetQuery("id")
	if !ok {
		c.JSON(400, gin.H{"error": "Post ID is required"})
		return
	}

	// Find the post by ID
	var post models.Post
	result := initializers.DB.First(&post, id)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Post not found"})
		return
	}

	// Return the post as a JSON response
	c.JSON(200, post)
}

func UpdatePost(c *gin.Context) {
	// Get the post ID from the query parameters
	id, ok := c.GetQuery("id")
	if !ok {
		c.JSON(400, gin.H{"error": "Post ID is required"})
		return
	}

	// Get the updated data from the request body
	var body struct {
		Message string `json:"message"`
		Title   string `json:"title"`
	}
	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	// Find the post by ID and update it
	var post models.Post
	result := initializers.DB.First(&post, id)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Post not found"})
		return
	}

	post.Title = body.Title
	post.Message = body.Message

	// Save the updated post to the database
	result = initializers.DB.Save(&post)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to update post"})
		return
	}

	// Return the updated post as a JSON response
	c.JSON(200, post)
}

func DeletePost(c *gin.Context) {
	// Get the post ID from the query parameters
	id, ok := c.GetQuery("id")
	if !ok {
		c.JSON(400, gin.H{"error": "Post ID is required"})
		return
	}

	// Find the post by ID and delete it
	var post models.Post
	result := initializers.DB.First(&post, id)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Post not found"})
		return
	}

	result = initializers.DB.Delete(&post)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to delete post"})
		return
	}

	// Return a success message as a JSON response
	c.JSON(200, gin.H{"message": "Post deleted successfully"})
}
