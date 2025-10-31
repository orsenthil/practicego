// _Defer_ is used to ensure that a function call is
// performed later in a program's execution, usually for
// purposes of cleanup. `defer` is often used where e.g.
// `ensure` and `finally` would be used in other languages.

package main

import (
	"fmt"
	"os"
)

// Suppose we wanted to create a file, write to it,
// and then close when we're done. Here's how we could
// do that with `defer`.
func main() {

	// Immediately after getting a file object with
	// `createFile`, we defer the closing of that file
	// with `closeFile`. This will be executed at the end
	// of the enclosing function (`main`), after
	// `writeFile` has finished.

	// TODO: Create a file with createFile("/tmp/defer.txt")
	// TODO: Defer the closing of the file with defer closeFile(f)
	// TODO: Write to the file with writeFile(f)

	f := createFile("/tmp/defer.txt")
	defer closeFile(f)
	writeFile(f)

}

// TODO: Create function createFile(p string) *os.File that creates a file at the given path and returns the file
// Inside, create a file with os.Create(p) and check if err is not nil, panic with err
// Return the file

func createFile(p string) *os.File {
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

// TODO: Create function writeFile(f *os.File) that writes "data" to the file
// Inside, use fmt.Fprintln(f, "data")

func writeFile(f *os.File) {
	fmt.Fprintln(f, "data")
}

// TODO: Create function closeFile(f *os.File) that closes the file
// Inside, use f.Close() and check if err is not nil, panic with err

// It's important to check for errors when closing a
// file, even in a deferred function.

func closeFile(f *os.File) {
	err := f.Close()
	if err != nil {
		panic(err)
	}
}