package main

import (
	"flag"
	"fmt"
)

// Creates a new variable from a flag
var name = flag.String("name", "World", "A name to say hello to.")
// New variable to store flag value
var spanish bool


// Sets variable to the flag value
func init() {
	flag.BoolVar(&spanish, "spanish", false, "Use Spanish language.")
	flag.BoolVar(&spanish, "s", false, "Use Spanish language.")
}


func main() {
	// Parses the flags, placing values in variables
	flag.Parse()

	if spanish == true {
		fmt.Printf("Hola %s!\n", *name)
	} else {
		fmt.Printf("Hello %s!\n", *name)
	}
}


// $ flag -name=payam -s
