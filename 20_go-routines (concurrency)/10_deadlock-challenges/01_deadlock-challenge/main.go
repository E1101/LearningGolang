package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	c <- 1
	fmt.Println(<-c)
}

// This results in a deadlock.
// Can you determine why?
// And what would you do to fix it?


// When the send operation executes, it blocks the
// main goroutine, which means that it blocks the
// execution of the entire program because the send operation
// is waiting for a corresponding receive operation on the
// same channel.
// Because the send operation blocks
// execution, the receive operation cannot be executed,
// which causes a deadlock.
