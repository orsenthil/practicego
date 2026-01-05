// A _goroutine_ is a lightweight thread of execution.

package main

import (
	"fmt"
	"time"
)

// TODO: Define function f(from string) that prints from and i in each iteration
// inside, use i:= range 3 to iterate


func main() {

	// Suppose we have a function call `f(s)`. Here's how
	// we'd call that in the usual way, running it
	// synchronously.

	// TODO: Call f("direct")


	// To invoke this function in a goroutine, use
	// `go f(s)`. This new goroutine will execute
	// concurrently with the calling one.

	// TODO: Call go f("goroutine")


	// You can also start a goroutine for an anonymous
	// function call.

	// TODO: Call go func(msg string) {
	// fmt.Println(msg)
	// }("going")

	// Our two function calls are running asynchronously in
	// separate goroutines now. Wait for them to finish
	// (for a more robust approach, use a [WaitGroup](waitgroups)).

	// TODO: Sleep for 1 second
}