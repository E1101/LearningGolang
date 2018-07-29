package main

import (
	"fmt"
	"regexp/syntax"
)


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
	// in a multi-line statement
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


	// Declare an integer array of five elements.
	// Initialize index 1 and 2 with specific values.
	// The rest of the elements contain their zero value.
	arrayIndex := [5]int{1: 10, 2: 20}
	fmt.Printf("%T ", arrayIndex)


	// Declare an integer pointer array of five elements.
	// Initialize index 0 and 1 of the array with integer pointers.
	array := [5]*int{0: new(int), 1: new(int)}
	// Assign values to index 0 and 1.
	*array[0] = 10
	*array[1] = 20
	// Copy the values from array2 into array1.
	// After the copy, you have two arrays pointing to the same values
	array2 := array

	fmt.Printf("%v ", array)
	fmt.Printf("%v ", array2)
}
