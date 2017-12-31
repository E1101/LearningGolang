package main

import (
	"fmt"
	"sync"
)

func main() {

	c := make(chan int)

	var wg sync.WaitGroup

	// Fix: is remove wg.Add from within closures
	//      add here.

	// wg.Add(2)

	go func() {
		wg.Add(1)             // waitGroup want add in same time with Lines Marked X
		for i := 0; i < 10; i++ {
			c <- i
		}
		wg.Done()
	}()

	go func() {
		wg.Add(1)            // X
		for i := 0; i < 10; i++ {
			c <- i
		}
		wg.Done()
	}()

	go func() {
		wg.Wait()                 // X
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}
