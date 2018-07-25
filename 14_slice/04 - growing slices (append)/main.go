package main

import "fmt"


// One of the advantages of using a slice over using
// an array is that you can grow the capacity of your
// slice as needed.

func main() {
	// 3 is length - number of elements referred to by the slice
	// 5 is capacity - number of elements in the underlying array
	greeting := make([]string, 3, 5)


	greeting[0] = "Good morning!"
	greeting[1] = "Bonjour!"
	greeting[2] = "buenos dias!"
	greeting = append(greeting, "Suprabadham")

	fmt.Println(greeting[3])
}
