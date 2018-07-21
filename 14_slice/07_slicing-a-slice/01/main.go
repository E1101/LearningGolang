package main

import "fmt"

func main() {

	var results []int

	fmt.Println(results) // []



	mySlice := []string {"a", "b", "c", "g", "m", "z"}

	fmt.Println(mySlice)       // [a b c g m z]

	// slicing a slice
	fmt.Println(mySlice[2:4])  // [c g]

	// index access; accessing by index
	fmt.Println(mySlice[2])    // [c]


	// index access; accessing by index
	fmt.Println("myString"[2]) // 83
}
