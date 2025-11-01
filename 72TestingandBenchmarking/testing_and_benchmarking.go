// Unit testing is an important part of writing
// principled Go programs. The `testing` package
// provides the tools we need to write unit tests
// and the `go test` command runs tests.

// For the sake of demonstration, this code is in package
// `main`, but it could be any package. Testing code
// typically lives in the same package as the code it tests.
package main

import (
	"fmt"
	"testing"
)

// TODO: Create function IntMin(a, b int) int that returns the minimum of a and b
func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// A test is created by writing a function with a name
// beginning with `Test`.

// TODO: Create function TestIntMinBasic(t *testing.T) that tests IntMin(2, -2) = -2
func TestIntMinBasic(t *testing.T) {
	ans := IntMin(2, -2)
	if ans != -2 {
		t.Errorf("IntMin(2, -2) = %d; want -2", ans)
	}
}

// `t.Error*` will report test failures but continue
// executing the test. `t.Fatal*` will report test
// failures and stop the test immediately.

func TestIntMinBasic(t *testing.T) {
	// TODO: Create ans := IntMin(2, -2)
	// TODO: If ans != -2, t.Errorf("IntMin(2, -2) = %d; want -2", ans)
	ans := IntMin(2, -2)
	if ans != -2 {
		t.Errorf("IntMin(2, -2) = %d; want -2", ans)
	}
}

// Writing tests can be repetitive, so it's idiomatic to
// use a *table-driven style*, where test inputs and
// expected outputs are listed in a table and a single loop
// walks over them and performs the test logic.
func TestIntMinTableDriven(t *testing.T) {

	// TODO: Create tests = []struct {
	// a, b int
	// want int
	//}{
	// {0, 1, 0},
	// {1, 0, 0},
	// {2, -2, -2},
	// {0, -1, -1},
	// {-1, 0, -1},
	//}
	tests := []struct {
		a, b int
		want int
	}{
		{0, 1, 0},
		{1, 0, 0},
		{2, -2, -2},
		{0, -1, -1},
		{-1, 0, -1},
	}

	// TODO: For _, tt := range tests, create testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
	// TODO: t.Run(testname, func(t *testing.T) {
	// TODO: Create ans := IntMin(tt.a, tt.b)
	// TODO: If ans != tt.want, t.Errorf("got %d, want %d", ans, tt.want)
	// TODO: })
	for _, tt := range tests {
		testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := IntMin(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

// Benchmark tests typically go in `_test.go` files and are
// named beginning with `Benchmark`.
// Any code that's required for the benchmark to run but should
// not be measured goes before this loop.

// TODO: Create function BenchmarkIntMin(b *testing.B) that benchmarks IntMin(1, 2)
func BenchmarkIntMin(b *testing.B) {
	for b.Loop() {
		IntMin(1, 2)
	}
}