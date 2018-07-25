package main

// When your variable is an array, this means the entire array, regardless
// of its size, is copied and passed to the function.

func main() {
	// Declare an array of 8 megabytes.
	var array [1e6]int

	// Pass the array to the function foo.
	fooFix(&array)
}

// Function foo accepts an array of one million integers.
func fooFix(array *[1e6]int) {
	// ...
}
