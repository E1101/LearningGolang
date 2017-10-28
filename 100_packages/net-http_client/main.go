package net_http_client

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// Performs a GET request
	res, _ := http.Get("http://goinpracticebook.com")
	// Reads the body of the response and closes
	// the Body reader when done reading it
	b, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()

	// Prints the body to Standard Output
	fmt.Printf("%s", b)


	// Creates a new request object set up for a delete HTTP method
	req, _ := http.NewRequest("DELETE", "http://example.com/foo/bar", nil)
	// Performs the request with the default client
	res, _ := http.DefaultClient.Do(req)

	// Displays the status code from performing the request
	fmt.Printf("%s", res.Status)
}
