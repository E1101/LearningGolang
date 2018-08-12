package main

import "os"
import "bufio"

import "bytes"
import "fmt"
import "io/ioutil"


func main() {

	// Read as reader
	//
	fmt.Println("### Read as reader ###")
	f, err := os.Open("temp/file.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Read the
	// file with reader
	wr := bytes.Buffer{}
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		wr.WriteString( sc.Text() )
	}
	fmt.Println( wr.String() )



	// ReadFile
	// for smaller files
	//
	fmt.Println("### ReadFile ###")
	fContent, err := ioutil.ReadFile("temp/file.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println( string(fContent) )

}