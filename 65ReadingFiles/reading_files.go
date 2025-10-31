// Reading and writing files are basic tasks needed for
// many Go programs. First we'll look at some examples of
// reading files.

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// Reading files requires checking most calls for errors.
// This helper will streamline our error checks below.

// TODO: Create function check(e error) that checks if e is not nil, panic with e

func check(e error) {
	if e != nil {
		panic(e)
	}
}


func main() {

	// Perhaps the most basic file reading task is
	// slurping a file's entire contents into memory.

	// TODO: Create dat, err := os.ReadFile("/tmp/dat")
	// TODO: Print err
	// TODO: Print fmt.Print(string(dat))

	dat, err := os.ReadFile("/tmp/dat")
	check(err)
	fmt.Print(string(dat))

	// You'll often want more control over how and what
	// parts of a file are read. For these tasks, start
	// by `Open`ing a file to obtain an `os.File` value.

	// TODO: Create f, err := os.Open("/tmp/dat")
	// TODO: Print err
	f, err := os.Open("/tmp/dat")
	check(err)

	// Read some bytes from the beginning of the file.
	// Allow up to 5 to be read but also note how many
	// actually were read.

	// TODO: Create b1 := make([]byte, 5)
	// TODO: Create n1, err := f.Read(b1)
	// TODO: Print err
	// TODO: Print fmt.Printf("%d bytes: %s ", n1, string(b1[:n1]))
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))
	// You can also `Seek` to a known location in the file
	// and `Read` from there.

	// TODO: Create o2, err := f.Seek(6, io.SeekStart)
	o2, err := f.Seek(6, io.SeekStart)
	check(err)
	// TODO: Print err
	// TODO: Create b2 := make([]byte, 2)
	// TODO: Create n2, err := f.Read(b2)
	// TODO: Print err
	// TODO: Print fmt.Printf("%d bytes @ %d: ", n2, o2)
	// TODO: Print fmt.Printf("%v ", string(b2[:n2]))
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(b2[:n2]))

	// Other methods of seeking are relative to the
	// current cursor position,

	// TODO: Create _, err = f.Seek(2, io.SeekCurrent)
	// TODO: Print err
	_, err = f.Seek(2, io.SeekCurrent)
	check(err)
	// and relative to the end of the file.

	// TODO: Create _, err = f.Seek(-4, io.SeekEnd)
	// TODO: Print err
	_, err = f.Seek(-4, io.SeekEnd)	
	check(err)
	// The `io` package provides some functions that may
	// be helpful for file reading. For example, reads
	// like the ones above can be more robustly
	// implemented with `ReadAtLeast`.

	// TODO: Create o3, err := f.Seek(6, io.SeekStart)
	// TODO: Print err
	// TODO: Create b3 := make([]byte, 2)
	// TODO: Create n3, err := io.ReadAtLeast(f, b3, 2)
	// TODO: Print err
	// TODO: Print fmt.Printf("%d bytes @ %d: %s ", n3, o3, string(b3))
	o3, err := f.Seek(6, io.SeekStart)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))
	// There is no built-in rewind, but
	// `Seek(0, io.SeekStart)` accomplishes this.

	// TODO: Create _, err = f.Seek(0, io.SeekStart)
	// TODO: Print err
	_, err = f.Seek(0, io.SeekStart)
	check(err)
	// The `bufio` package implements a buffered
	// reader that may be useful both for its efficiency
	// with many small reads and because of the additional
	// reading methods it provides.

	// TODO: Create r4 := bufio.NewReader(f)
	// TODO: Create b4, err := r4.Peek(5)
	// TODO: Print err
	// TODO: Print fmt.Printf("5 bytes: %s ", string(b4))
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))
	// Close the file when you're done (usually this would
	// be scheduled immediately after `Open`ing with
	// `defer`).

	// TODO: Create f.Close()
	f.Close()
}