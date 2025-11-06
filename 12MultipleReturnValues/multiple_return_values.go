// Go has built-in support for _multiple return values_.
// This feature is used often in idiomatic Go, for example
// to return both result and error values from a function.

package main

import "fmt"

func vals() (int, int) {
	return 3, 7
}

func main() {

	// Here we use the 2 different return values from the
	// call with _multiple assignment_.

	a, b := vals()
	fmt.Println(a, b)
	// TODO: Print a and b

	// If you only want a subset of the returned values,
	// use the blank identifier `_`.
	_, c := vals()
	fmt.Println(c)
	// TODO: Print c
}
