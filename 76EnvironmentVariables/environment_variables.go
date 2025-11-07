// [Environment variables](https://en.wikipedia.org/wiki/Environment_variable)
// are a universal mechanism for [conveying configuration
// information to Unix programs](https://www.12factor.net/config).
// Let's look at how to set, get, and list environment variables.

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	// To set a key/value pair, use `os.Setenv`. To get a
	// value for a key, use `os.Getenv`. This will return
	// an empty string if the key isn't present in the
	// environment.

	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))
	// TODO: Create fmt.Println("FOO:", os.Getenv("FOO"))
	// TODO: Create fmt.Println("BAR:", os.Getenv("BAR"))

	// Use `os.Environ` to list all key/value pairs in the
	// environment. This returns a slice of strings in the
	// form `KEY=value`. You can `strings.SplitN` them to
	// get the key and value. Here we print all the keys.
	fmt.Println()
	// TODO: iterate over os.Environ() with range and print the key
	// Use strings.SplitN(e, "=", 2) to get the key and value
}