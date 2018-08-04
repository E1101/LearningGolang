package calc

import "testing"

// Test case for function Sum
func TestSum(t *testing.T) {
	input, expected := []int{7, 8, 10}, 25
	result := Sum(input...)
	if result != expected {
		t.Errorf("Result: %d, Expected: %d", result, expected)
	}
}

// Test case for function Average
func TestAverage(t *testing.T) {
	input, expected := []int{7, 8, 10}, 8.33
	result := Average(input...)
	if result != expected {
		t.Errorf("Result: %f, Expected: %f", result, expected)
	}
}

// Benchmark for function Sum
func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum(7, 8, 10)
	}
}

// Benchmark for function Average
func BenchmarkAverage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Average(7, 8, 10)
	}
}

// $ go test -v -bench .
//

