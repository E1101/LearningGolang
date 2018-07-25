package main

import "fmt"

// A nil slice is the most common way you create slices in Go.
// They can be used with many of the standard library and
// built-in functions that work with slices.


func main() {

	// nil Slice
	//
	// Use a slice literal to create an empty slice of integers.
	// Create a nil slice of integers.
	var slice []int


	// Empty Slice
	//
	// empty slices are useful when you want to represent
	// an empty collection, such as when a database query
	// returns zero results

	// This is same as variable slice defined below
	sliceA := []int{}

	// Use make to create an empty slice of integers.
	slice1 := make([]int, 0)



	// !! Regardless of whether you’re using a nil slice or an empty slice,
	// the built-in functions append , len , and cap work the same.
	//

	fmt.Printf("%T", slice)
	fmt.Printf("%T", sliceA)
	fmt.Printf("%T", slice1)
}
