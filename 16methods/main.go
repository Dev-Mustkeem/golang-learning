package main

import "fmt"

func main() {
	user := User{Name: "John Doe", Age: 30, Email: "john.doe@example.com"}
	user.GetEmail()
}

type User struct {
	Name  string
	Age   int
	Email string
}

func (user User) GetEmail() {
	fmt.Println("Hello User: ", user.Email)
}
