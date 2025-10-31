// Go supports <em><a href="https://en.wikipedia.org/wiki/Pointer_(computer_programming)">pointers</a></em>,
// allowing you to pass references to values and records
// within your program.

package main

import "fmt"

// We'll show how pointers work in contrast to values with
// 2 functions: `zeroval` and `zeroptr`. `zeroval` has an
// `int` parameter, so arguments will be passed to it by
// value. `zeroval` will get a copy of `ival` distinct
// from the one in the calling function.

// TODO: Create function zeroval(ival int) that sets ival = 0

func zeroval(ival int) {
	ival = 0
}

// `zeroptr` in contrast has an `*int` parameter, meaning
// that it takes an `int` pointer. The `*iptr` code in the
// function body then _dereferences_ the pointer from its
// memory address to the current value at that address.
// Assigning a value to a dereferenced pointer changes the
// value at the referenced address.

// TODO: Create function zeroptr(iptr *int) that sets *iptr = 0

func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {

	// TODO: Create variable i := 1
	i := 1
	// TODO: Call zeroval(i) and print result
	zeroval(i)
	fmt.Println(i)
	// The `&i` syntax gives the memory address of `i`,
	// i.e. a pointer to `i`.

	// TODO: Call zeroptr(&i) and print result
	zeroptr(&i)
	fmt.Println(i)
	// Pointers can be printed too
	// TODO: Print pointer of i using &i
	fmt.Println(&i)
}