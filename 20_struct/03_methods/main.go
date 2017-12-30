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

/*
(p person) is the "receiver"
it is another parameter
not idiomatic to use "this" or "self"


"Not many people know this, but method notation, i.e. v.Method() is actually syntactic sugar and Go also understands the de-sugared version of it: (T).Method(v). You can see an example here. Naming the receiver like any other parameter reflects that it is, in fact, just another parameter quite well.
This also implies that the receiver-argument inside a method may be nil. This is not the case with this in e.g. Java."
SOURCE:
https://www.reddit.com/r/golang/comments/3qoo36/question_why_is_self_or_this_not_considered_a/?utm_source=golangweekly&utm_medium=email
 */