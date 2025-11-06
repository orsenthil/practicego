// In a [previous](range-over-built-in-types) example we saw how `for` and
// `range` provide iteration over basic data structures.
// We can also use this syntax to iterate over
// values received from a channel.

package main

import "fmt"

func main() {

	// We'll iterate over 2 values in the `queue` channel.

	// Create queue channel of strings with buffer size 2
	queue := make(chan string, 2)

	// Send "one" and "two" to queue channel
	queue <- "one"
	queue <- "two"

	// Close the queue channel
	close(queue)


	// This `range` iterates over each element as it's
	// received from `queue`. Because we `close`d the
	// channel above, the iteration terminates after
	// receiving the 2 elements.

	// Use range to iterate over queue and print each element
	for msg := range queue {
		fmt.Println(msg)
	}

}