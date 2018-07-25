package main

import (
	"net/http"
	"time"
	"fmt"
	"os"
	"io/ioutil"
)

func main() {
	// Handles any errors such as a client timeout
	// Timeouts are typically detected by the net package
	// when a timeout is explicitly set
	cc := &http.Client{Timeout: time.Second}

	// Performs a GET request using the custom client
	res, err := cc.Get("http://goinpracticebook.com")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}


	b, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()

	fmt.Printf("%s", b)
}
