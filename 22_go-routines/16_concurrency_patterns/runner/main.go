package runner

import (
	"log"
	"time"
	"os"


)

// timeout is the number of second the program has to finish.
const timeout = 3 * time.Second

// main is the entry point for the program.
func main() {
	log.Println("Starting work.")

	// Create a new timer value for this run.
	r := New(timeout)
	// Add the tasks to be run.
	r.Add(createTask(), createTask(), createTask())

	// Run the tasks and handle the result.
	if err := r.Start(); err != nil {
		switch err {
			case ErrTimeout:
				log.Println("Terminating due to timeout.")
				os.Exit(1)
			case ErrInterrupt:
				log.Println("Terminating due to interrupt.")
				os.Exit(2)
			}
	}

	log.Println("Process ended.")
}

// createTask returns an example task that sleeps for the specified
// number of seconds based on the id.
func createTask() func(int) {
	return func(id int) {
		log.Printf("Processor - Task #%d.", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}
