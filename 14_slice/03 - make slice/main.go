package main

import "fmt"


func main() {

	// If you want to create a slice, you should use the built-in make function:
	// To set array length, we can use make() function.
	//
	mySlice := make([]int, 0, 3)


	fmt.Println("-----------------")
	fmt.Println(mySlice)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))
	fmt.Println("-----------------")

	for i := 0; i < 80; i++ {
		mySlice = append(mySlice, i)
		fmt.Println("Len:", len(mySlice), "Capacity:", cap(mySlice), "Value: ", mySlice[i])
	}
}
