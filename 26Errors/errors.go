// In Go it's idiomatic to communicate errors via an
// explicit, separate return value. This contrasts with
// the exceptions used in languages like Java, Python and
// Ruby and the overloaded single result / error value
// sometimes used in C. Go's approach makes it easy to
// see which functions return errors and to handle them
// using the same language constructs employed for other,
// non-error tasks.
//
// See the documentation of the [errors package](https://pkg.go.dev/errors)
// and [this blog post](https://go.dev/blog/go1.13-errors) for additional
// details.

package main

import (
	"errors"
	"fmt"
)

// By convention, errors are the last return value and
// have type `error`, a built-in interface.

// TODO: Create function f(arg int) (int, error) that returns -1, errors.New("can't work with 42") if arg == 42,
// otherwise returns arg + 3, nil

func f(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}
	return arg + 3, nil
}

// A sentinel error is a predeclared variable that is used to
// signify a specific error condition.

// TODO: Define var ErrOutOfTea = fmt.Errorf("no more tea available")
// TODO: Define var ErrPower = fmt.Errorf("can't boil water")

var ErrOutOfTea = errors.New("no more tea available")	
var ErrPower = errors.New("can't boil water")


// TODO: Define function makeTea(arg int) error that returns ErrOutOfTea if arg == 2, ErrPower if arg == 4, nil otherwise

func makeTea(arg int) error {
	if arg == 2 {
		return ErrOutOfTea
	}
	if arg == 4 {
		return ErrPower
	}
	return nil
}

// We can wrap errors with higher-level errors to add
// context. The simplest way to do this is with the
// `%w` verb in `fmt.Errorf`. Wrapped errors
// create a logical chain (A wraps B, which wraps C, etc.)
// that can be queried with functions like `errors.Is`
// and `errors.As`.

func main() {

	// TODO: Use range to iterate over []int{7, 42} and check if f(i) is nil
	// TODO: Print result
	// It's idiomatic to use an inline error check in the `if`
	// line.

	for _, i := range []int{7, 42} {
		if r, err := f(i); err != nil {
			fmt.Println("f(%d) failed:", i, err)
		} else {
			fmt.Println("f(%d) = %d", i, r)
		}
	}


	// TODO: Use range to iterate over 5 and check if makeTea(i) is nil
	// Check if err is ErrOutOfTea or ErrPower, print the error message

	// `errors.Is` checks that a given error (or any error in its chain)
	// matches a specific error value. This is especially useful with wrapped or
	// nested errors, allowing you to identify specific error types or sentinel
	// errors in a chain of errors.

	for i := range 5 {
		if err := makeTea(i); err != nil {
			if errors.Is(err, ErrOutOfTea) {
				fmt.Println("out of tea")
			} else if errors.Is(err, ErrPower) {
				fmt.Println("can't boil water")
			}
		}
		fmt.Println("tea made")
	}

}