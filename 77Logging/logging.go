// The Go standard library provides straightforward
// tools for outputting logs from Go programs, with the
// log package for free-form output and the log/slog
// package for structured output.

package main

import (
	"fmt"
	"log"
	"log/slog"
)

func main() {
	// TODO: Implement logging concepts
	log.Println("Practicing: Logging")
	slog.Info("Practicing: Logging")
	slog.Error("Practicing: Logging")
	slog.Warn("Practicing: Logging")
	slog.Debug("Practicing: Logging")
}