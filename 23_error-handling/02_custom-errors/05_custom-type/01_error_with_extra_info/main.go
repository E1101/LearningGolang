// you may want your errors to contain more information than a simple string.
// In such cases, you may choose to create a custom error type.
package main

import "fmt"

type ParseError struct {
	Message    string
	Line, Char int
}

// Implements the Error interface
func (p *ParseError) Error() string {
	format := "%s o1n Line %d, Char %d"
	return fmt.Sprintf(format, p.Message, p.Line, p.Char)
}

