// A _line filter_ is a common type of program that reads
// input on stdin, processes it, and then prints some
// derived result to stdout. `grep` and `sed` are common
// line filters.

// Here's an example line filter in Go that writes a
// capitalized version of all input text. You can use this
// pattern to write your own Go line filters.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// Wrapping the unbuffered `os.Stdin` with a buffered
	// scanner gives us a convenient `Scan` method that
	// advances the scanner to the next token; which is
	// the next line in the default scanner.

	// TODO: Create scanner := bufio.NewScanner(os.Stdin)
	scanner := bufio.NewScanner(os.Stdin)

	// TODO: For scanner.Scan(), use strings.ToUpper(scanner.Text()) to uppercase the line
	// TODO: Write out the uppercased line
	for scanner.Scan() {
		fmt.Println(strings.ToUpper(scanner.Text()))
	}	
	// Check for errors during `Scan`. End of file is
	// expected and not reported by `Scan` as an error.

	// TODO: Create err := scanner.Err()
	err := scanner.Err()
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
	// TODO: If err is not nil, print "error:", err and exit with 1
}