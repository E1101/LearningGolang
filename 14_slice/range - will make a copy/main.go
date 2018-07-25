package main

import "fmt"

// !! It’s important to know that range is making a copy
// of the value, not returning a reference.

func main() {
	// Create a slice of integers.
	// Contains a length and capacity of 4 elements.
	slice := []int{10, 20, 30, 40}

	// Iterate over each element and display the value and addresses.
	for index, value := range slice {
		fmt.Printf("Value: %d Value-Addr: %X ElemAddr: %X\n",
			value, &value, &slice[index])

		// Value: 10	Value-Addr: 10500168    ElemAddr: 1052E100
		// ..
	}
}
