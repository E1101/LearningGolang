package main

import "fmt"

func main() {

	var myGreeting = make(map[string]string)

	// the length of a map (found by doing len(x) )
	// can change as we add new items to it.
	myGreeting["Tim"] = "Good morning."
	myGreeting["Jenny"] = "Bonjour."

	fmt.Println(myGreeting)
}
