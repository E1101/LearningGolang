package main

import "fmt"


func main() {

	myGreeting := map[string]string{}

	myGreeting["Tim"] = "Good morning."
	myGreeting["Jenny"] = "Bonjour."


	fmt.Println(myGreeting)
	// map[Tim:Good morning. Jenny:Bonjour.]


	// ..

	// Create a map with a key and value of type string.
	// Initialize the map with 2 key/value pairs.
	dict := map[string]string{"Red": "#da1337", "Orange": "#e95a22"}
	dict["Blue"] = "#0000FF"

	fmt.Println(dict)
	// map[Red:#da1337 Orange:#e95a22 Blue:#0000FF]
}
