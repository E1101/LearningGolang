package main

//  switch on types
//  -- normally we switch on value of variable
//  -- go allows you to switch on type of variable

import "fmt"

type contact struct {
	greeting string
	name     string
}

// SwitchOnType works with interfaces
func SwitchOnType(x interface{}) {
	// this is an assert; asserting, "x is of this type"
	switch x.(type) {
	case int:
		fmt.Println("int")

	case string:
		fmt.Println("string")

	case contact:
		fmt.Println("contact")

	default:
		fmt.Println("unknown")
	}
}

func main() {
	SwitchOnType(7)

	SwitchOnType("McLeod")

	var t = contact{"Good to see you,", "Tim"}
	SwitchOnType(t)

	SwitchOnType(t.greeting)

	SwitchOnType(t.name)
}
