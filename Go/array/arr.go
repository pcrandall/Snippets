package main

import "fmt"

func main() {
	/* Arrays are static in go*/
	scores := [4]int{10, 593, 9001, 43}
	/* Using for loop with range to iterate through an array */

	for index, value := range scores {
		fmt.Println("Index:", index, "Value:", value)
	}
}
