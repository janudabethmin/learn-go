package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")

	user := User{
		Name:     "Januda Bethmin",
		Age:      24,
		Password: "securepassword",
	}
	fmt.Println("User:", user)

	userConvertedTOByteString, error := json.Marshal(user)

	if error != nil {
		fmt.Println("Error converting user to JSON:", error)
		return
	}

	userConvertedTOJson := string(userConvertedTOByteString)
	fmt.Println("User in JSON format:", userConvertedTOJson)
}

type User struct {
	Name     string `json:"name,omitempty"`
	Age      int    `json:"age"`
	Password string `json:"password"`
}
