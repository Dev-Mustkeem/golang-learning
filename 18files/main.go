package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to files in go lang")

	content := "This needs to go in a file - Mustkeem.in"

	file, err := os.Create("./myFile.txt")
	checkNil(err)

	length, error := io.WriteString(file, content)
	checkNil(error)

	fmt.Println("Length is:", length)

	defer file.Close()
	readFile("./myFile.txt")
}

func readFile(fileName string) {
	databyte, err := os.ReadFile(fileName)
	checkNil(err)
	fmt.Println("Data is:", string(databyte))
}

func checkNil(err error) {
	if err != nil {
		panic(err)
	}
}
