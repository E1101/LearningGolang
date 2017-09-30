package main

import "fmt"

func main()  {

	var message = "Hello World!"
	var a, b, c = 1, 2, 3

	fmt.Println(message, a, b, c)


	// Mixed Type
	var message1 = "Hello World!"
	var a1, b1, c1 = 1, false, 3

	fmt.Println(message1, a1, b1, c1)
}
