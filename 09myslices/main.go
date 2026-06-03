package main

import (
	"fmt"
	"sort"
)

func main() {

	var fruitList = []string{"Apple", "Tomato", "Peach"}

	fmt.Printf("Type of fruitList: %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")

	fmt.Println("Fruit List:", fruitList)

	fruitList = append(fruitList[1:])
	// fmt.Println("Fruit List:", fruitList)

	fruitList = append(fruitList[1:3])
	// fmt.Println("Fruit List:", fruitList)

	fruitList = append(fruitList[:3])
	// fmt.Println("Fruit List:", fruitList)

	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 986
	highScores[2] = 654
	highScores[3] = 765

	fmt.Println("High Scores:", highScores)

	highScores = append(highScores, 555, 666, 321)
	fmt.Println("High Scores:", highScores)

	sort.Ints(highScores)
	fmt.Println("Sorted High Scores:", highScores)
	fmt.Println("Is sorted: ", sort.IntsAreSorted(highScores))

	//how to remove a value from slice based on index
	var courses = []string{"ReactJS", "NodeJS", "Angular", "VueJS", "Django"}

	fmt.Println("Courses:", courses)
	var index int = 2

	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println("Courses:", courses)
}
