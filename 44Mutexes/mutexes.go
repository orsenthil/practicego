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

type Container struct {
	mu sync.Mutex
	counters map[string]int
}

func (c *Container) inc(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}


func main() {
	
	// Note that the zero value of a mutex is usable as-is, so no
	// initialization is required here.

	c := Container{
		counters: map[string]int{"a": 0, "b": 0},
	}
	var wg sync.WaitGroup

	// This function increments a named counter
	// in a loop.

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			c.inc("a")
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(c.counters)
}