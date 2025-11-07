// The primary mechanism for managing state in Go is
// communication over channels. We saw this for example
// with [worker pools](worker-pools). There are a few other
// options for managing state though. Here we'll
// look at using the `sync/atomic` package for _atomic
// counters_ accessed by multiple goroutines.

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	// We'll use an atomic integer type to represent our
	// (always-positive) counter.

	ops := atomic.Uint64(0)

	// A WaitGroup will help us wait for all goroutines
	// to finish their work.

	wg := sync.WaitGroup{}

	// We'll start 50 goroutines that each increment the
	// counter exactly 1000 times.

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			ops.Add(1000)
			wg.Done()
		}()
	}

	// Wait until all the goroutines are done.

	wg.Wait()

	// Here no goroutines are writing to 'ops', but using
	// `Load` it's safe to atomically read a value even while
	// other goroutines are (atomically) updating it.

	fmt.Println("ops:", ops.Load())
}