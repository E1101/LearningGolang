package calc

import (
	"testing"
	"time"
)

// Test case for the function Sum to be executed in parallel
func TestSumInParallel(t *testing.T) {
	t.Parallel() // <-------- [signal that this test is to be run in parallel]

	// Delaying 1 second for the sake of demonstration
	time.Sleep(1 * time.Second)
	input, expected := []int{7, 8, 10}, 25
	result := Sum(input...)
	if result != expected {
		t.Errorf("Result: %d, Expected: %d", result, expected)
	}
}

// Test case for the function Sum to be executed in parallel
func TestAverageInParallel(t *testing.T) {
	t.Parallel() // <-------- [signal that this test is to be run in parallel]

	// Delaying 1 second for the sake of demonstration
	time.Sleep(2 * time.Second)
	input, expected := []int{7, 8, 10}, 8.33
	result := Average(input...)
	if result != expected {
		t.Errorf("Result: %f, Expected: %f", result, expected)
	}
}
