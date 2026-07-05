package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"course_name"`
	Price    int      `json:"course_price"`
	Platform string   `json:"course_platform"`
	Password string   `json:"-"`
	Tag      []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome JSON video")
	DecodeJson()
}

func EncodeJson() {
	lcoCourses := []course{
		{"React JS bootcamp", 299, "LearnCodeOnline.in", "abc123", []string{"Web Dev", "js"}},
		{"MERN bootcamp", 199, "LearnCodeOnline.in", "bcd123", []string{"full-stack", "js"}},
		{"Angular bootcamp", 299, "LearnCodeOnline.in", "hit123", nil},
	}

	//package this data as json data

	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	//finalJson, err := json.Marshal(lcoCourses)
	fmt.Println(string(finalJson))
	CheckNil(err)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(
		`
	{
		"course_name": "React JS bootcamp",
		"course_price": 299,
		"course_platform": "LearnCodeOnline.in",
		"tags": ["Web Dev","js"]
	}
`)

	var lcoCourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Printf("Json Valid\n")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("Json Not Valid")
	}

	// some cases where you just want to add data to key value
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for key, value := range myOnlineData {
		fmt.Printf("key: %s, value: %v\n", key, value)
	}
}

func CheckNil(err error) {
	if err != nil {
		panic(err)
	}
}
