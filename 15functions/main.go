package main

import "fmt"

func main() {
	greeter()
	result, message := multiReturn(1, 2, 3, 4, 5)
	fmt.Printf("%s%v\n", message, result)
}

func add(a int, b int) int {
	return a + b
}

func proAdd(values ...int) int {
	sum := 0
	for _, v := range values {
		sum += v
	}
	return sum
}

func multiReturn(values ...int) (int, string) {
	sum := 0
	for _, v := range values {
		sum += v
	}
	return sum, "Kya re haule: "
}

func greeter() {
	fmt.Print("Namasty from go lang")
}
