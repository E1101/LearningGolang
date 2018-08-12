package main

import (
	"io"
	"log"
	"os"
	"os/exec"
)


// The example shows the piping output
// from the executed command to the standard
// output of the parent program.
func main() {
	// creates the in-memory pipe and
	// returns both ends of the pipe
	//
	// Each Write to PipeWriter is blocked
	// until it is consumed by Read on the other end.
	pReader, pWriter := io.Pipe()

	cmd := exec.Command("echo", "Hello Go!\nThis is example")
	cmd.Stdout = pWriter


	// By assigning the pWriter to cmd.Stdout,
	// the standard output of the child process
	// is written to the pipe, and the io.Copy
	// in goroutine consumes the written data,
	// by copying the data to os.Stdout.
	go func() {
		defer pReader.Close()

		if _, err := io.Copy(os.Stdout, pReader); err != nil {
			log.Fatal(err)
		}
	}()

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

}
