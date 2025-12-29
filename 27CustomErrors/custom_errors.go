// It's possible to define custom error types by
// implementing the `Error()` method on them. Here's a
// variant on the example above that uses a custom type
// to explicitly represent an argument error.

package main

import (
	"errors"
	"fmt"
)

// A custom error type usually has the suffix "Error".

// TODO: Define struct argError with arg int and message string fields


// Adding this `Error` method makes `argError` implement
// the `error` interface.

// TODO: Define method Error() string on argError that returns fmt.Sprintf("%d - %s", e.arg, e.message)


// TODO: Define function f(arg int) (int, error) that returns -1, &argError{arg, "can't work with it"} if arg == 42,
// otherwise returns arg + 3, nil


func main() {

	// `errors.As` is a more advanced version of `errors.Is`.
	// It checks that a given error (or any error in its chain)
	// matches a specific error type and converts to a value
	// of that type, returning `true`. If there's no match, it
	// returns `false`.
	_, err := f(42)
	var ae *argError

	// TODO: Use errors.As to check if err is an argError
	// TODO: Print arg and message

}