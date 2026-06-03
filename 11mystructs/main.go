package main

import "fmt"

func main() {
	fmt.Println("Structs in Go Lang")

	// there is no inheritance in Go Lang also there is no super or parent keyword in Go Lang

	user1 := User{"Mustkeem Baraskar", "abc@gmail.com", true, 25}
	fmt.Printf("User1: %+v\n", user1)

	fmt.Printf("Name: %v\nEmail: %v\nStatus: %v\nAge: %v\n", user1.Name, user1.Email, user1.Status, user1.Age)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
