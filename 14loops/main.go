package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in golang")

	// days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for index, day := range days {
	// 	fmt.Printf("The index is %v and the value is %v\n", index, day)
	// }

	rougeValue := 1
	for rougeValue < 10 {

		if rougeValue == 2 {
			goto lco
		}

		fmt.Printf("The value of rougeValue is %v\n", rougeValue)
		rougeValue++
	}

lco:
	fmt.Println("This is lco")
}
