package main

import "fmt"

type foo int

func main() {
	var myAge foo

	myAge = 44
	fmt.Printf("%T %v \n", myAge, myAge)
}

/*
user defined types - we declare a new type, foo
the underlying type of foo: int

conversion:int(myAge)
converting type foo to type int

THIS CODE IS ONLY FOR EXAMPLE
IT IS A BAD PRACTICE TO ALIAS TYPES
one exception: if you need to attach methods to a type
	see the time package for an example of this
godoc.org/time
type Duration int64

Duration has methods attached to it
*/