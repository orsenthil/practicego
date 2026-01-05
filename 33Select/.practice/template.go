// Go's _select_ lets you wait on multiple channel
// operations. Combining goroutines and channels with
// select is a powerful feature of Go.

package main

import (
	"fmt"
	"time"
)

func main() {

	// For our example we'll select across two channels.

	// TODO: Create c1 channel of strings

	// TODO: Create c2 channel of strings


	// Each channel will receive a value after some amount
	// of time, to simulate e.g. blocking RPC operations
	// executing in concurrent goroutines.

	// TODO Creat a goroutine that sends "one" to c1 after 1 second

	// TODO: Creat a goroutine that sends "two" to c2 after 2 seconds


	// We'll use `select` to await both of these values
	// simultaneously, printing each one as it arrives.

	// TODO: Use for range 2 to receive from c1 and c2

}