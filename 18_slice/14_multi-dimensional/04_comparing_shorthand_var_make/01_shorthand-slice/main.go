package main

import (
	"fmt"
)

func main() {
	student := []string{}
	students := [][]string{}

	// student[0] = "Todd"
	// it can't be used because length is empty;
	// instead must use append
	student = append(student, "Todd")

	fmt.Println(student)
	fmt.Println(students)
}
