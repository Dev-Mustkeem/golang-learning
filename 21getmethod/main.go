package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Print("Welcome to web verb video: GET")
	performGetRequest()
}

func performGetRequest() {
	const myUrl = "http://localhost:8000"

	response, err := http.Get(myUrl)
	checkError(err)

	defer response.Body.Close()

	fmt.Println("status code:", response.StatusCode)
	fmt.Println("Content Length:", response.ContentLength)

	content, err := io.ReadAll(response.Body)
	fmt.Println("Content Length:", len(content))
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
