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
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {

	// In order to use our pool of workers we need to send
	// them work and collect their results. We make 2
	// channels for this.

	// Create numJobs = 5
	numJobs := 5

	// Create jobs channel of int with buffer size numJobs
	jobs := make(chan int, numJobs)

	// Create results channel of int with buffer size numJobs
	results := make(chan int, numJobs)

	// This starts up 3 workers, initially blocked
	// because there are no jobs yet.

	// Create 3 workers using worker function
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// Here we send 5 `jobs` and then `close` that
	// channel to indicate that's all the work we have.

	// Send 5 jobs to the workers
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// Finally we collect all the results of the work.
	// This also ensures that the worker goroutines have
	// finished. An alternative way to wait for multiple
	// goroutines is to use a [WaitGroup](waitgroups).

	// Receive from results channel
	for a := 1; a <= numJobs; a++ {
		<-results
	}

}