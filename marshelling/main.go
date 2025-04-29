package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")

	user1 := User{
		Name:     "Januda",
		Age:      24,
		Password: "securepassword",
	}
	fmt.Println("User1 :", user1)

	// Marshalling the user struct to JSON
	userConvertedTOByteString, error1 := json.Marshal(user1)

	if error1 != nil {
		fmt.Println("Error converting user1 to JSON:", error1)
		return
	}

	userConvertedTOJson := string(userConvertedTOByteString)
	fmt.Println("User1 in JSON format:", userConvertedTOJson)

	// Unmarshalling the JSON to a struct
	jsonUser := `{"name":"Bethmin","age":24,"password":"mysecurepassword"}`

	user2 := User{}

	error2 := json.Unmarshal([]byte(jsonUser), &user2)

	if error2 != nil {
		fmt.Println("Error converting JSON to user2:", error2)
		return
	}

	fmt.Println("User2 :", user2)
}

type User struct {
	Name     string `json:"name,omitempty"`
	Age      int    `json:"age"`
	Password string `json:"password"`
}
