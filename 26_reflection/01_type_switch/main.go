package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a uint8 = 2
	var b int = 37
	var c string = "3.2"

	// Sums a uint8, an int,
	// and a string
	res := sum(a, b, c)
	fmt.Printf("Result: %f\n", res)
}

func sum(v ...interface{}) float64 {
	var res float64 = 0

	// Loops through all of the values
	// given, and switches over them
	// based on type
	for _, val := range v {
		switch val.(type) {
		// For each type that you
		// support (int, int64, uint8,
		// string), converts to float64 and sum
		case int:
			res += float64(val.(int))
		case int64:
			res += float64(val.(int64))
		case uint8:
			res += float64(val.(uint8))
		case string:
			// For a string, you use the strconv
			// library to convert the string to a float64
			a, err := strconv.ParseFloat(val.(string), 64)
			if err != nil {
				panic(err)
			}
			res += a
		default:
			// If the type isn’t one of the four you
			// support, prints an error and ignores
			fmt.Printf("Unsupported type %T. Ignoring.\n", val)
		}
	}

	return res
}
