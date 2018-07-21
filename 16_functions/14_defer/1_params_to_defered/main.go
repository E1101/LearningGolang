package main

import "fmt"

// Params passed to a deferred func are evaluated
// when the defer is registered not when it runs.
func main() {
	func() {
		m := &message{content: "Hello"}
		defer fmt.Print(m.print())
		m.set("World")
		// deferred func runs
	}()

	// Output:
	// "Hello"
}


// ..

type message struct {
	content string
}

func (p *message) set(c string) {
	p.content = c
}

func (p *message) print() string {
	return p.content
}
