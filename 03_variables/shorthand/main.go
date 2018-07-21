package main

import "fmt"

func main() {

	// you can only do this inside a func
	message := "Hello World!"
	a, b, c := 1, false, 3
	d := 4
	e := true

	fmt.Printf("%v %T \n", a, a)

	fmt.Println(message, a, b, c, d, e)
}
