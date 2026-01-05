// A common requirement in programs is getting the number
// of seconds, milliseconds, or nanoseconds since the
// [Unix epoch](https://en.wikipedia.org/wiki/Unix_time).
// Here's how to do it in Go.

package main

import (
	"fmt"
	"time"
)

func main() {

	// Use `time.Now` with `Unix`, `UnixMilli` or `UnixNano`
	// to get elapsed time since the Unix epoch in seconds,
	// milliseconds or nanoseconds, respectively.

	// TODO: Create now := time.Now() and print it

	// TODO: Print now.Unix()
	// TODO: Print now.UnixMilli()
	// TODO: Print now.UnixNano()

	// You can also convert integer seconds or nanoseconds
	// since the epoch into the corresponding `time`.

	// TODO: Print time.Unix(now.Unix(), 0)
	// TODO: Print time.Unix(0, now.UnixNano())
}