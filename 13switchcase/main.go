package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to game of ludo")

	rand.Seed(time.Now().UnixNano())

	diceNumber := rand.Intn(6) + 1

	fmt.Printf("You got %d on the dice\n", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("You can move 1 step")
	case 2:
		fmt.Println("You can move 2 steps")
	case 3:
		fmt.Println("You can move 3 steps")
	case 4:
		fmt.Println("You can move 4 steps")
	case 5:
		fmt.Println("You can move 5 steps")
	case 6:
		fmt.Println("You can move 6 steps and roll the dice again")
	default:
		fmt.Println("Invalid dice number")
	}
}
