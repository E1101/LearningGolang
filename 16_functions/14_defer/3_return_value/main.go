package main

import "errors"

// Returning values from a deferred func practically has no effect on the caller.
// However, you can still use named result values to change the result values.
func main() {

	release() // nil


	releaseWithNamedValue() // "error"

}


// ..

func release() error {
	defer func() error {
		return errors.New("error")
	}()

	return nil
}


func releaseWithNamedValue() (err error) {
	defer func() {
		err = errors.New("error")
	}()

	return nil
}
