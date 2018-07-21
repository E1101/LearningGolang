package main

import "fmt"

// global package private variable
var x = 42

func main() {
	fmt.Println(x) // 42
	foo()
}

func foo() {
	fmt.Println(x) // 42
}
