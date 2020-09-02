package main

import (
	"fmt"
)

type Saiyan struct {
	Name  string
	Power int
}

func main() {

	goku := Saiyan{
		Name:  "Goku",
		Power: 9000,
	}

	power := getPower()

	fmt.Printf("we deed it %d\n", power)

	fmt.Printf("%s we be saiyan it %d\n", goku.Name, goku.Power)
}

func getPower() int {
	return 9001
}
