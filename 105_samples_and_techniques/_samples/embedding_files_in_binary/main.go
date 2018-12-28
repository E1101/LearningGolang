// https://github.com/GeertJohan/go.rice
//
// $ go get github.com/GeertJohan/go.rice/rice
// After this tool is installed, you can build a Go binary with the following two commands:
// $ rice embed-go
// $ go build
package main

import "net/http"

import (
	// Imports go.rice to handle
	// file locations for you
	"github.com/GeertJohan/go.rice"
	"net/http"
)

func main() {
	// Creates a box to represent a location
	// on the filesystem
	box := rice.MustFindBox("../files/")
	httpbox := box.HTTPBox()
	// An HTTPBox provides files using
	// the http.FileSystem interface
	http.ListenAndServe(":8080", http.FileServer(httpbox))
}
