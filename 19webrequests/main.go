package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://instagrp.com"

func main() {
	fmt.Println("Welcome to web requests in Go!")

	response, err := http.Get(url)
	checkNil(err)
	fmt.Printf("Response is of type %T\n", response)
	fmt.Println("Status code:", response.StatusCode)

	defer response.Body.Close() // caller's responsibility to close the response body when done with it

	databytes, err := io.ReadAll(response.Body)
	checkNil(err)
	content := string(databytes)
	fmt.Println(content)
}

func checkNil(err error) {
	if err != nil {
		panic(err)
	}
}
