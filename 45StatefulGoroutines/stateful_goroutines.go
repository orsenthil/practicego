// In the previous example we used explicit locking with
// [mutexes](mutexes) to synchronize access to shared state
// across multiple goroutines. Another option is to use the
// built-in synchronization features of  goroutines and
// channels to achieve the same result. This channel-based
// approach aligns with Go's ideas of sharing memory by
// communicating and having each piece of data owned
// by exactly 1 goroutine.

package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

// In this example our state will be owned by a single
// goroutine. This will guarantee that the data is never
// corrupted with concurrent access. In order to read or
// write that state, other goroutines will send messages
// to the owning goroutine and receive corresponding
// replies. These `readOp` and `writeOp` `struct`s
// encapsulate those requests and a way for the owning
// goroutine to respond.

// TODO: Define struct readOp with key (int) and resp (chan int) fields
type readOp struct {
	key  int
	resp chan int
}

// TODO: Define struct writeOp with key (int), val (int), and resp (chan bool) fields
type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func main() {

	// As before we'll count how many operations we perform.

	// TODO: Create readOps uint64
	var readOps uint64
	// TODO: Create writeOps uint64
	var writeOps uint64

	// The `reads` and `writes` channels will be used by
	// other goroutines to issue read and write requests,
	// respectively.

	// TODO: Create reads channel of readOp
	reads := make(chan readOp)
	// TODO: Create writes channel of writeOp
	writes := make(chan writeOp)
	// Here is the goroutine that owns the `state`, which
	// is a map as in the previous example but now private
	// to the stateful goroutine. This goroutine repeatedly
	// selects on the `reads` and `writes` channels,
	// responding to requests as they arrive. A response
	// is executed by first performing the requested
	// operation and then sending a value on the response
	// channel `resp` to indicate success (and the desired
	// value in the case of `reads`).

	// TODO: Create a goroutine that owns the state, which is a map as in the previous example but now private to the stateful goroutine.
	// Inside, use for range to select on reads and writes channels and perform the requested operation and send a value on the response channel resp to indicate success (and the desired value in the case of reads).
	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	// This starts 100 goroutines to issue reads to the
	// state-owning goroutine via the `reads` channel.
	// Each read requires constructing a `readOp`, sending
	// it over the `reads` channel, and then receiving the
	// result over the provided `resp` channel.

	// TODO: Iterate over 100 and create a goroutine that issues reads to the state-owning goroutine via the reads channel.
	// Inside, create read readOp with key rand.Intn(5) and resp make(chan int)
	// Send read to reads channel
	// Receive the result from read.resp
	// Add 1 to readOps
	// Sleep for 1 millisecond

	for range 100 {
		go func() {
			for {
				read := readOp{key: rand.Intn(5), resp: make(chan int)}
				reads <- read
				<-read.resp
				atomic.AddUint64(&readOps, 1)
				time.Sleep(1 * time.Millisecond)
			}
		}()
	}

	// We start 10 writes as well, using a similar
	// approach.

	// TODO: Iterate over 10 and create a goroutine that issues writes to the state-owning goroutine via the writes channel.
	// Inside, create write writeOp with key rand.Intn(5) and val rand.Intn(100) and resp make(chan bool)
	// Send write to writes channel
	// Receive the result from write.resp
	// Add 1 to writeOps
	// Sleep for 1 millisecond

	for range 10 {
		go func() {
			for {
				write := writeOp{key: rand.Intn(5), val: rand.Intn(100), resp: make(chan bool)}
				writes <- write
				<-write.resp
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(1 * time.Millisecond)
			}
		}()
	}

	// Let the goroutines work for a second.

	// TODO: Sleep for 1 second
	time.Sleep(1 * time.Second)
	// Finally, capture and report the op counts.

	// TODO: Print the result of the reads with readOps
	fmt.Println("readOps:", readOps)
	// TODO: Print the result of the writes with writeOps

	fmt.Println("writeOps:", writeOps)
}
