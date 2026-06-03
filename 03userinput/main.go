package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcom := "Welcome to Go Programming"
	fmt.Println(welcom)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our pizza: ")

	//comma ok syntax or comma error syntax

	input, _ := reader.ReadString('\n')

	fmt.Printf("Type of this rating is %T:", input)
}
