package main


// Since the data associated with a slice is contained in
// the underlying array, there are no problems passing a
// copy of a slice to any function. Only the slice is being
// copied, not the underlying array

func main() {
	// Allocate a slice of 1 million integers.
	slice := make([]int, 1e6)

	// Pass the slice to the function foo.
	slice = foo(slice)
}

// Function foo accepts a slice of integers and returns the slice back.
func foo(slice []int) []int {
	// ..
	return slice
}
