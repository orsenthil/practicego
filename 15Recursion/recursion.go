// Go supports
// <a href="https://en.wikipedia.org/wiki/Recursion_(computer_science)"><em>recursive functions</em></a>.
// Here's a classic example.

package main

import "fmt"

// This `fact` function calls itself until it reaches the
// base case of `fact(0)`.
// TODO: Create recursive function fact(n int) int
// Base case: if n == 0 return 1
// Recursive case: return n * fact(n-1)

func main() {
	
	// TODO: Call fact(7) and print result

	// Anonymous functions can also be recursive, but this requires
	// explicitly declaring a variable with `var` to store
	// the function before it's defined.

	// TODO: Create variable fib of type func(int) int
	// Assign anonymous function that calculates fibonacci recursively

	// Since `fib` was previously declared, you can call that with the anonymous function


	// TODO: Call fib(7) and print result
}