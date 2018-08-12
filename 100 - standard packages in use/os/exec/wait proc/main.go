package main

import (
	"fmt"
	"os/exec"
	"bytes"
)

func main() {
	prc := exec.Command("ls", "-a")
	out := bytes.NewBuffer([]byte{})
	prc.Stdout = out
	err := prc.Start()
	if err != nil {
		fmt.Println(err)
	}

	// The Wait method could be used to wait until the process ends.
	// After the Wait method finishes, the resources of the process are released.
	prc.Wait()

	if prc.ProcessState.Success() {
		fmt.Println("Process run successfully with output:\n")
		fmt.Println(out.String())
	}
}
