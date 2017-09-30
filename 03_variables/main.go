package main

import "fmt"

var a = "this is stored in the variable a"     // package scope
var b, c string = "stored in b", "stored in c" // package scope
var d string                                   // package scope

func main()  {

	d = "stored in d" // declaration above; assignment here; package scope

	/*
	 function scope - subsequent variables have func scope:
	 */

	// Declare Variable
	var message string
	message = "Hello World."
	fmt.Println(message)

	// variable creation with assignment
	var x string = "Hello, World"
	fmt.Println(x)

	// variables can change their value throughout the lifetime of a program
	x = "second"
	fmt.Println(x)

	x = x + " later"
	x += " concat"
	fmt.Println(x)


	// Declare Many At Onnce
	var message1 string
	var a1, b1, c1 int

	a1 = 1
	message1 = "Hello World!"

	fmt.Println(message1, a1, b1, c1)


	// shorter statement
	x2 := "Hello, World"
	fmt.Println(x2)


	// When variables created without assignment it has zero state
	var a int
	var b string
	var c float64
	var d bool

	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)

	fmt.Println()
}
