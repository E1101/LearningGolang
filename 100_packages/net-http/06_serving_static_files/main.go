package main

import (
	"net/http"
)

func main() {
	dir := http.Dir("./files")
	// It’s capable of looking at the If-Modified-Since HTTP header and responding with a
	// 304 Not Modified response if the version of the file a user already has matches the one
	// currently being served.
	http.ListenAndServe(":8080", http.FileServer(dir))
}
