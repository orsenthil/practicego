// [Timers](timers) are for when you want to do
// something once in the future - _tickers_ are for when
// you want to do something repeatedly at regular
// intervals. Here's an example of a ticker that ticks
// periodically until we stop it.

package main

import (
	"fmt"
	"time"
)

func main() {

	// Tickers use a similar mechanism to timers: a
	// channel that is sent values. Here we'll use the
	// `select` builtin on the channel to await the
	// values as they arrive every 500ms.

	// TODO: Create ticker with 500 milliseconds using NewTicker

	// TODO: Create done channel of bool

	// TODO: Creat a goroutine that receives from ticker.C and prints the result

	// Tickers can be stopped like timers. Once a ticker
	// is stopped it won't receive any more values on its
	// channel. We'll stop ours after 1600ms.

	// TODO: Give the ticker enough time to fire, if it ever
	// was going to, to show it is in fact stopped.

	// TODO: If done is true, print "Ticker stopped"

}