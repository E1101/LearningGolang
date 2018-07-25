package main

import (
	"net/http"
	"time"
	"io/ioutil"
	"fmt"
	"net/url"
	"net"
	"strings"
)

func main() {
	cc := &http.Client{Timeout: time.Second}

	// Performs a GET request using the custom client
	res, err := cc.Get("http://goinpracticebook.com")
	if err != nil && hasTimedOut(err) {
		fmt.Println("A timeout error occured")
		return
	}

	b, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()

	fmt.Printf("%s", b)
}

func hasTimedOut(err error) bool {
	switch err := err.(type) {
		// A url.Error may be caused by an underlying net
		// error that can checked for a timeout.
		case *url.Error:
			if err, ok := err.Err.(net.Error); ok && err.Timeout() {
				return true
			}

		// Looks for timeouts
		// detected by the net package
		case net.Error:
			if err.Timeout() {
				return true
			}
		case *net.OpError:
			if err.Timeout() {
				return true
			}
	}

	// Some errors, without a custom
	// type or variable to check against,
	// can indicate a timeout.
	errTxt := "use of closed network connection"
	if err != nil && strings.Contains(err.Error(), errTxt) {
		return true
	}

	return false
}
