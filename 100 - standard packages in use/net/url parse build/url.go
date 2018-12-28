package main

import (
	"encoding/json"
	"fmt"
	"net/url"
)

func main() {

	// Assemble
	//
	u := &url.URL{}
	u.Scheme = "http"
	u.Host = "localhost"
	u.Path = "index.html"
	u.RawQuery = "id=1&name=John"
	u.User = url.UserPassword("admin", "1234")

	// > http://admin:1234@localhost/index.html?id=1&name=John
	fmt.Printf("Assembled URL:\n%v\n\n\n", u)


	// Parse
	//
	parsedURL, err := url.Parse( u.String() )
	if err != nil {
		panic(err)
	}

	jsonURL, err := json.Marshal(parsedURL)
	if err != nil {
		panic(err)
	}

	// > {"Scheme":"http","Opaque":"","User":{},"Host":"localhost","Path":"/index.html","RawPath":""
	// ,"ForceQuery":false,"RawQuery":"id=1\u0026name=John","Fragment":""}
	fmt.Println("Parsed URL:")
	fmt.Println( string(jsonURL) )

}
