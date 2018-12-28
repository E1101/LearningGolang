package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"
)

type StringServer string

const addr = "localhost:7070"


func main() {

	s := createServer(addr)
	go s.ListenAndServe()


	// Connect with plain TCP
	// Once the Dial function is successful,
	// the Conn type is returned, which serves
	// as a reference to the opened socket.
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()


	// request a page
	_, err = io.WriteString(conn, "GET / HTTP/1.1\r\nHost: localhost:7070\r\n\r\n")
	if err != nil {
		panic(err)
	}


	// Note that the Scanner, in this case,
	// works because of the brake lines.
	// Otherwise, the more generic Read method should be used.
	scanner := bufio.NewScanner(conn)
	// This means the deadline is set as a time point in the future.
	// time.Now().Add(10*time.Second)
	conn.SetReadDeadline( time.Now().Add(time.Second) )
	for scanner.Scan() {
		fmt.Println( scanner.Text() )
	}

	ctx, _ := context.WithTimeout( context.Background(), 5*time.Second )
	s.Shutdown(ctx)

}


// ..

func (s StringServer) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	rw.Write( []byte(string(s)) )
}


func createServer(addr string) http.Server {
	return http.Server {
		Addr: addr,
		Handler: StringServer("HELLO GOPHER!\n"),
	}
}
