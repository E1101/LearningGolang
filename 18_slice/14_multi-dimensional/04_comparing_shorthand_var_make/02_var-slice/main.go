package main

import (
	"fmt"
)

func main() {
	var student []string
	var students [][]string


	// student[0] = "Todd"
	// it cause an error; instead must use append because student initialized to nil value
	student = append(student, "Todd")
	fmt.Println(student)
	fmt.Println(students)
}
