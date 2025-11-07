// Sometimes we'd like our Go programs to intelligently
// handle [Unix signals](https://en.wikipedia.org/wiki/Unix_signal).
// For example, we might want a server to gracefully
// shutdown when it receives a `SIGTERM`, or a command-line
// tool to stop processing input if it receives a `SIGINT`.
// Here's how to handle signals in Go with channels.

package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	// Go signal notification works by sending `os.Signal`
	// values on a channel. We'll create a channel to
	// receive these notifications. Note that this channel
	// should be buffered.

	// TODO: Create sigs channel of os.Signal with buffer size 1
	sigs := make(chan os.Signal, 1)

	// `signal.Notify` registers the given channel to
	// receive notifications of the specified signals.

	// TODO: Call signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	// We could receive from `sigs` here in the main
	// function, but let's see how this could also be
	// done in a separate goroutine, to demonstrate
	// a more realistic scenario of graceful shutdown.

	// TODO: Create done channel of bool with buffer size 1
	done := make(chan bool, 1)
	go func() {
		sig := <-sigs
		fmt.Println(sig)
		done <- true
	}()
	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
}