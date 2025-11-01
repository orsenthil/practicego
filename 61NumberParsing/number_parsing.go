// Parsing numbers from strings is a basic but common task
// in many programs; here's how to do it in Go.

package main

// The built-in package `strconv` provides the number
// parsing.
import (
	"fmt"
	"strconv"
)

func main() {

	// With `ParseFloat`, this `64` tells how many bits of
	// precision to parse.

	// TODO: Create f, _ := strconv.ParseFloat("1.234", 64)
	// TODO: Print f
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println("f:", f)

	// For `ParseInt`, the `0` means infer the base from
	// the string. `64` requires that the result fit in 64
	// bits.

	// TODO: Create i, _ := strconv.ParseInt("123", 0, 64)
	// TODO: Print i
	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println("i:", i)

	// `ParseInt` will recognize hex-formatted numbers.

	// TODO: Create d, _ := strconv.ParseInt("0x1c8", 0, 64)
	// TODO: Print d
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println("d:", d)

	// A `ParseUint` is also available.

	// TODO: Create u, _ := strconv.ParseUint("789", 0, 64)
	// TODO: Print u
	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println("u:", u)

	// `Atoi` is a convenience function for basic base-10
	// `int` parsing.

	// TODO: Create k, _ := strconv.Atoi("135")
	// TODO: Print k
	k, _ := strconv.Atoi("135")
	fmt.Println("k:", k)
	// Parse functions return an error on bad input.

	// TODO: Create _, e := strconv.Atoi("wat")
	// TODO: Print e
	_, e := strconv.Atoi("wat")
	fmt.Println("e:", e)
}