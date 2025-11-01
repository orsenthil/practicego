// In the previous example we saw how to manage simple
// counter state using [atomic operations](atomic-counters).
// For more complex state we can use a [_mutex_](https://en.wikipedia.org/wiki/Mutual_exclusion)
// to safely access data across multiple goroutines.

package main

import (
	"fmt"
	"sync"
)

// Container holds a map of counters; since we want to
// update it concurrently from multiple goroutines, we
// add a `Mutex` to synchronize access.
// Note that mutexes must not be copied, so if this
// `struct` is passed around, it should be done by
// pointer.

// TODO: Define struct Container with mu (sync.Mutex) and counters (map[string]int) fields

type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

// TODO: Create method inc(name string) on Container that locks the mutex and increments the counter for the given name
// Lock the mutex before accessing `counters`; unlock
// it at the end of the function using a [defer](defer)
// statement.

func (c *Container) inc(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

func main() {

	// Note that the zero value of a mutex is usable as-is, so no
	// initialization is required here.

	// TODO: Create c Container with counters map[string]int{"a": 0, "b": 0}

	c := Container {
		counters: map[string]int{"a":0 , "b":0},
	}

	// TODO: Create wg sync.WaitGroup
	var wg sync.WaitGroup

	// This function increments a named counter
	// in a loop.

	// TODO: Define function doIncrement(name string, n int) that increments the counter for the given name in a loop

	doIncrement := func(name string, n int) {
		for range n {
			c.inc(name)
		}
	}

	// Run several goroutines concurrently; note
	// that they all access the same `Container`,
	// and two of them access the same counter.

	// TODO: Launch 3 goroutines using wg.Go that call doIncrement with "a" and 10000, "a" and 10000, and "b" and 10000

	wg.Go(func() {
		doIncrement("a", 10000)
	})

	// Wait for the goroutines to finish

	// TODO: Wait for all the goroutines to finish.
	wg.Go(func() {
		doIncrement("a", 10000)
	})

	wg.Go(func() {
		doIncrement("b", 10000)
	})

	wg.Wait()

	// TODO: Print the result of the counters with c.counters
	fmt.Print(c.counters)
}
