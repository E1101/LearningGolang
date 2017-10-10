package main

import (
	"html/template"
	"net/http"
)

var t *template.Template

type Page struct {
	Title, Content string
}


func init() {
	// Loads the two templates
	// into a template object
	t = template.Must(template.ParseFiles("index.html", "head.html"))
}

func diaplayPage(w http.ResponseWriter, r *http.Request) {
	p := &Page{
		Title: "An Example",
		Content: "Have fun stormin’ da castle.",
	}

	// Invokes the template
	// with the page data
	t.ExecuteTemplate(w, "index.html", p)
}

func main() {
	http.HandleFunc("/", diaplayPage)
	http.ListenAndServe(":8080", nil)
}
