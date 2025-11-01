// Go offers excellent support for string formatting in
// the `printf` tradition. Here are some examples of
// common string formatting tasks.

package main

import (
	"fmt"
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

	// TODO: Create p := point{1, 2}
	// TODO: Print p using %v
	p := point{1, 2}
	fmt.Printf("p: %v\n", p)

	// If the value is a struct, the `%+v` variant will
	// include the struct's field names.

	// TODO: Print p using %+v
	fmt.Printf("p: %+v\n", p)
	// The `%#v` variant prints a Go syntax representation
	// of the value, i.e. the source code snippet that
	// would produce that value.

	// TODO: Print p using %#v
	fmt.Printf("p: %#v\n", p)
	// To print the type of a value, use `%T`.

	// TODO: Print p using %T
	fmt.Printf("p: %T\n", p)
	// Formatting booleans is straight-forward.

	// TODO: Print true using %t
	fmt.Printf("true: %t\n", true)
	// There are many options for formatting integers.
	// Use `%d` for standard, base-10 formatting.

	// TODO: Print 123 using %d
	fmt.Printf("123: %d\n", 123)
	// This prints a binary representation.
	// TODO: Print 14 using %b

	// This prints the character corresponding to the
	// given integer.

	// TODO: Print 33 using %c
	fmt.Printf("33: %c\n", 33)
	// `%x` provides hex encoding.
	
	// TODO: Print 456 using %x
	fmt.Printf("456: %x\n", 456)
	// There are also several formatting options for
	// floats. For basic decimal formatting use `%f`.

	// TODO: Print 78.9 using %f
	fmt.Printf("78.9: %f\n", 78.9)
	// `%e` and `%E` format the float in (slightly
	// different versions of) scientific notation.

	// TODO: Print 123400000.0 using %e
	// TODO: Print 123400000.0 using %E
	fmt.Printf("123400000.0: %e\n", 123400000.0)
	fmt.Printf("123400000.0: %E\n", 123400000.0)

	// For basic string printing use `%s`.

	// TODO: Print ""string"" using %s
	fmt.Printf("\"string\": %s\n", "string")
	// To double-quote strings as in Go source, use `%q`.

	// TODO: Print ""string"" using %q
	fmt.Printf("\"string\": %q\n", "string")
	// As with integers seen earlier, `%x` renders
	// the string in base-16, with two output characters
	// per byte of input.

	// TODO: Print "hex this" using %x
	fmt.Printf("\"hex this\": %x\n", "hex this")
	// To print a representation of a pointer, use `%p`.

	// TODO: Print pointer of p using %p
	fmt.Printf("pointer of p: %p\n", &p)
	// When formatting numbers you will often want to
	// control the width and precision of the resulting
	// figure. To specify the width of an integer, use a
	// number after the `%` in the verb. By default the
	// result will be right-justified and padded with
	// spaces.

	// TODO: Print "12" and "345" using %6d
	fmt.Printf("\"12\": %6d\n", 12)
	fmt.Printf("\"345\": %6d\n", 345)
	// You can also specify the width of printed floats,
	// though usually you'll also want to restrict the
	// decimal precision at the same time with the
	// width.precision syntax.

	// TODO: Print "1.2" and "3.45" using %6.2f
	fmt.Printf("\"1.2\": %6.2f\n", 1.2)
	fmt.Printf("\"3.45\": %6.2f\n", 3.45)
	// To left-justify, use the `-` flag.

	// TODO: Print "1.2" and "3.45" using %-6.2f
	fmt.Printf("\"1.2\": %-6.2f\n", 1.2)
	fmt.Printf("\"3.45\": %-6.2f\n", 3.45)
	// You may also want to control width when formatting
	// strings, especially to ensure that they align in
	// table-like output. For basic right-justified width.

	// TODO: Print "foo" and "b" using %6s
	fmt.Printf("\"foo\": %6s\n", "foo")
	fmt.Printf("\"b\": %6s\n", "b")
	// To left-justify use the `-` flag as with numbers.

	// TODO: Print "foo" and "b" using %-6s
	fmt.Printf("\"foo\": %-6s\n", "foo")
	fmt.Printf("\"b\": %-6s\n", "b")
	// So far we've seen `Printf`, which prints the
	// formatted string to `os.Stdout`. `Sprintf` formats
	// and returns a string without printing it anywhere.

	// TODO: Print "sprintf: a string" using %Sprintf
	fmt.Printf("sprintf: a string: %s\n", "a string")
	// You can format+print to `io.Writers` other than
	// `os.Stdout` using `Fprintf`.

	// TODO: Print "io: an error" using %Fprintf
	fmt.Printf("io: an error: %s\n", "an error")
}