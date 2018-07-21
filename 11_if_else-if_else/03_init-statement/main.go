package main

import "fmt"


func main() {
	b := true

	// food is initiated to "Chocolate"
	if food := "Chocolate"; b {
		fmt.Println(food)
	}

	// fmt.Println(food) food variable can`t resolved here
}
