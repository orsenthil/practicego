// Writing files in Go follows similar patterns to the
// ones we saw earlier for reading.

package main

import (
	"bufio"
	"fmt"
	"os"
)

// TODO: Create function check(e error) that checks if e is not nil, panic with e
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// To start, here's how to dump a string (or just
	// bytes) into a file.

	// TODO: Create d1 := []byte("hello go ")
	// TODO: Create err := os.WriteFile("/tmp/dat1", d1, 0644)
	// TODO: Print err
	d1 := []byte("hello go ")
	err := os.WriteFile("/tmp/dat1", d1, 0644)
	fmt.Println("err:", err)
	// For more granular writes, open a file for writing.

	// TODO: Create f, err := os.Create("/tmp/dat2")
	// TODO: Print err
	f, err := os.Create("/tmp/dat2")
	fmt.Println("f:", f)
	fmt.Println("err:", err)

	// It's idiomatic to defer a `Close` immediately
	// after opening a file.

	// TODO: Defer the closing of the file with defer f.Close()
	defer f.Close()
	// You can `Write` byte slices as you'd expect.

	// TODO: Create d2 := []byte{115, 111, 109, 101, 10}
	// TODO: Create n2, err := f.Write(d2)
	// TODO: Print err
	// TODO: Print fmt.Printf("wrote %d bytes ", n2)

	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	fmt.Println("n2:", n2)
	fmt.Println("err:", err)
	fmt.Printf("wrote %d bytes ", n2)

	// A `WriteString` is also available.

	// TODO: Create n3, err := f.WriteString("writes ")
	// TODO: Print err
	// TODO: Print fmt.Printf("wrote %d bytes ", n3)
	n3, err := f.WriteString("writes ")
	fmt.Println("n3:", n3)
	fmt.Println("err:", err)
	fmt.Printf("wrote %d bytes ", n3)

	// Issue a `Sync` to flush writes to stable storage.

	// TODO: Create f.Sync()
	f.Sync()
	// `bufio` provides buffered writers in addition
	// to the buffered readers we saw earlier.

	// TODO: Create w := bufio.NewWriter(f)
	// TODO: Print err
	// TODO: Print fmt.Printf("wrote %d bytes ", n4)
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered ")
	fmt.Println("n4:", n4)
	fmt.Println("err:", err)
	fmt.Printf("wrote %d bytes ", n4)

	// Use `Flush` to ensure all buffered operations have
	// been applied to the underlying writer.

	// TODO: Create w.Flush()
	w.Flush()
}