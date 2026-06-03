package main

import "fmt"

func main() {
	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["PY"] = "Python"
	languages["RB"] = "Ruby"

	fmt.Println("langugages:", languages)

	fmt.Println("JS shorts for:", languages["JS"])

	delete(languages, "RB")
	fmt.Println("langugages:", languages)

	//loops in maps
	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
}
