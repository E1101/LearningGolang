// imagine you have a directory listing for a user. The listing contains
// information about a user and their activity that can be viewed by many other users.
// Obtaining the dataset to render would require multiple data source lookups. If that
// were cached, you could skip loading this information each time the page is viewed.
package main

import (
	"bytes"
	"html/template"
	"net/http"
)

var t *template.Template
var qc template.HTML

func init() {
	t = template.Must(template.ParseFiles("index.html", "quote.html"))
}

type Page struct {
	Title   string
	Content template.HTML
}

type Quote struct {
	Quote, Name string
}

func main() {
	q := &Quote{
		Quote: `You keep using that word. I do not think
		dataset to supply
		it means what you think it means.`,
		Name: "Inigo Montoya",
	}

	var b bytes.Buffer
	t.ExecuteTemplate(&b, "quote.html", q)
	qc = template.HTML(b.String())

	http.HandleFunc("/", displayPage)
	http.ListenAndServe(":8080", nil)
}

func displayPage(w http.ResponseWriter, r *http.Request) {
	p := &Page{
		Title: "A User",
		Content: qc,
	}

	t.ExecuteTemplate(w, "index.html", p)
}
