package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup


func main() {
	wg.Add(2)

	go foo()
	go bar()

	wg.Wait()
}


// ..

func foo() {
	defer wg.Done()

	for i := 0; i < 45; i++ {
		fmt.Println("Foo:", i)
	}
}

func bar() {
	defer wg.Done()

	for i := 0; i < 45; i++ {
		fmt.Println("Bar:", i)
	}
}
