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

	// TODO: Declare slice s of strings
	// Print uninit: s, s == nil, len(s) == 0
	var s []string
	fmt.Println("s:", s, s == nil, len(s) == 0)

	// To create a slice with non-zero length, use
	// the builtin `make`. Here we make a slice of
	// `string`s of length `3` (initially zero-valued).
	// By default a new slice's capacity is equal to its
	// length; if we know the slice is going to grow ahead
	// of time, it's possible to pass a capacity explicitly
	// as an additional parameter to `make`.

	// TODO: Create slice s with make, length 3
	// Print emp: s, len: len(s), cap: cap(s)
	s = make([]string, 3)
	fmt.Println("s:", s, len(s), cap(s))
	// We can set and get just like with arrays.

	// TODO: Set s[0] = "a", s[1] = "b", s[2] = "c"
	// Print set: s, get: s[2]
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("s:", s, s[2])


	// `len` returns the length of the slice as expected.
	// Print len: len(s)

	// In addition to these basic operations, slices
	// support several more that make them richer than
	// arrays. One is the builtin `append`, which
	// returns a slice containing one or more new values.
	// Note that we need to accept a return value from
	// `append` as we may get a new slice value.

	// TODO: Append "d" to s.
	// TODO: Then append "e" and "f" to s and print slice

	s = append(s, "d")
	fmt.Println("s:", s)
	s = append(s, "e", "f")
	fmt.Println("s:", s)

	// Slices can also be `copy`'d. Here we create an
	// empty slice `c` of the same length as `s` and copy
	// into `c` from `s`.

	// TODO: Create slice c with make, same length as s
	// TODO: Copy s into c
	// Print cpy: c
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("c:", c)

	// Slices support a "slice" operator with the syntax
	// `slice[low:high]`. For example, this gets a slice
	// of the elements `s[2]`, `s[3]`, and `s[4]`.

	// TODO: Create slice l := s[2:5]
	// Print sl1: l
	l := s[2:5]
	fmt.Println("l:", l)
	// This slices up to (but excluding) `s[5]`.

	// TODO: Create slice l := s[:5]
	// Print sl2: l
	l = s[:5]
	fmt.Println("l:", l)

	// And this slices up from (and including) `s[2]`.

	// TODO: Create slice l := s[2:]
	// Print sl3: l
	l = s[2:]
	fmt.Println("l:", l)

	// We can declare and initialize a variable for slice
	// in a single line as well.

	// TODO: Create slice t := []string{"g", "h", "i"}
	// Print dcl: t
	t := []string{"g", "h", "i"}
	fmt.Println("t:", t)
	// The `slices` package contains a number of useful
	// utility functions for slices.

	// TODO: Create slice t2 := []string{"g", "h", "i"}
	// TODO: Use slices.Equal to compare t and t2
	// Print t == t2
	t2 := []string{"g", "h", "i"}
	fmt.Println("t == t2:", slices.Equal(t, t2))

	// Slices can be composed into multi-dimensional data
	// structures. The length of the inner slices can
	// vary, unlike with multi-dimensional arrays.

	// TODO: Create 2D slice twoD := make([][]int, 3)
	// TODO: Use loop to populate each inner slice with different lengths
	// Print 2d: twoD
	twoD := make([][]int, 3)
	for i := range 3 {
		twoD[i] = make([]int, i+1)
		for j := range i+1 {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("twoD:", twoD)
}