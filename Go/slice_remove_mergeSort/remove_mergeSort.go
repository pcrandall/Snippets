package main

import (
	"fmt"
)

func main() {
	scores := []int{1, 2, 88, 46, 52, 98, 1, 35, 1356}
	fmt.Println("Original scores", scores)
	scores = removeAtIndex(scores, 2)
	fmt.Println("scores with removeAtIndex", scores)
	fmt.Println("sortedScores", mergeSort(scores))
}

func removeAtIndex(source []int, index int) []int {
	lastIndex := len(source) - 1
	// swap the last value and the value we want to remove
	source[index], source[lastIndex] = source[lastIndex], source[index]
	return source[:lastIndex]
}

func mergeSort(arr []int) []int {
	fmt.Println("heres my array", arr)
	if len(arr) <= 1 {
		return arr
	}

	left := []int(arr[0 : len(arr)/2])
	right := []int(arr[len(arr)/2:])

	left = mergeSort(left)
	right = mergeSort(right)

	results := make([]int, len(left)+len(right))

	j, k := 0, 0
	for i := 0; i < len(results); i++ {
		if j >= len(left) {
			results[i] = right[k]
			k++
		} else if k >= len(right) {
			results[i] = left[j]
			j++
		} else if left[j] > right[k] {
			results[i] = right[k]
			k++
		} else {
			results[i] = left[j]
			j++
		}
		fmt.Println(results)
	}
	return results
}
