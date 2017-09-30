package main

import "fmt"

func main() {
	slice1 := []int{1,2,3}
	slice2 := make([]int, 2)
	copy(slice2, slice1)

	// slice1 has [1,2,3] and slice2 has [1,2]
	// because slice2 has room for only two elements,
	// only the first two elements of slice1 are copied.
	fmt.Println(slice1, slice2)
}
