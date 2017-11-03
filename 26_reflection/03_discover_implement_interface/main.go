package main

import (
	"bytes"
	"fmt"
)

func main() {
	// Tests whether a
	// *bytes.Buffer is a
	// fmt.Stringer. It is.
	b := bytes.NewBuffer([]byte("Hello"))
	if isStringer(b) {
		fmt.Printf("%T is a stringer\n", b)
	}

	// Tests whether an
	// integer is a
	// fmt.Stringer. It’s not.
	i := 123
	if isStringer(i) {
		fmt.Printf("%T is a stringer\n", i)
	}
}

// Takes an interface{} value and runs a
// type assertion to the desired interface
func isStringer(v interface{}) bool {
	_, ok := v.(fmt.Stringer)
	return ok
}
