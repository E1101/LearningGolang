package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"golang.org/x/text/encoding/charmap"
)

func main() {

	// Write the string
	// encoded to Windows-1252
	encoder := charmap.Windows1252.NewEncoder()
	s, e := encoder.String("This is sample text with runes Š")
	if e != nil {
		panic(e)
	}

	ioutil.WriteFile("example.txt", []byte(s), os.ModePerm)

	// Decode to UTF-8
	f, e := os.Open("example.txt")
	if e != nil {
		panic(e)
	}
	defer f.Close()


	// Go, by default, expects that the strings used in
	// the program are UTF-8 based. If they are not, then
	// decoding from the given charset must be done to be
	// able to work with the string.
	decoder := charmap.Windows1252.NewDecoder()
	reader := decoder.Reader(f)
	b, err := ioutil.ReadAll(reader)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
}
