package main

import "fmt"

func main() {
	/*
	 !! An array is a numbered sequence of elements of a single type with a fixed length.
	 */
	var x [58]int

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[42])

	x[42] = 777
	fmt.Println(x[42])
}
