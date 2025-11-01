// _Closing_ a channel indicates that no more values
// will be sent on it. This can be useful to communicate
// completion to the channel's receivers.

package main

import "fmt"

// In this example we'll use a `jobs` channel to
// communicate work to be done from the `main()` goroutine
// to a worker goroutine. When we have no more jobs for
// the worker we'll `close` the `jobs` channel.
func main() {
	
	// TODO: Create jobs channel of int with buffer size 5
	jobs := make(chan int, 5)

	// TODO: Create done channel of bool
	done := make(chan bool)

	// Here's the worker goroutine. It repeatedly receives
	// from `jobs` with `j, more := <-jobs`. In this
	// special 2-value form of receive, the `more` value
	// will be `false` if `jobs` has been `close`d and all
	// values in the channel have already been received.
	// We use this to notify on `done` when we've worked
	// all our jobs.

	// TODO: Creat a goroutine that receives from jobs and prints the result
	// or prints "received all jobs" if the operation takes more than 1 second

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()


	// This sends 3 jobs to the worker over the `jobs`
	// channel, then closes it.

	// TODO: Send 3 jobs to the worker over the jobs channel

	for j := 1; j <= 3; j++ {
		jobs <- j
	}

	// TODO: Close the jobs channel
	close(jobs)

	fmt.Println("sent all jobs")

	// We await the worker using the
	// [synchronization](channel-synchronization) approach
	// we saw earlier.

	// TODO: Receive from done channel
	<-done
	// Reading from a closed channel succeeds immediately,
	// returning the zero value of the underlying type.
	// The optional second return value is `true` if the
	// value received was delivered by a successful send
	// operation to the channel, or `false` if it was a
	// zero value generated because the channel is closed
	// and empty.
	
	// TODO: Receive from jobs
	_, ok := <-jobs

	// TODO: Print the result
	fmt.Println(ok)

}