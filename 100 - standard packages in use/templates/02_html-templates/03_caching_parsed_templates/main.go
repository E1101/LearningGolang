package main

import (
	"html/template"
	"net/http"
)

// Parses the template when
// the package is initialized
// This is a subtle, simple way to speed up application responses.
var t = template.Must(template.ParseFiles("templates/simple.html"))

type Page struct {
	Title, Content string
}

func diaplayPage(w http.ResponseWriter, r *http.Request) {
	p := &Page{
		Title: "An Example",
		Content: "Have fun stormin’ da castle.",
	}

	t.Execute(w, p)
}

func main() {
	http.HandleFunc("/", diaplayPage)
	http.ListenAndServe(":8080", nil)
}
