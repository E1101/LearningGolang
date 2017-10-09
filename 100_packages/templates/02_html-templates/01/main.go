package main

import (
	"html/template"
	"log"
	"os"
	"strings"
)

type Page struct {
	Title string
	Body  string
}

func main() {
	var err error

	tpl := template.New("tpl.gohtml")
	tpl = tpl.Funcs(template.FuncMap{
		"uppercase": func(str string) string {
			return strings.ToUpper(str)
		},
	})

	tpl, err = tpl.ParseFiles("templates/tpl.ghtml")
	if err != nil {
		log.Fatalln(err)
	}

	// Data object to pass to template containing properties to print
	p := &Page{
		Title: "My Title 2",
		Body:  `hello world <script>alert("Yow!");</script>`,
	}
	err = tpl.Execute(os.Stdout, p)
	if err != nil {
		log.Fatalln(err)
	}
}
