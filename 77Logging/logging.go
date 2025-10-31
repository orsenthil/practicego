// The Go standard library provides straightforward
// tools for outputting logs from Go programs, with the
// log package for free-form output and the log/slog
// package for structured output.

package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"log/slog"
)

func main() {
	// TODO: Implement logging concepts
	fmt.Println("Practicing: Logging")
	log.Println("message")

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("message")

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("message")

	logger := log.New(os.Stdout, "prefix: ", log.LstdFlags)
	logger.Println("message")

	logger.SetPrefix("prefix: ")

	logger.SetOutput(os.Stderr)
	logger.Println("message")

	var buf bytes.Buffer

	logger.Println("message")
	fmt.Println(buf.String())

	jsonHandler := slog.NewJSONHandler(os.Stdout, nil)
	jsonLogger := slog.New(jsonHandler)
	jsonLogger.Info("message")
	fmt.Println(buf.String())
}