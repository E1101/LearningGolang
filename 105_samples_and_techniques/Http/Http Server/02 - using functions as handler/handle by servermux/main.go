package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", index)
	mux.HandleFunc("/welcome", welcome)
	mux.HandleFunc("/message", message)

	log.Println("Listening...")
	http.ListenAndServe(":8080", mux)
}


// ..

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	html :=
		`<doctype html>
		<html>
		<head>
		<title>Hello Gopher</title>
		</head>
		<body>
		<b>Hello Gopher!</b>
		<p>
		<a href="/welcome">Welcome</a> | <a href="/message">Message</a>
		</p>
		</body>
		</html>`
	fmt.Fprintf(w, html)
}

func welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Go Web Programming")
}

func message(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "net/http package is used to build web apps")
}
