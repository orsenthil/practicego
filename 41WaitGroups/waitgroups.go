// To wait for multiple goroutines to finish, we can
// use a *wait group*.

package main

import (
	"fmt"
	"sync"
	"time"
)

// This is the function we'll run in every goroutine.


func worker(id int) {
	fmt.Println("Worker", id, "starting")
	time.Sleep(time.Second)
	fmt.Println("Worker", id, "done")
}

func main() {

	// This WaitGroup is used to wait for all the
	// goroutines launched here to finish. Note: if a WaitGroup is
	// explicitly passed into functions, it should be done *by pointer*.

	wg := sync.WaitGroup{}

	// Launch several goroutines using `WaitGroup.Go`

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i)
	}

	// Block until all the goroutines started by `wg` are
	// done. A goroutine is done when the function it invokes
	// returns.

	wg.Wait()

	// Note that this approach has no straightforward way
	// to propagate errors from workers. For more
	// advanced use cases, consider using the
	// [errgroup package](https://pkg.go.dev/golang.org/x/sync/errgroup).

}