// To wait for multiple goroutines to finish, we can
// use a *wait group*.

package main

import (
	"fmt"
	"sync"
	"time"
)

// This is the function we'll run in every goroutine.

// TODO: Define function worker(id int) that prints "Worker %d starting" and "Worker %d done" after sleeping for 1 second

func worker(id int) {
	fmt.Println("Worker", id, "starting")
	time.Sleep(time.Second)
	fmt.Println("Worker", id, "done")
}
func main() {

	// This WaitGroup is used to wait for all the
	// goroutines launched here to finish. Note: if a WaitGroup is
	// explicitly passed into functions, it should be done *by pointer*.

	// TODO: Create wg sync.WaitGroup
	wg := sync.WaitGroup{}
	// Launch several goroutines using `WaitGroup.Go`

	// TODO: Launch 5 workers using worker function

	for i := 0; i < 5; i++ {
		wg.Go(func() {
			worker(i)
		})
	}
	// Block until all the goroutines started by `wg` are
	// done. A goroutine is done when the function it invokes
	// returns.

	// TODO: Wait for all the goroutines to finish
	wg.Wait()

	// Note that this approach has no straightforward way
	// to propagate errors from workers. For more
	// advanced use cases, consider using the
	// [errgroup package](https://pkg.go.dev/golang.org/x/sync/errgroup).

}