// Go offers excellent support for string formatting in
// the `printf` tradition. Here are some examples of
// common string formatting tasks.

package main

import (
	"fmt"
	"os"
)

// TODO: Define struct point with x (int) and y (int) fields

func main() {

	// Go offers several printing "verbs" designed to
	// format general Go values. For example, this prints
	// an instance of our `point` struct.

	// TODO: Create p := point{1, 2}
	// TODO: Print p using %v

	// If the value is a struct, the `%+v` variant will
	// include the struct's field names.

	// TODO: Print p using %+v

	// The `%#v` variant prints a Go syntax representation
	// of the value, i.e. the source code snippet that
	// would produce that value.

	// TODO: Print p using %#v

	// To print the type of a value, use `%T`.

	// TODO: Print p using %T

	// Formatting booleans is straight-forward.

	// TODO: Print true using %t

	// There are many options for formatting integers.
	// Use `%d` for standard, base-10 formatting.

	// TODO: Print 123 using %d

	// This prints a binary representation.
	// TODO: Print 14 using %b

	// This prints the character corresponding to the
	// given integer.

	// TODO: Print 33 using %c

	// `%x` provides hex encoding.
	// TODO: Print 456 using %x

	// There are also several formatting options for
	// floats. For basic decimal formatting use `%f`.

	// TODO: Print 78.9 using %f

	// `%e` and `%E` format the float in (slightly
	// different versions of) scientific notation.

	// TODO: Print 123400000.0 using %e
	// TODO: Print 123400000.0 using %E

	// For basic string printing use `%s`.

	// TODO: Print ""string"" using %s

	// To double-quote strings as in Go source, use `%q`.

	// TODO: Print ""string"" using %q

	// As with integers seen earlier, `%x` renders
	// the string in base-16, with two output characters
	// per byte of input.

	// TODO: Print "hex this" using %x

	// To print a representation of a pointer, use `%p`.

	// TODO: Print pointer of p using %p

	// When formatting numbers you will often want to
	// control the width and precision of the resulting
	// figure. To specify the width of an integer, use a
	// number after the `%` in the verb. By default the
	// result will be right-justified and padded with
	// spaces.

	// TODO: Print "12" and "345" using %6d

	// You can also specify the width of printed floats,
	// though usually you'll also want to restrict the
	// decimal precision at the same time with the
	// width.precision syntax.

	// TODO: Print "1.2" and "3.45" using %6.2f

	// To left-justify, use the `-` flag.

	// TODO: Print "1.2" and "3.45" using %-6.2f

	// You may also want to control width when formatting
	// strings, especially to ensure that they align in
	// table-like output. For basic right-justified width.

	// TODO: Print "foo" and "b" using %6s

	// To left-justify use the `-` flag as with numbers.

	// TODO: Print "foo" and "b" using %-6s

	// So far we've seen `Printf`, which prints the
	// formatted string to `os.Stdout`. `Sprintf` formats
	// and returns a string without printing it anywhere.

	// TODO: Print "sprintf: a string" using %Sprintf

	// You can format+print to `io.Writers` other than
	// `os.Stdout` using `Fprintf`.

	// TODO: Print "io: an error" using %Fprintf
}