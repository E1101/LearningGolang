package main

import "fmt"

func max(x int) int {
	return 42 + x
}

func main() {
	max := max(7)
	fmt.Println(max) // max is now the result, not the function

	max(4) // cannot call non-function max (type int)
}

// don't do this; bad coding practice to shadow variables
