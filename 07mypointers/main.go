package main

import "fmt"

func main() {
	fmt.Println("Welcom to class on pointers")

	// var ptr *int
	// fmt.Printf("The value of ptr is %v\n", ptr)

	myNumber := 23

	var ptr = &myNumber

	fmt.Println("Value of actual pointer is", ptr)
	fmt.Println("Value of the number is", *ptr)

}
