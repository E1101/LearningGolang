package main

import "fmt"


func main() {
	// Create a slice of integers.
	// Contains a length and capacity of 4 elements.
	slice := []int{10, 20, 30, 40}

	// Iterate over each element and display each value.
	for index, value := range slice {
		fmt.Printf("Index: %d Value: %d\n", index, value)
		// Index: 0	Value: 10
		// ..
	}
}
