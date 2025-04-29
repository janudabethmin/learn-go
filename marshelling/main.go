package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")

	user := User{
		Name:     "John Doe",
		Age:      30,
		Password: "securepassword",
	}
	fmt.Println("User:", user)
}

type User struct {
	Name     string
	Age      int
	Password string
}
