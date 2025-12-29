// In a [previous](range-over-built-in-types) example we saw how `for` and
// `range` provide iteration over basic data structures.
// We can also use this syntax to iterate over
// values received from a channel.

package main

import "fmt"

func main() {

	// We'll iterate over 2 values in the `queue` channel.

	// TODO: Create queue channel of strings with buffer size 2

	// TODO: Send "one" and "two" to queue channel

	// TODO: Close the queue channel


	// This `range` iterates over each element as it's
	// received from `queue`. Because we `close`d the
	// channel above, the iteration terminates after
	// receiving the 2 elements.

	// TODO: Use range to iterate over queue and print each element

}