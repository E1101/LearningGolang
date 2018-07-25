package main

import "fmt"

func main() {
	// Declare a two dimensional integer array of four elements
	// by two elements.
	var array [4][2]int

	// Use an array literal to declare and initialize a two
	// dimensional integer array.
	array = [4][2]int{{10, 11}, {20, 21}, {30, 31}, {40, 41}}
	// Declare and initialize index 1 and 3 of the outer array.
	array = [4][2]int{1: {20, 21}, 3: {40, 41}}
	// Declare and initialize individual elements of the outer
	// and inner array.
	array = [4][2]int{1: {0: 20}, 3: {1: 41}}

	fmt.Printf("%v", array)



	// To access an individual element, use the [ ] operator again and a bit of composition.
	//

	// Declare a two dimensional integer array of two elements.
	var array2 [2][2]int

	// Set integer values to each individual element.
	array2[0][0] = 10
	array2[0][1] = 20
	array2[1][0] = 30
	array2[1][1] = 40


	// Assigning multidimensional arrays of the same type
	//

	// Declare two different two dimensional integer arrays.
	var array1 [2][2]int
	var array3 [2][2]int
	// Add integer
	array3[0][0] = 10
	array3[0][1] = 20
	array3[1][0] = 30
	array3[1][1] = 40

	// Copy the values from array2 into array1.
	array1 = array3

	fmt.Printf("%v", array1)


	// Assigning multidimensional arrays by index
	//

	// Copy index 1 of array1 into a new array of the same type.
	var array4 [2]int = array1[1]

	// Copy the integer found in index 1 of the outer array
	// and index 0 of the interior array into a new variable of
	// type integer.
	var value int = array1[1][0]

	fmt.Printf("%v", array4)
	fmt.Printf("%T", value)
}
