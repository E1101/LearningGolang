package main

import (
	"time"
	"context"
	"os/exec"
	"bufio"
	"fmt"
)

func main() {
	// The command line tool
	// "ping" is executed for
	// 2 seconds
	timeout := 2 * time.Second
	ctx, _ := context.WithTimeout(context.TODO(), timeout)
	proc := exec.CommandContext(ctx, "ping", "example.com")

	// The process output is obtained
	// in form of io.ReadCloser. The underlying
	// implementation use the os.Pipe
	stdout, _ := proc.StdoutPipe()
	defer stdout.Close()

	// Start the process
	proc.Start()

	// For more comfortable reading the
	// bufio.Scanner is used.
	// The read call is blocking.
	s := bufio.NewScanner(stdout)
	for s.Scan() {
		fmt.Println( s.Text() )
	}
}

