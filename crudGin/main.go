package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	// Initialize the Gin router
	router := gin.Default()

	// Route a specific route to a function
	router.GET("/books", getAllBooks)
	router.POST("/books", createBook)
	router.GET("/books/:id", getBookById)
	router.PATCH("/books/checkout", checkoutBook)
	router.PATCH("/books/return", returnBook)

	// Start the server on localhost:8080
	err := router.Run("localhost:8080")
	if err != nil {
		return
	}
}

// *gin.Context is a pointer to the context of the request that contains all the information about the request, allowing us to create the response
func getAllBooks(c *gin.Context) {

	// Return the list of books as nicely indented JSON
	c.IndentedJSON(http.StatusOK, books)
}

func createBook(c *gin.Context) {

	// Create a new variable to hold the new book data
	var newBook Book

	err := c.BindJSON(&newBook)
	if err != nil {
		// gin.H is used to create a json object without creating a struct
		// If the JSON is invalid, return a 400 Bad Request
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid input"})
		return
	}

	// Add the new book to the in-memory database by
	books = append(books, newBook)

	// Return the newly created book with a 201 Created status
	c.IndentedJSON(http.StatusCreated, newBook)
}

func getBookById(c *gin.Context) {

	// Get the book ID from the URL via Path parameters
	id := c.Param("id")

	// Call the helper function to get the book by ID
	book, err := bookById(id)
	if err != nil {
		// If the book is not found, return a 404 Not Found
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}
	// Return the book as nicely indented JSON
	c.IndentedJSON(http.StatusOK, book)
}

// This is a helper function to get a book by its ID
func bookById(id string) (*Book, error) {

	// Loop through the books to find the book with the given ID
	for i, book := range books {
		if book.ID == id {
			return &books[i], nil
		}
	}

	// Uses the errors package to create a new error
	return nil, errors.New("Book not found")
}

func checkoutBook(c *gin.Context) {

	// Get the book ID from the URL via Query parameters
	id, ok := c.GetQuery("id")

	if !ok {
		// If the ID is not provided, return a 400 Bad Request
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "ID is required"})
		return
	}

	// Call the helper function to get the book by ID
	book, err := bookById(id)
	if err != nil {
		// If the book is not found, return a 404 Not Found
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}

	// Check if the book is available for checkout
	if book.Quantity <= 0 {
		// If the book is not available, return a 400 Bad Request
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Book not available"})
		return
	}

	// Decrease the quantity of the book by 1
	book.Quantity--

	// Return the updated book with a 200 OK status
	c.IndentedJSON(http.StatusOK, book)
}

func returnBook(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		// If the ID is not provided, return a 400 Bad Request
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "ID is required"})
		return
	}

	// Call the helper function to get the book by ID
	book, err := bookById(id)
	// If the book is not found, return a 404 Not Found
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}

	// Increase the quantity of the book by 1
	book.Quantity++

	// Return the updated book with a 200 OK status
	c.IndentedJSON(http.StatusOK, book)
}

type Book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

// In-memory database of books
var books = []Book{
	{ID: "1", Title: "Book One", Author: "Author A", Quantity: 5},
	{ID: "2", Title: "Book Two", Author: "Author B", Quantity: 3},
	{ID: "3", Title: "Book Three", Author: "Author C", Quantity: 2},
}
