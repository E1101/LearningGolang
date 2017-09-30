package main

import (
	"fmt"
	"time"
)

func main() {
	var c chan string = make(chan string)

	go pinger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)
}

// Functions:

// pinger is only allowed to send to c.
// Attempting to receive from c will result in a compile-time error.
func pinger(c chan<- string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

// Allowed only to receive from c.
func printer(c <-chan string) {
	for {
		msg := <- c

		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}
