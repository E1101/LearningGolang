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

	var obj map[string]interface{}

	err := json.Unmarshal([]byte(jsonData), &obj)
	if err != nil {
		panic(err)
	}

	fmt.Println(obj)                        // map[name:Todd McLeod age:44]
	fmt.Println(obj["name"])                // Todd McLeod
	fmt.Println(obj["age"])                 // 44
	fmt.Printf("%T\n", obj["name"])  // string
	fmt.Printf("%T\n", obj["age"])   // float64
}
