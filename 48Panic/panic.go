// A `panic` typically means something went unexpectedly
// wrong. Mostly we use it to fail fast on errors that
// shouldn't occur during normal operation, or that we
// aren't prepared to handle gracefully.

package main

import "os"
import "fmt"

func main() {

	// We'll use panic throughout this site to check for
	// unexpected errors. This is the only program on the
	// site designed to panic.

	// Panic with "a problem"
	panic("a problem")


	// A common use of panic is to abort if a function
	// returns an error value that we don't know how to
	// (or want to) handle. Here's an example of
	// `panic`king if we get an unexpected error when creating a new file.

	// Create a new file with os.Create("/tmp/file")
	f, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	write, err := f.WriteString("data")
	if err != nil {
		panic(err)
	}
	fmt.Println("wrote", write, "bytes")

}