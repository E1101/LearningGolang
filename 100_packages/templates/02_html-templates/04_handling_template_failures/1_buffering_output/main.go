package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
)

type Page struct {
	Title, Content string
}

var t *template.Template

func init() {
	t = template.Must(template.ParseFiles("./templates/simple.html"))
}

func diaplayPage(w http.ResponseWriter, r *http.Request) {
	p := &Page{
		Title: "An Example",
		Content: "Have fun stormin’ da castle.",
	}

	// Creates a buffer to store
	// the output of the executed template
	var b bytes.Buffer
	err := t.Execute(&b, p)
	if err != nil {
		// Handles any errors from
		// template execution
		fmt.Fprint(w, "A error occured.")
		return
	}

	// Copies the buffered output
	// to the response writer
	b.WriteTo(w)
}

func main() {
	http.HandleFunc("/", diaplayPage)
	http.ListenAndServe(":8080", nil)
}
