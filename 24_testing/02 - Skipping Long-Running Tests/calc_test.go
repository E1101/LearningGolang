package calc

import (
	"testing"
	"time"
)

// Test case for the Sum function
func TestSum(t *testing.T) {
	input, expected := []int{7, 8, 10}, 25
	result := Sum(input...)
	if result != expected {
		t.Errorf("Result: %d, Expected: %d", result, expected)
	}
}

// Test case for the Sum function
func TestAverage(t *testing.T) {
	input, expected := []int{7, 8, 10}, 8.33
	result := Average(input...)
	if result != expected {
		t.Errorf("Result: %f, Expected: %f", result, expected)
	}
}

// TestLongRun is a time-consuming test
func TestLongRun(t *testing.T) {
	// Checks whether the short flag is provided  <--- [check short flag]
	//
	if testing.Short() {
		t.Skip("Skipping test in short mode")
	}
	// Long running implementation goes here
	time.Sleep(5 * time.Second)
}


// $ go test -v -short
//
