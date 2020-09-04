package main

import (
	"fmt"
)

func main() {
	/* To initiate a slice leave the value in square brackets empty */
	scores := []int{10, 593, 9001, 43}
	/* Using for loop with range to iterate through an array */
	for index, value := range scores {
		fmt.Println("Index:", index, "Value:", value)
	}

	/* use make for new slice with len of 10 and cap of 10 */

	slScores := make([]int, 10)
	for index, value := range slScores {
		fmt.Println("Index:", index, "Value:", value)
	}

	/* use make for new slice with len of 0 and cap of 10 */

	/* empty nothing gets printed out */
	lenScors := make([]int, 0, 10)
	for index, value := range lenScors {
		fmt.Println("Index:", index, "Value:", value)
	}

	/* 15 get insterted at first index */
	lenScors = append(lenScors, 15)
	fmt.Println(lenScors)

	// without specifying the len you create a slice with len and cap of 5
	// when you append to the slice, the value gets added on the end
	nSlice := make([]int, 5)
	nSlice = append(nSlice, 9332)
	fmt.Println(nSlice)

	nScores := []int{1, 2, 3, 4, 5}
	// only prints out index 2 & 3
	slice := nScores[2:4]
	fmt.Println(slice)

	// changes the nScores slice
	slice[0] = 999
	fmt.Println(nScores)
}
