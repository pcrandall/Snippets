package main

import (
	"fmt"
)

type Saiyan struct {
	Name  string
	Power int
}

/*
 Structures don't have constructors. Instead, you create a function that
 returns an instance of the desired type (like a factory)
*/

func main() {
	goku := newSaiyan("Goku", 9000)
	goku.print()
}

func newSaiyan(name string, power int) Saiyan {
	return Saiyan{
		Name:  name,
		Power: power,
	}
}

func (s *Saiyan) print() {
	fmt.Printf("%s's power: %d\n", s.Name, s.Power)
}

/*
Despite the lack of constructors, Go does have a built-in new function which is
used to allocate the memory required by a type. The result of new(X) is the same as &X{}:

goku := new(Saiyan)
// same as
goku := &Saiyan{}

Which you use is up to you, but you'll find that most people prefer the latter
whenever they have fields to initialize, since it tends to be easier to read:

goku := new(Saiyan)
goku.name = "goku"
goku.power = 9001

//vs

goku := &Saiyan {
  name: "goku",
  power: 9000,
}

Whichever approach you choose, if you follow the factory pattern above, you can
shield the rest of your code from knowing and worrying about any of the allocation details.
*/
