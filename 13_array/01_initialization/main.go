package main

import "fmt"


func main() {
	//
	// !! An array is a numbered sequence of elements of a single type with a fixed length.
	//
	var x [58]int

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[42])

	x[42] = 777
	fmt.Println(x[42])


	// use an array literal to declare and initialize arrays
	y := [5]int {5,10,15,20,25}
	fmt.Printf("%T ", y)


	// When you declare and initialize arrays with the array literal you can provide the initialization expression
	// in a multiline statement
	z := [5]int {
		5,
		10,
		15,
		20,
		25,
	}
	fmt.Printf("%T ", z)


	// When arrays are initialized using an array literal, you can provide values for specific elements
	langs := [4]string {0: "Go", 3: "Julia"}
	fmt.Printf("%T ", langs)


	// Here the length of the array is determined by the number of elements provided in the initialization
	// expression.
	t := [...]int {5,10,15,20,25}
	fmt.Printf("%T ", t)
}
