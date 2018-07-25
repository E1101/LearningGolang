package main

import "fmt"

func main() {

	mySlice := []string{"Monday", "Tuesday"}
	mySlice = append(mySlice, []string{"Wednesday", "Thursday", "Friday"}...)

	fmt.Println(mySlice)
}
