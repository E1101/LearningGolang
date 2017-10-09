package main

import (
	"net/http"
	"html/template"
)

type Page struct {
	Title, Content string
}

func displayPage(w http.ResponseWriter, r *http.Request) {
	p := &Page{
		Title:   "An Example",
		Content: "Have fun stormin’ da castle.",
	}

	// Parses a template for later use
	t := template.Must(template.ParseFiles("templates/simple.ghtml"))
	// Writes to HTTP output using template and dataset
	t.Execute(w, p)
}

func main() {
	http.HandleFunc("/", displayPage)
	http.ListenAndServe(":8080", nil)
}
