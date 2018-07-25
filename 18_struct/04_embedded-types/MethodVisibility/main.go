package MethodVisibility

import (
	"fmt"
)

func main()  {

	// Create a variable of the unexported type using the exported
	// New function from the package counters.
	counter := New(10)

	fmt.Printf("Counter: %d\n", counter)
}
