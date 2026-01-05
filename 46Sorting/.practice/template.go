// Go's `slices` package implements sorting for builtins
// and user-defined types. We'll look at sorting for
// builtins first.

package main

import (
	"fmt"
	"slices"
)

func main() {

	// Sorting functions are generic, and work for any
	// _ordered_ built-in type. For a list of ordered
	// types, see [cmp.Ordered](https://pkg.go.dev/cmp#Ordered).

	// TODO: Create slice strs of strings with values "c", "a", "b"

	// TODO: Sort strs using slices.Sort

	// TODO: Print Strings: strs


	// An example of sorting `int`s.

	// TODO: Create slice ints of ints with values 7, 2, 4

	// TODO: Sort ints using slices.Sort

	// TODO: Print Ints: ints


	// We can also use the `slices` package to check if
	// a slice is already in sorted order.

	// TODO: Check if ints is sorted using slices.IsSorted

	// TODO: Print Sorted: s

}