package main

import "fmt"

func main() {
	fmt.Println("Welcome to class on arrays")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Orange"
	fruitList[3] = "Peach"

	fmt.Println("Fruit List is: ", fruitList)                //Output: Array list with extra space for uninitialized element :- Fruit List is:  [Apple Orange  Peach]
	fmt.Println("Length of Fruit List is: ", len(fruitList)) //Output: Length of Fruit List is:  4 cause we have declared the array with size 4

	var vegList = [3]string{"Potato", "Tomato", "Cabbage"}
	fmt.Println("Veg List is: ", vegList) //Output: Veg List is:  [Potato Tomato Cabbage]

	// We can also declare an array without size and let the compiler decide the size based on the number of elements we are initializing
	var anotherFruitList = [...]string{"Mango", "Banana", "Grapes"}
	fmt.Println("Another Fruit List is: ", anotherFruitList) //Output: Another Fruit List is:  [Mango Banana Grapes]
}
