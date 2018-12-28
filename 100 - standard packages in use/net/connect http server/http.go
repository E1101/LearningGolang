// Communication with the HTTP server at a higher level
//
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


type StringServer string

const addr = "localhost:7070"


func main() {
	s := createServer(addr)
	go s.ListenAndServe()

	useRequest()
	simplePost()

}


// StringServer:

func (s StringServer) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	fmt.Printf("Received form data: %v\n", req.Form)
	rw.Write( []byte(string(s)) )
}


// ..

func createServer(addr string) http.Server {
	return http.Server{
		Addr: addr,
		Handler: StringServer("Hello world"),
	}
}


func useRequest() {

	hc := http.Client{}


	// Build Request
	//
	form := url.Values{}
	form.Add("name", "Radek")
	form.Add("surname", "Sohlich")

	req, err := http.NewRequest(
		"POST",
		"http://localhost:7070",
		strings.NewReader( form.Encode() ))

	req.Header.Add("Content-Type",
		"application/x-www-form-urlencoded")


	// Send Request with Client
	//
	res, err := hc.Do(req)
	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	res.Body.Close()
	fmt.Println( "Response from server:" + string(data) )
}

func simplePost() {

	res, err := http.Post(
		"http://localhost:7070",
		"application/x-www-form-urlencoded",
		strings.NewReader("name=Radek&surname=Sohlich"))
	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	res.Body.Close()
	fmt.Println("Response from server:" + string(data))
}
