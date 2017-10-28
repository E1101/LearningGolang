package main

import "net/http"

func displayError(w http.ResponseWriter, r *http.Request) {
	// Returns an HTTP status
	// 403 with a message
	http.Error(w, "An Error Occurred", http.StatusForbidden)
}

func main() {
	// Sets up all paths to serve the
	// HTTP handler displayError
	http.HandleFunc("/", displayError)
	http.ListenAndServe(":8080", nil)
}
