package main

import "fmt"

type Person struct {
	Name string
}

type Saiyan struct {
	*Person
	Power int
}

func (p *Person) Introduce() {
	fmt.Printf("Hi, I'm %s\n", p.Name)
}

/*
Uncomment the following for an example of overloading,
you can overwrite the inherited Person introduce method with your own
func (p *Saiyan) Introduce() {
	fmt.Printf("Hi, I'm %s, I just wanna let you know my power is %d\n", p.Name, p.Power)
}
*/

func main() {
	goku := &Saiyan{
		Person: &Person{"Goku"},
		Power:  9001,
	}

	goku.Introduce()
}
