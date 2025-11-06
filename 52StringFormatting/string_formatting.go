// Go offers excellent support for string formatting in
// the `printf` tradition. Here are some examples of
// common string formatting tasks.

package main

import (
	"fmt"
	"os"
)

// TODO: Define struct point with x (int) and y (int) fields
type point struct {
	x int
	y int
}

func main() {

	// Go offers several printing "verbs" designed to
	// format general Go values. For example, this prints
	// an instance of our `point` struct.

	// Create p := point{1, 2}
	p := point{1, 2}
	// Print p using %v
	fmt.Printf("p using %v: %v\n", p)

	// If the value is a struct, the `%+v` variant will
	// include the struct's field names.

	// Print p using %+v
	fmt.Printf("p using %+v: %+v\n", p)

	// The `%#v` variant prints a Go syntax representation
	// of the value, i.e. the source code snippet that
	// would produce that value.

	// Print p using %#v
	fmt.Printf("p using %#v: %#v\n", p)

	// To print the type of a value, use `%T`.

	// Print p using %T
	fmt.Printf("p using %T: %T\n", p)

	// Formatting booleans is straight-forward.

	// Print true using %t
	fmt.Printf("true using %t: %t\n", true)

	// There are many options for formatting integers.
	// Use `%d` for standard, base-10 formatting.

	// Print 123 using %d
	fmt.Printf("123 using %d: %d\n", 123)

	// This prints a binary representation.
	// Print 14 using %b
	fmt.Printf("14 using %b: %b\n", 14)

	// This prints the character corresponding to the
	// given integer.

	// Print 33 using %c
	fmt.Printf("33 using %c: %c\n", 33)

	// `%x` provides hex encoding.
	// Print 456 using %x
	fmt.Printf("456 using %x: %x\n", 456)

	// There are also several formatting options for
	// floats. For basic decimal formatting use `%f`.

	// Print 78.9 using %f
	fmt.Printf("78.9 using %f: %f\n", 78.9)

	// `%e` and `%E` format the float in (slightly
	// different versions of) scientific notation.

	// Print 123400000.0 using %e
	fmt.Printf("123400000.0 using %e: %e\n", 123400000.0)
	// Print 123400000.0 using %E
	fmt.Printf("123400000.0 using %E: %E\n", 123400000.0)

	// For basic string printing use `%s`.

	// Print ""string"" using %s
	fmt.Printf("\"string\" using %s: %s\n", "string")

	// To double-quote strings as in Go source, use `%q`.

	// Print ""string"" using %q
	fmt.Printf("\"string\" using %q: %q\n", "string")

	// As with integers seen earlier, `%x` renders
	// the string in base-16, with two output characters
	// per byte of input.

	// Print "hex this" using %x
	fmt.Printf("\"hex this\" using %x: %x\n", "hex this")

	// To print a representation of a pointer, use `%p`.

	// Print pointer of p using %p
	fmt.Printf("pointer of p using %p: %p\n", &p)

	// When formatting numbers you will often want to
	// control the width and precision of the resulting
	// figure. To specify the width of an integer, use a
	// number after the `%` in the verb. By default the
	// result will be right-justified and padded with
	// spaces.

	// Print "12" and "345" using %6d
	fmt.Printf("\"12\" and \"345\" using %6d: %6d\n", 12, 345)

	// You can also specify the width of printed floats,
	// though usually you'll also want to restrict the
	// decimal precision at the same time with the
	// width.precision syntax.

	// Print "1.2" and "3.45" using %6.2f
	fmt.Printf("\"1.2\" and \"3.45\" using %6.2f: %6.2f\n", 1.2, 3.45)

	// To left-justify, use the `-` flag.

	// Print "1.2" and "3.45" using %-6.2f
	fmt.Printf("\"1.2\" and \"3.45\" using %-6.2f: %-6.2f\n", 1.2, 3.45)

	// You may also want to control width when formatting
	// strings, especially to ensure that they align in
	// table-like output. For basic right-justified width.

	// Print "foo" and "b" using %6s
	fmt.Printf("\"foo\" and \"b\" using %6s: %6s\n", "foo", "b")

	// To left-justify use the `-` flag as with numbers.

	// Print "foo" and "b" using %-6s
	fmt.Printf("\"foo\" and \"b\" using %-6s: %-6s\n", "foo", "b")

	// So far we've seen `Printf`, which prints the
	// formatted string to `os.Stdout`. `Sprintf` formats
	// and returns a string without printing it anywhere.

	// Print "sprintf: a string" using %Sprintf
	fmt.Printf("\"sprintf: a string\" using %Sprintf: %Sprintf\n", "sprintf: a string")

	// You can format+print to `io.Writers` other than
	// `os.Stdout` using `Fprintf`.

	// Print "io: an error" using %Fprintf
	fmt.Printf("\"io: an error\" using %Fprintf: %Fprintf\n", "io: an error")
}