package main

import (
	"io"
	"net/http"
)

func upTown(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "doggy doggy doggy")
}

func youUp(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "catty catty catty")
}

// If you provide a nil value, package http will take DefaultServeMux as the ServeMux value. When you
// work with DefaultServeMux as the ServeMux value, you can use the function http.HandleFunc to register
// a handler function for the given URL pattern. Inside the function http.HandleFunc , it calls the function
// HandleFunc of DefaultServeMux .
func main() {

	http.HandleFunc("/", upTown)
	http.HandleFunc("/cat/", youUp)

	http.ListenAndServe(":9000", nil)
}
