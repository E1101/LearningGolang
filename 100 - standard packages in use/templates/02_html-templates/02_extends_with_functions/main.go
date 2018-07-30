package main

import (
	"html/template"
	"net/http"
	"time"
)

// An HTML Template as String
// Pipes Date through the dateFormat command
var tpl = `<!DOCTYPE HTML>
<html>
<head>
<meta charset="utf-8">
<title>Date Example</title>
</head>
<body>
<p>{{.Date | dateFormat "Jan 2, 2006"}}</p>
</body>
</html>`

// Maps Go functions to template functions
var funcMap = template.FuncMap{
	"dateFormat": dateFormat,
}

func dateFormat(layout string, d time.Time) string {
	return d.Format(layout)
}

func serveTemplate(res http.ResponseWriter, req *http.Request) {
	// Creates a new template.Template instance
	t := template.New("date")
	// Passes additional functions in
	// map into template engine
	t.Funcs(funcMap)
	// Parses the template string
	// into the template engine
	t.Parse(tpl)

	// Creates a dataset to pass into
	// template to display
	data := struct{ Date time.Time }{
		Date: time.Now(),
	}

	// Sends template with data to output response
	t.Execute(res, data)
}

func main() {
	http.HandleFunc("/", serveTemplate)
	http.ListenAndServe(":8080", nil)
}
