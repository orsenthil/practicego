// When using channels as function parameters, you can
// specify if a channel is meant to only send or receive
// values. This specificity increases the type-safety of
// the program.

package main

import "fmt"

// This `ping` function only accepts a channel for sending
// values. It would be a compile-time error to try to
// receive on this channel.

// TODO: Define function ping(pings chan<- string, msg string) that sends msg to pings channel

func ping(pings chan<- string, msg string) {
	pings <- msg
}

// The `pong` function accepts one channel for receives
// (`pings`) and a second for sends (`pongs`).

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

// TODO: Define function pong(pings <-chan string, pongs chan<- string) that
// receives msg from pings channel and sends it to pongs channel

func main() {

	// TODO: Create pings channel of strings with buffer size 1
	pings := make(chan string, 1)
	// TODO: Create pongs channel of strings with buffer size 1
	pongs := make(chan string, 1)

	// TODO: Call ping(pings, "passed message")
	ping(pings, "passed message")

	// TODO: Call pong(pings, pongs)
	pong(pings, pongs)
	// TODO: Receive from pongs channel and print it
	fmt.Println(<-pongs)
}
