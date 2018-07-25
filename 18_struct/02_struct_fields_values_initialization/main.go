package main

import "fmt"

type Circle struct {
	x, y, r float64
}

func main() {

	var circ Circle
	// We can also use the new function
	// This allocates memory for all the fields, sets each of them to their zero value, and
	// returns a pointer to the struct ( *Circle ).
	circl := new(Circle)


	// Initialization with values:

	c := Circle{x: 0, y: 0, r: 5}
	// to leave off the field names if we know the order they were
	c1 := Circle{0, 0, 5}
	// If you want a pointer to the struct
	c2 := &Circle{0, 0, 5}

	fmt.Println(c.x, c.y, c.r)
	fmt.Println(c1.x, c1.y, c1.r)
	fmt.Println(c2.x, c2.y, c2.r)

	fmt.Printf("%T", circ)
	fmt.Printf("%T", circl)
}
