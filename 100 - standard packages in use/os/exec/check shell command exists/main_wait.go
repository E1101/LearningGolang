package main

import (
	"fmt"
	"os/exec"
	"runtime"
	"time"
)


func main() {
	// determine command based on OS
	//
	var cmd string
	if runtime.GOOS == "windows" {
		cmd = "timeout"
	} else {
		cmd = "sleep"
	}


	proc := exec.Command(cmd, "1")
	proc.Start()

	// Wait function will
	// wait till the process ends.
	proc.Wait()

	// After the process terminates
	// the *os.ProcessState contains
	// simple information
	// about the process run
	fmt.Printf("PID: %d\n", proc.ProcessState.Pid())  // 23777

	fmt.Printf("Process took: %dms\n",
		proc.ProcessState.SystemTime()/time.Microsecond)     // 818ms

	// ProcessState structure of the os package is available,
	// only after the process terminates.
	fmt.Printf("Exited sucessfuly : %t\n",
		proc.ProcessState.Success())                         // true
}
