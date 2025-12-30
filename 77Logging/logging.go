// The Go standard library provides straightforward
// tools for outputting logs from Go programs, with the
// log package for free-form output and the log/slog
// package for structured output.

package main

import (
	"fmt"
	"log"
)

func main() {
	log.Println("Practicing: Logging")
	fmt.Println("Practicing: Logging")
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("This will be logged with the file and line number")
}