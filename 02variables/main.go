package main

import "fmt"

func main() {
	var username string = "Mustkeem"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallValue uint8 = 255
	fmt.Println(smallValue)

	var smallFloat float32 = 255.455234234234234
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	var bigFloat float64 = 255.455234234234234
	fmt.Println(bigFloat)
	fmt.Printf("Variable is of type: %T \n", bigFloat)
}
