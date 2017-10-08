package main

import (
	"errors"
	"strings"
	"os"
	"fmt"
)

func Concat(parts ...string) (string, error) {
	if len(parts) == 0 {
		// Returns an error if nothing was passed in
		return "", errors.New("No strings supplied")
	}

	// Returns the new string and nil
	return strings.Join(parts, " "), nil
}

func main() {
	// Uses just the args after Args[0].
	// You don’t want the program name.
	args := os.Args[1:]

	if result, err := Concat(args...); err != nil { // Handles the error
		fmt.Printf("Error: %s\n", err)
	} else {
		// non-error case
		fmt.Printf("Concatenated string: '%s'\n", result)
	}
}
