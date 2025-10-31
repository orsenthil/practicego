// [_Rate limiting_](https://en.wikipedia.org/wiki/Rate_limiting)
// is an important mechanism for controlling resource
// utilization and maintaining quality of service. Go
// elegantly supports rate limiting with goroutines,
// channels, and [tickers](tickers).

package main

import (
	"fmt"
	"time"
)

func main() {

	// First we'll look at basic rate limiting. Suppose
	// we want to limit our handling of incoming requests.
	// We'll serve these requests off a channel of the
	// same name.

	// TODO: Create requests channel of int with buffer size 5
	requests := make(chan int, 5)

	// TODO: Send 5 requests to the requests channel
	for i := 1; i <= 5; i++ {
		requests <- i
	}

	// TODO: Close the requests channel
	close(requests)
	// This `limiter` channel will receive a value
	// every 200 milliseconds. This is the regulator in
	// our rate limiting scheme.

	// TODO: Create limiter channel of time.Tick 200 milliseconds
	limiter := time.Tick(200 * time.Millisecond)

	// By blocking on a receive from the `limiter` channel
	// before serving each request, we limit ourselves to
	// 1 request every 200 milliseconds.

	// TODO: Iterate over requests channel and retrieve the limiter channel and print the request and time

	for req := range requests {
		<-limiter
		fmt.Println("request", req, "time", time.Now())
	}

	// We may want to allow short bursts of requests in
	// our rate limiting scheme while preserving the
	// overall rate limit. We can accomplish this by
	// buffering our limiter channel. This `burstyLimiter`
	// channel will allow bursts of up to 3 events.

	// TODO: Create burstyLimiter channel of time.Time with buffer size 3
	burstyLimiter := make(chan time.Time, 3)

	// Fill up the channel to represent allowed bursting.

	// TODO: Iterate over 3 and send the time to the burstyLimiter channel
	for i := 1; i <= 3; i++ {
		burstyLimiter <- time.Now()
	}
	// Every 200 milliseconds we'll try to add a new
	// value to `burstyLimiter`, up to its limit of 3.

	// TODO: Creat a goroutine that sends the time to the burstyLimiter channel every 200 milliseconds
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()
	// Now simulate 5 more incoming requests. The first
	// 3 of these will benefit from the burst capability
	// of `burstyLimiter`.

	// TODO: Create burstyRequests channel of int with buffer size 5
	burstyRequests := make(chan int, 5)

	// TODO: Send 5 requests to the burstyRequests channel
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	// TODO: Close the burstyRequests channel

	// TODO: Iterate over burstyRequests channel and retrieve the burstyLimiter channel and print the request and time
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, "time", time.Now())
	}
}
