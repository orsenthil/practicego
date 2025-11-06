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
	fmt.Println("Practicing: Logging")
	log.Println("Hello from log.Println")
	log.Printf("Hello from log.Printf with %d\n", 42)
	log.Fatalln("Hello from log.Fatalln")
	log.Panicln("Hello from log.Panicln")
}
