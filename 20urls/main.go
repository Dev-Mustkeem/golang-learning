package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://jsonplaceholder.typicode.com/comments?postId=1"

func main() {
	fmt.Println("Welcome to handling URLs in GO")
	// fmt.Println(myUrl)

	//Parsing
	result, _ := url.Parse(myUrl)

	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()

	fmt.Printf("Type of Q Params is %T\n", qparams)
	fmt.Println(qparams["postId"])

	for _, value := range qparams {
		fmt.Println("Params is: ", value)
	}

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "jsonplaceholder.typicode.com",
		Path:    "/comments",
		RawPath: "postId=1",
	}

	anotherUrl :=
		partsOfUrl.String()
	fmt.Println(anotherUrl)

}

func checkNil(err error) {
	if err != nil {
		panic(err)
	}
}
