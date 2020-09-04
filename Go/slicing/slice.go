package main

import "fmt"

func main() {
	scores := []int{1, 2, 3, 4, 5}
	scores = removeAtIndex(scores, 2)
	fmt.Println("sortedScores", mergeSort(scores))
}

func removeAtIndex(source []int, index int) []int {
	lastIndex := len(source) - 1
	// swap the last value and the value we want to remove
	source[index], source[lastIndex] = source[lastIndex], source[index]
	fmt.Println(source)
	return source[:lastIndex]
}

func merge(arr1 []int, arr2 []int) []int {
	println("arr1", arr1, len(arr1), "arr2", arr2, len(arr2))
	results := make([]int, 1)
	i, j := 0, 0
	for i < len(arr1) && j < len(arr2) {
		if arr2[j] > arr1[i] {
			results = append(results, arr1[i])
			i++
		} else {
			results = append(results, arr2[j])
			j++
		}
		fmt.Println("i: ", i, "j: ", j)
	}
	return results
}

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	left := []int(arr[0 : len(arr)/2])
	right := []int(arr[len(arr)/2:])
	fmt.Println("left", left, "right", right)
	return merge(left, right)
}
