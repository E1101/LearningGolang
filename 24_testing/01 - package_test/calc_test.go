package calc

import "testing"

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

// $ go test -v
// === RUN
// TestSum
// --- PASS: TestSum (0.00s)
// === RUN
// TestAverage
// --- PASS: TestAverage (0.00s)
// PASS ok
// github.com/shijuvar/go-recipes/ch08/calc 0.121s
