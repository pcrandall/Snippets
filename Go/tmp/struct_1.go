package main

import (
	"fmt"
)

func main() {

	goku := Saiyan{"Goku", 9000}

	Super(goku)

	fmt.Println("we making changes to a copy of goku, need to pass in the address of")

	fmt.Println(goku.Power)

}

type Saiyan struct {
	Name  string
	Power int
}

func Super(s Saiyan) {
	s.Power += 10000
}
