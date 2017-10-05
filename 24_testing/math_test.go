package math

import (
	"fmt"
	"testing"
	"testing/quick"
	"time"
)

/*
$ go test

look for any tests in any of the files in the current folder
and run them.

The go test command provides a coverage (-cover) flag that helps
you to get coverage of the test cases written against your code.
*/

func TestAdder(t *testing.T) {
	result := Adder(4, 7)
	if result != 11 {
		t.Fatal("4 + 7 did not equal 11")
	}
}

func BenchmarkAdder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Adder(4, 7)
	}
}

func ExampleAdder() {
	fmt.Println(Adder(4, 7))
	// Output:
	// 11
}

func ExampleAdder_multiple() {
	fmt.Println(Adder(3, 6, 7, 4, 61))
	// Output:
	// 81
}

func TestAdderBlackbox(t *testing.T) {
	err := quick.Check(a, nil)
	if err != nil {
		t.Fatal(err)
	}
}

// The type testing.T provides a method Skip that can be used to skip unit tests.
// To skip those unit tests, you can give a signal by providing a short (-short)
// flag to the go test command.

// TestLongRun is a time-consuming test
func TestLongRun(t *testing.T) {
	// Checks whether the short flag is provided
	if testing.Short() {
		t.Skip("Skipping test in short mode")
	}

	// Long running implementation goes here
	time.Sleep(5 * time.Second)
}

func TestSkipped(t *testing.T)  {
	t.Skip("Skipping test in short mode")
}

func a(x, y int) bool {
	return Adder(x, y) == x+y
}
