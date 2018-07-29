package main

import "fmt"


func main() {
	// Call variadic function
	//
	n := average(43, 56, 87, 12, 45, 57)
	fmt.Println(n)


	// Providing a slice as an argument
	//
	nums := []float64{1, 2, 3, 4, 5}
	total := average(nums...)
	fmt.Println("The Average is:", total)
}

func average(sf ...float64) float64 {
	fmt.Println(sf)
	fmt.Printf("%T \n", sf) // []float64

	var total float64

	for _, v := range sf {
		total += v
	}

	return total / float64( len(sf) )
}
