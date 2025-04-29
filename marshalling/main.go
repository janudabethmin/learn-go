package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	fmt.Println("Hello, World!")

	// Marshalling the user struct to JSON

	user1 := User{
		Name:     "Januda",
		Age:      24,
		Password: "securepassword",
	}
	fmt.Println("User1 :", user1)

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

	user2Pointer := &user2

	jsonUserConvertedTOByteString := []byte(jsonUser)

	error2 := json.Unmarshal(jsonUserConvertedTOByteString, user2Pointer)

	if error2 != nil {
		fmt.Println("Error converting JSON to user2:", error2)
		return
	}

	fmt.Println("User2 :", user2)

	// Omitting empty fields

	user3 := UserWithOmitEmpty{
		Name:     "",
		Age:      24,
		Password: "mysecurepassword",
	}

	fmt.Println("User3 :", user3)

	// Removing a field from the JSON output

}

type User struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Password string `json:"password"`
}

type UserWithOmitEmpty struct {
	// omitempty omits the field from the JSON output if it is empty
	Name     string `json:"name,omitempty"`
	Age      int    `json:"age"`
	Password string `json:"password"`
}

type UserWhenPasswordIsNotSharedViaJson struct {
	// "-" omits the field from the JSON output no matter what
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Password string `json:"-"`
}
