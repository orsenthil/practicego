// In this example we'll look at how to implement
// a _worker pool_ using goroutines and channels.

package main

import (
	"fmt"
	"time"
)

// Here's the worker, of which we'll run several
// concurrent instances. These workers will receive
// work on the `jobs` channel and send the corresponding
// results on `results`. We'll sleep a second per job to
// simulate an expensive task.

// TODO: Define function worker(id int, jobs <-chan int, results chan<- int) that receives from jobs and sends the result to results

func main() {

	// In order to use our pool of workers we need to send
	// them work and collect their results. We make 2
	// channels for this.

	// TODO: Create numJobs = 5

	// TODO: Create jobs channel of int with buffer size numJobs

	// TODO: Create results channel of int with buffer size numJobs

	// This starts up 3 workers, initially blocked
	// because there are no jobs yet.

	// TODO: Create 3 workers using worker function

	// Here we send 5 `jobs` and then `close` that
	// channel to indicate that's all the work we have.

	// TODO: Send 5 jobs to the workers
	// TODO: Close the jobs channel

	// Finally we collect all the results of the work.
	// This also ensures that the worker goroutines have
	// finished. An alternative way to wait for multiple
	// goroutines is to use a [WaitGroup](waitgroups).

	// TODO: Receive from results channel

}