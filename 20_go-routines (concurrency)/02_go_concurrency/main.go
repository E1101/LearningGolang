package main

import (
	"fmt"
)

func main() {
	go foo()
	go bar()

	var input string
	fmt.Scanln(&input)
}


// ..

func foo() {
	for i := 0; i < 45; i++ {
		fmt.Println("Foo:", i)
	}
}

func bar() {
	for i := 0; i < 45; i++ {
		fmt.Println("Bar:", i)
	}
}
