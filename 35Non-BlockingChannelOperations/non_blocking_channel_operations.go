// Basic sends and receives on channels are blocking.
// However, we can use `select` with a `default` clause to
// implement _non-blocking_ sends, receives, and even
// non-blocking multi-way `select`s.

package main

import "fmt"

func main() {
	
	// Create messages channel of strings
	messages := make(chan string)

	// Create signals channel of bools
	signals := make(chan bool)

	// Here's a non-blocking receive. If a value is
	// available on `messages` then `select` will take
	// the `<-messages` `case` with that value. If not
	// it will immediately take the `default` case.

	// Use select to receive from messages and print the result
	// or print "no message received" if the operation takes more than 1 second
	select {
	case msg := <-messages:
		fmt.Println("received", msg)
	default:
		fmt.Println("no message received")
	}

	// A non-blocking send works similarly. Here `msg`
	// cannot be sent to the `messages` channel, because
	// the channel has no buffer and there is no receiver.
	// Therefore the `default` case is selected.

	// Use select to send to messages and print the result
	// or print "no message sent" as default
	select {
	case msg := <-messages:
		fmt.Println("message received", msg)
	default:
		fmt.Println("no message received")
	}
	// We can use multiple `case`s above the `default`
	// clause to implement a multi-way non-blocking
	// select. Here we attempt non-blocking receives
	// on both `messages` and `signals`.

		select {
	case msg := <-messages:
		fmt.Println("received", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}

}