package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our pizza shop!")
	fmt.Println("Please rate our pizza from 1 to 5")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	fmt.Println("You rated our pizza:", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println("Error converting input to number:", err)
	} else {
		fmt.Println("Added 1 to your rating", numRating+1)
	}

}
