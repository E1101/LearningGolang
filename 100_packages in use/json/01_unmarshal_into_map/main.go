package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var obj map[string]string

	err := json.Unmarshal([]byte(
		`{
		   "name": "Todd McLeod"
		}`), &obj)

	if err != nil {
		panic(err)
	}

	fmt.Println(obj)                       // map[name:Todd McLeod]
	fmt.Println(obj["name"])               // Todd McLeod
	fmt.Printf("%T\n", obj["name"]) // string
}
