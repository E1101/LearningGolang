package main

import "fmt"

func main() {

	slice := []int {
		0,
		1,
		2,
	}

	fmt.Printf("%v", slice)
	// [0 1 2]


	// The values changes through func
	foo(slice)


	fmt.Printf("%v\r\n", slice)
	// [10 11 12]
}


// ..

func foo(slice []int)  {
	for index := range slice {
		slice[index] += 10
	}
}
