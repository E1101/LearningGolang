package main

import (
	"fmt"
	"regexp"
)

const refString3 = "Mary had a little lamb"


func main() {
	regex := regexp.MustCompile("l[a-z]+")
	out := regex.ReplaceAllString(refString3, "replacement")

	fmt.Println(out) // Mary had a replacement replacement
}
