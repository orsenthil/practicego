// Go supports [_anonymous functions_](https://en.wikipedia.org/wiki/Anonymous_function),
// which can form <a href="https://en.wikipedia.org/wiki/Closure_(computer_science)"><em>closures</em></a>.
// Anonymous functions are useful when you want to define
// a function inline without having to name it.

package main

import "fmt"

// This function `intSeq` returns another function, which
// we define anonymously in the body of `intSeq`. The
// returned function _closes over_ the variable `i` to
// form a closure.

// TODO: Create function intSeq() func() int
// Inside, create variable i := 0
// Return anonymous function that increments i and returns it

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

	// We call `intSeq`, assigning the result (a function)
	// to `nextInt`. This function value captures its
	// own `i` value, which will be updated each time
	// we call `nextInt`.

	// TODO: Call intSeq() and assign to nextInt
	nextInt := intSeq()
	// See the effect of the closure by calling `nextInt`
	// a few times.

	// TODO: Call nextInt() multiple times and print results for each call
	fmt.Println(nextInt())
	fmt.Println(nextInt())		
	// To confirm that the state is unique to that
	// particular function, create and test a new one.

	// TODO: Create newInts := intSeq() and call it to show separate state
	newInts := intSeq()
	fmt.Println(newInts())
	fmt.Println(newInts())	
}