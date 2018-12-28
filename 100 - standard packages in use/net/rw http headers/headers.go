package main

import (
	"fmt"
	"net/http"
)


func main() {

	header := http.Header{}

	// Using the header as slice
	header.Set("Auth-X", "abcdef1234")
	header.Add("Auth-X", "defghijkl")
	fmt.Println(header)
	// > map[Auth-X:[abcdef1234 defghijkl]]


	// retrieving slice of values in header
	resSlice := header["Auth-X"]
	fmt.Println(resSlice)
	// > [abcdef1234 defghijkl]


	// get the first value
	resFirst := header.Get("Auth-X")
	fmt.Println(resFirst)
	// > abcdef1234


	// replace all existing values with
	// this one
	header.Set("Auth-X", "newvalue")
	fmt.Println(header)

	// Remove header
	header.Del("Auth-X")
	fmt.Println(header)

}
