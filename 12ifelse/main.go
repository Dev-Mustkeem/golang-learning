package main

import "fmt"

func main() {

	fmt.Println("If else in Go Lang")

	loginCount := 23

	var result string

	if loginCount < 10 {
		result = "Regular User"
	} else if loginCount > 20 {
		result = "Active User"
	} else {
		result = "Not a Regular User"
	}

	fmt.Printf("Result: %v\n", result)

	if 9%2 == 0 {
		fmt.Println("Even Number")
	} else {
		fmt.Println("Odd Number")
	}

	if num := 9; num%2 == 0 {
		fmt.Println("Even Number")
	} else {
		fmt.Println("Odd Number")
	}
}
