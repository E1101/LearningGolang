package main

import "fmt"

func main() {
	switch "Marcus" {
	case "Tim":
		fmt.Println("Wassup Tim")
	case "Jenny":
		fmt.Println("Wassup Jenny")
	case "Marcus":
		fmt.Println("Wassup Marcus")
		// fall execution to next case v
		fallthrough
	case "Medhi":
		fmt.Println("Wassup Medhi")
		// fall execution to next case v
		fallthrough
	case "Julian":
		fmt.Println("Wassup Julian")
		// execution stop after this _
	case "Sushant":
		fmt.Println("Wassup Sushant")
	}
}
