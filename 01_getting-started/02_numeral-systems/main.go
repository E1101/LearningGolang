package main

import "fmt"

func main()  {

	// Decimal
	fmt.Println(42)

	// Binary
	fmt.Printf("%d - %b \n", 42, 42)

	// Hexadecimal
	fmt.Printf("%d - %b - %x \n", 42, 42, 42)
	fmt.Printf("%d - %b - %#x \n", 42, 42, 42)
	fmt.Printf("%d - %b - %#X \n", 42, 42, 42)
	fmt.Printf("%d \t %b \t %#X \n", 42, 42, 42)

	// Loop
	for i := 1000000; i < 1000100; i++ {
		fmt.Printf("%d \t %b \t %x \n", i, i, i)
	}
}
