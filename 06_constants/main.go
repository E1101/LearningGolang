package main

import "fmt"

const p = "death & taxes"

// Multiple Initialization
const (
	pi       = 3.14
	language = "Go"
)

const (
	numberGoroutines = 4 // Number of goroutines to use.
	taskLoad             // Amount of work to process. 4
)

func main() {

	const q = 42

	fmt.Println("p - ", p)
	fmt.Println("q - ", q)

	// ..

	fmt.Println(pi)
	fmt.Println(language)

	// ..

	fmt.Println(numberGoroutines)
	fmt.Println(taskLoad)

}

// a CONSTANT is a simple unchanging value
