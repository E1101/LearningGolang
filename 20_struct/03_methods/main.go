package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

// value receiver, the method will always be operating against a copy of the value
// used to make the method call
func (p person) fullName() string {
	return p.first + p.last
}

// This time, the receiver is not a value of type user but a pointer of
// type user . When you call a method declared with a pointer receiver, the value used to
// make the call is shared with the method.
// changeEmail implements a method with a pointer receiver.
func (p *person) changeName(first string) {
	p.first = first
}

func main() {
	p1 := person{"James", "Bond", 20}
	p2 := person{"Miss", "Moneypenny", 18}
	fmt.Println(p1.fullName())
	fmt.Println(p2.fullName())

	lisa := &person{"Lisa", "lisaeian", 34}
	lisa.fullName()
	// You can imagine that Go is performing the following operation.
	(*lisa).fullName()


	(&lisa).changeName("Lis")
}
