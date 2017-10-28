package main

import (
	"net/http"
	"fmt"
)

func main() {
	res, _ := http.Get("http://goinpracticebook.com/")

	switch {
	case 300 <= res.StatusCode && res.StatusCode < 400:
		fmt.Println("Redirect message")
	case 400 <= res.StatusCode && res.StatusCode < 500:
		fmt.Println("Client error")
	case 500 <= res.StatusCode && res.StatusCode < 600:
		fmt.Println("Server error")
	}

}
