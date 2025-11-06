// _Slices_ are an important data type in Go, giving
// a more powerful interface to sequences than arrays.

package main

import (
	"fmt"
	"slices"
)

func main() {

	// Unlike arrays, slices are typed only by the
	// elements they contain (not the number of elements).
	// An uninitialized slice equals to nil and has
	// length 0.

	s := []string{}
	fmt.Println(s, s == nil, len(s) == 0)

	// To create a slice with non-zero length, use
	// the builtin `make`. Here we make a slice of
	// `string`s of length `3` (initially zero-valued).
	// By default a new slice's capacity is equal to its
	// length; if we know the slice is going to grow ahead
	// of time, it's possible to pass a capacity explicitly
	// as an additional parameter to `make`.

	s = make([]string, 3)
	fmt.Println(s, len(s), cap(s))
	// Print emp: s, len: len(s), cap: cap(s)

	// We can set and get just like with arrays.

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println(s)


	// `len` returns the length of the slice as expected.
	fmt.Println(len(s))

	// In addition to these basic operations, slices
	// support several more that make them richer than
	// arrays. One is the builtin `append`, which
	// returns a slice containing one or more new values.
	// Note that we need to accept a return value from
	// `append` as we may get a new slice value.

	s = append(s, "d")
	fmt.Println(s)
	s = append(s, "e", "f")
	fmt.Println(s)


	// Slices can also be `copy`'d. Here we create an
	// empty slice `c` of the same length as `s` and copy
	// into `c` from `s`.

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println(c)


	// Slices support a "slice" operator with the syntax
	// `slice[low:high]`. For example, this gets a slice
	// of the elements `s[2]`, `s[3]`, and `s[4]`.

	l := s[2:5]
	fmt.Println(l)

	// This slices up to (but excluding) `s[5]`.

	l = s[:5]
	fmt.Println(l)
	

	// And this slices up from (and including) `s[2]`.

	l = s[2:]
	fmt.Println(l)
	

	// We can declare and initialize a variable for slice
	// in a single line as well.

	t := []string{"g", "h", "i"}
	fmt.Println(t)

	// The `slices` package contains a number of useful
	// utility functions for slices.

	t2 := []string{"g", "h", "i"}
	fmt.Println(slices.Equal(t, t2))


	// Slices can be composed into multi-dimensional data
	// structures. The length of the inner slices can
	// vary, unlike with multi-dimensional arrays.

	// TODO: Create 2D slice twoD := make([][]int, 3)
	// TODO: Use loop to populate each inner slice with different lengths
	twoD := make([][]int, 3)
	for i := range 3 {
		twoD[i] = make([]int, i + 1)
		for j := range i + 1 {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD)
}