package main

import (
	"fmt"
	"encoding/json"
	"os"
)

func main() {
	ks := `{"jsondata": "json"}`

	// A variable instance of type
	// interface{} to hold the JSON data
	var f interface{}
	// Parses the JSON data and puts it
	// into the interface{} type variable
	err := json.Unmarshal([]byte(ks), &f)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	printJSON(f)
}


func printJSON(v interface{}) {
	switch vv := v.(type) {
	case string:
		fmt.Println("is string", vv)
	case float64:
		fmt.Println("is float64", vv)
	case []interface{}:
		fmt.Println("is an array:")
		for i, u := range vv {
			fmt.Print(i, " ")
			printJSON(u)
		}
	case map[string]interface{}:
		fmt.Println("is an object:")
		for i, u := range vv {
			fmt.Print(i, " ")
			printJSON(u)
		}
	default:
		fmt.Println("Unknown type")
	}
}
