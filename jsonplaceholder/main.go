package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	url := "https://jsonplaceholder.typicode.com/posts"

	response, error := http.Get(url)

	if error != nil {
		fmt.Println("Error fetching data:", error)
		return
	}

	// Ensure the response body is closed after all operations
	defer response.Body.Close()

	// Check if the response status code is 200 OK
	if response.StatusCode != http.StatusOK {
		fmt.Println("Error: Status code", response.StatusCode)
		return
	}

	// Get the pointer to the response body
	responseBodyPointer := response.Body

	// Use io.ReadAll to read the response body
	responseBodyByteType, errorBody := io.ReadAll(responseBodyPointer)

	if errorBody != nil {
		fmt.Println("Error reading response body:", errorBody)
		return
	}

	responseBody := string(responseBodyByteType)

	fmt.Println("Response Body:", responseBody)
}
