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

func f(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}
	return arg + 3, nil
}

// A sentinel error is a predeclared variable that is used to
// signify a specific error condition.

var ErrOutOfTea = fmt.Errorf("no more tea available")
var ErrPower = fmt.Errorf("can't boil water")

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

	for _, i := range []int{7, 42} {
		if r, err := f(i); err != nil {
			fmt.Println("f failed:", err)
		} else {
			fmt.Println("f worked:", r)
		}
	}

	for _, i := range []int{2, 4, 5} {
		if err := makeTea(i); err != nil {
			fmt.Println("makeTea failed:", err)
		}
	}

	if errors.Is(makeTea(2), ErrOutOfTea) {
		fmt.Println("out of tea")
	}
	if errors.Is(makeTea(4), ErrPower) {
		fmt.Println("power error")
	}
}