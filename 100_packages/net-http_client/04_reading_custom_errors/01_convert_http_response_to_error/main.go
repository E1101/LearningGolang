package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"os"
)

func main() {
	res, err := get("http://localhost:8080")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	b, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s", b)
}


// ..

func get(u string) (*http.Response, error) {

	res, err := http.Get(u)
	if err != nil {
		return res, err
	}

	// Checks if the response code was
	// outside the 200 range of
	// successful responses
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		// Checks the response content type and
		// returns an error if it’s not correct
		if res.Header.Get("Content-Type") != "application/json" {
			sm := "Unknown error. HTTP status: %s"
			return res, fmt.Errorf(sm, res.Status)
		}

		// Reads the body of the
		// response into a buffer
		b, _ := ioutil.ReadAll(res.Body)
		res.Body.Close()

		// Parses the JSON response and
		// places the data into a struct and
		// responds to any errors
		var data struct {
			Err Error `json:"error"` }
		err = json.Unmarshal(b, &data)
		if err != nil {
			sm := "Unable to parse json: %s. HTTP status: %s"
			return res, fmt.Errorf(sm, err, res.Status)
		}

		// Adds the HTTP status code to the
		// Error instance
		data.Err.HTTPCode = res.StatusCode
		return res, data.Err
	}

	return res, nil
}


// Error:

// Structure to hold
// data from the error
type Error struct {
	HTTPCode int    `json:"-"`
	Code     int    `json:"code,omitempty"`
	Message  string `json:"message"`
}

// The Error method implements the
// error interface on the Error struct.
func (e Error) Error() string {
	fs := "HTTP: %d, Code: %d, Message: %s"
	return fmt.Sprintf(fs, e.HTTPCode, e.Code, e.Message)
}


