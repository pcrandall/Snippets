package main

import (
	"fmt"
)

type Saiyan struct {
	Name  string
	Power int
}

func main() {

	goku := &Saiyan{"Goku", 9000}
	gohan := &Saiyan{"Gohan", 1000}

	fmt.Println("we making changes to the originals")

	goku.Print()
	gohan.Print()

	goku.Super()
	gohan.Super()

	fmt.Println("after calling the super method")

	goku.Print()
	gohan.Print()
}

func (s *Saiyan) Super() {
	s.Power += 10000
}

func (s *Saiyan) Print() {
	fmt.Printf("%ss power: %d\n", s.Name, s.Power)
}
