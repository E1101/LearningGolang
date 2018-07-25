package main

import "fmt"

func main() {
	// Create two slices each initialized with two integers.
	s1 := []int{1, 2}
	s2 := []int{3, 4}

	// Append the two slices together and display the results.
	fmt.Printf("%v\n", append(s1, s2...)) // [1, 2, 3, 4]
}
