package main

import "fmt"


func main() {
	// It’s similar to creating an array, except you don’t
	// specify a value inside of the [ ] operator. The initial length
	// and capacity will be based on the number of elements you initialize
	mySlice := []int {1, 3, 5, 7, 9, 11}


	fmt.Printf("%T\n", mySlice)
	fmt.Println(mySlice)
}
