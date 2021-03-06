package main

import "fmt"


func main() {
	var x [256]byte

	fmt.Println(len(x))  // 256
	fmt.Println(x[42])   // 0

	for i := 0; i < 256; i++ {
		x[i] = byte(i)
	}

	for i, v := range x {
		fmt.Printf("%v - %T - %b\n", v, v, v) // 2 - uint8 - 10
		if i > 50 {
			break
		}
	}
}
