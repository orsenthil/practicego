// Go makes it possible to _recover_ from a panic, by
// using the `recover` built-in function. A `recover` can
// stop a `panic` from aborting the program and let it
// continue with execution instead.

// An example of where this can be useful: a server
// wouldn't want to crash if one of the client connections
// exhibits a critical error. Instead, the server would
// want to close that connection and continue serving
// other clients. In fact, this is what Go's `net/http`
// does by default for HTTP servers.

package main

import "fmt"

func mayPanic() {
	panic("a problem")
}

func main() {
	// `recover` must be called within a deferred function.
	// When the enclosing function panics, the defer will
	// activate and a `recover` call within it will catch
	// the panic.

	// Defer a function that recovers from a panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	mayPanic()
	fmt.Println("After mayPanic()")
}