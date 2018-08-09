package main

import (
	"fmt"
	"time"
)

func main() {

	// Set the epoch from int64
	t := time.Unix(0, 0)
	fmt.Println(t) // 1970-01-01 03:30:00 +0330 +0330

	// Get the epoch
	// from Time instance
	epoch := t.Unix()
	fmt.Println(epoch) // 0


	// Current epoch time
	apochNow := time.Now().Unix()
	fmt.Printf("Epoch time in seconds: %d\n", apochNow)
	// > Epoch time in seconds: 1533795853

	apochNano := time.Now().UnixNano()
	fmt.Printf("Epoch time in nano-seconds: %d\n", apochNano)
	// > Epoch time in nano-seconds: 1533795853704560226
}
