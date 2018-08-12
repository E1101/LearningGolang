package _1

import (
	"flag"
	"fmt"
)

func main() {
	// The user input for the filename flag (-filename ) is stored
	// in the pointer fileName, with type *string.
	fileName := flag.String("filename", "logfile", "File name for the log file")
	logLevel := flag.Int("loglevel", 0, "An integer value for Level (0-4)")
	isEnable := flag.Bool("enable", false, "A boolean value for enabling log options")


	// If you want to bind the flag to an existing variable, you can
	// use the functions flag.IntVar, flag.BoolVar, and flag.StringVar.
	// Bind the flag to a variable.
	var num int
	flag.IntVar(&num, "num", 25, "An integer value")


	// Parse parses flag definitions from the argument list.
	flag.Parse()

	// Get the values from pointers
	fmt.Println("filename:", *fileName)
	fmt.Println("loglevel:", *logLevel)
	fmt.Println("enable:", *isEnable)

	// Get the value from a variable
	fmt.Println("num:", num)

	// Args returns the non-flag command-line arguments.
	args := flag.Args()
	if len(args) > 0 {
		fmt.Println("The non-flag command-line arguments are:")
		// Print the arguments
		for _, v := range args {
			fmt.Println(v)
		}
	}
}

// Build and Run Program:
// the flag -h or --help provides help for the use of the command-line program.
//
//
// $ ./flag -filename=applog -loglevel=2 -enable -num=50 10 20 30 test
//   filename: applog
//   loglevel: 2
//   enable: true
//   num: 50
//   The non-flag command-line arguments are:
//   10
//   20
//   30
//   test
//
