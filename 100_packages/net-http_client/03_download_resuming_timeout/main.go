package main

import (
	"os"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
	"net/url"
	"net"
	"strings"
)

func main() {
	// Creates a local file to
	// store the download
	file, err := os.Create("file.zip")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Downloads the remote file to the local file,
	// retrying up to 100 times
	location := "https://example.com/file.zip"
	err = download(location, file, 100)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Displays the size of the file after
	// the download is complete
	fi, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Got it with %v bytes downloaded", fi.Size())
}


func download(location string, file *os.File, retries int64) error {
	// Creates a new GET
	// request for the file
	// being downloaded
	req, err := http.NewRequest("GET", location, nil)
	if err != nil {
		return err
	}

	// Starts the local file to find
	// the current file information
	fi, err := file.Stat()
	if err != nil {
		return err
	}

	current := fi.Size()
	if current > 0 {
		// When the local file already has
		// content, sets a header
		// requesting where the local file
		// left off. Ranges have an index of
		// 0, making the current length the
		// index for the next needed byte.
		start := strconv.FormatInt(current, 10)
		req.Header.Set("Range", "bytes="+start+"-")
	}

	// An HTTP client configured to
	// explicitly check for timeout
	cc := &http.Client{Timeout: 5 * time.Minute}
	// Performs the request for the file
	// or part if part of the file is already stored locally
	res, err := cc.Do(req)
	if err != nil && hasTimedOut(err) {
		if retries > 0 {
			// When checking for an error, tries return err
			// the request again if the error was caused by a timeout
			return download(location, file, retries-1)
		}

		return err

	} else if err != nil {
		return err
	}

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		// Handles nonsuccess HTTP status codes
		errFmt := "Unsuccess HTTP request. Status: %s"
		return fmt.Errorf(errFmt, res.Status)
	}

	// If the server doesn’t
	// support serving partial
	// files, sets retries to 0
	if res.Header.Get("Accept-Ranges") != "bytes" {
		retries = 0
	}

	// Copies the remote response
	// to the local file
	_, err = io.Copy(file, res.Body)
	if err != nil && hasTimedOut(err) {
		// If a timeout error occurs while
		// copying the file, tries retrieving
		// the remaining content
		if retries > 0 {
			return download(location, file, retries-1)
		}

		return err

		} else if err != nil {
		return err
	}

	return nil
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
