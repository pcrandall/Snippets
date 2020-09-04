package main

import (
	"fmt"
)

func main() {
	power := getPower()
	fmt.Printf("we deed it %d\n", power)
}

func getPower() int {
	return 9001
}
