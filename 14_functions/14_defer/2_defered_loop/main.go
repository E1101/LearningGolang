package main

import "fmt"

func main() {

	for i := 0; i < 3; i++ {
		defer func() {
			fmt.Println(i)
		}()
	}

	// Output:
	// 3
	// 3
	// 3


	// Solution:

	for i := 0; i < 3; i++ {
		defer func(i int) {
			fmt.Println(i)
		}(i)
	}

	// Or:

	for i := 0; i < 3; i++ {
		defer fmt.Println(i)
	}

}

