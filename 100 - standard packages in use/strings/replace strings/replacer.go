package main

import (
	"fmt"
	"strings"
)

const refString2 = "Mary had a little lamb"


func main() {
	replacer := strings.NewReplacer("lamb", "wolf", "Mary", "Jack")
	out := replacer.Replace(refString2)

	fmt.Println(out)
}
