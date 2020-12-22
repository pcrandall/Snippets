package main

import (
	"fmt"
)

func main() {
	//memoization
	//struct, keys will be arg to fn, value will be the return value
	memo := make(map[int]int)

	myFib := fib(40, memo)

	fmt.Println(myFib)
}

func fib(n int, m map[int]int) int {
	// check if it's already been calculated
	val, ok := m[n]
	if ok == true {
		fmt.Println("fib num:", n, " val", val)
		return val
	}

	if n <= 2 {
		return 1
	}
	m[n] = fib(n-1, m) + fib(n-2, m)
	return fib(n-1, m) + fib(n-2, m)
}
