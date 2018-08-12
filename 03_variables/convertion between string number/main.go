package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 123
	y := fmt.Sprintf("%d", x)
	fmt.Println(y, strconv.Itoa(x)) // "123 123"


	// FormatInt and FormatUint can be used to format numbers in a different base:
	fmt.Println( strconv.FormatInt(int64(x), 2) ) // "1111011"


	x1, _ := strconv.Atoi("123")
	// x is an int
	y1, _ := strconv.ParseInt("123", 10, 64) // base 10, up to 64 bits

	fmt.Println( x1, y1 )
}
