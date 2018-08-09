package main

import (
	"fmt"
	"time"
)


func main() {
	today := time.Now()
	fmt.Println(today)
}

// note:
// The pointer to the Time type should not be used.
// If only the value (not pointer to variable) is used,
// the Time instance is considered to be safe for use across
// multiple goroutines. The only exception is with serialization.
