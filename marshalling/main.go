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

	user3ConvertedTOByteString, error3 := json.Marshal(user3)

	if error3 != nil {
		fmt.Println("Error converting user3 to JSON:", error3)
		return
	}

	user3ConvertedTOJson := string(user3ConvertedTOByteString)

	fmt.Println("User3 in JSON format:", user3ConvertedTOJson)
	fmt.Println("User3 :", user3)

	// Removing a field from the JSON output

	user4 := UserWhenPasswordIsNotSharedViaJson{
		Name:     "Januda",
		Age:      24,
		Password: "mysecurepassword",
	}

	user4ConvertedTOByteString, error4 := json.Marshal(user4)

	if error4 != nil {
		fmt.Println("Error converting user4 to JSON:", error4)
		return
	}

	user4ConvertedTOJson := string(user4ConvertedTOByteString)

	fmt.Println("User4 in JSON format:", user4ConvertedTOJson)
	fmt.Println("User4 :", user4)

}

type User struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Password string `json:"password"`
}

type UserWithOmitEmpty struct {
	// omitempty omits the field from the JSON output if it is empty as well as the field from the struct
	Name     string `json:"name,omitempty"`
	Age      int    `json:"age"`
	Password string `json:"password"`
}

type UserWhenPasswordIsNotSharedViaJson struct {
	// "-" omits the field from the JSON output no matter what but keeps it in the struct
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Password string `json:"-"`
}
