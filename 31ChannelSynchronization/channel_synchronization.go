// We can use channels to synchronize execution
// across goroutines. Here's an example of using a
// blocking receive to wait for a goroutine to finish.
// When waiting for multiple goroutines to finish,
// you may prefer to use a [WaitGroup](waitgroups).

package main

import (
	"fmt"
	"time"
)

// This is the function we'll run in a goroutine. The
// `done` channel will be used to notify another
// goroutine that this function's work is done.

// TODO: Define function worker(done chan bool) that prints "working..." and sleeps for 1 second, 
// then prints "done" and sends true to the done channel.


func main() {

	// Start a worker goroutine, giving it the channel to
	// notify on.

	// TODO: Create done channel of bool with buffer size 1

	// TODO: Call go worker(done)

	// Block until we receive a notification from the
	// worker on the channel.

	// TODO: Receive from done channel
}