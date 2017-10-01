// The panic function is often called in situations in which continuing the execution
// is almost impossible.
// For example, if you are trying to connect to a database and are unable to,
// then continuing the execution of the program doesn’t make any sense because
// your application depends on the database.
//
// all deferred functions are executed before stopping the execution.
package main

import (
	"io/ioutil"
	"os"
)

func ReadFile(filename string) ([]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		panic (err) // calls panic
	}
	defer f.Close()

	return ioutil.ReadAll(f)
}

func main()  {
	ReadFile(`nofile.txt`)
}
