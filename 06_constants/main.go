package main

import "fmt"

const p = "death & taxes"

// Multiple Initialization
const (
	pi       = 3.14
	language = "Go"
)

func main() {

	const q = 42

	fmt.Println("p - ", p)
	fmt.Println("q - ", q)

	// ..

	fmt.Println(pi)
	fmt.Println(language)
}

// a CONSTANT is a simple unchanging value
