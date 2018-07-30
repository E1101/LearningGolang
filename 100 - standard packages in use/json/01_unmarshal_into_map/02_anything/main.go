package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonData := `
	{
	   "name": "Todd McLeod",
	   "age": 44
	}
	`

	type Anything interface{}

	var obj map[string]Anything

	err := json.Unmarshal([]byte(jsonData), &obj)
	if err != nil {
		panic(err)
	}

	fmt.Println(obj)
	fmt.Printf("%T\n", obj["name"]) // string
}
