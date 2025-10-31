// In Go, an _array_ is a numbered sequence of elements of a
// specific length. In typical Go code, [slices](slices) are
// much more common; arrays are useful in some special
// scenarios.

package main

import "fmt"

func main() {

	// Here we create an array `a` that will hold exactly
	// 5 `int`s. The type of elements and length are both
	// part of the array's type. By default an array is
	// zero-valued, which for `int`s means `0`s.

	// TODO: Create an array `a` that will hold exactly 5 `int`s and print it
	a := [5]int{1, 2, 3, 4, 5}
	fmt.Println(a)

	// We can set a value at an index using the
	// `array[index] = value` syntax, and get a value with
	// `array[index]`.

	// TODO: Set a[4] to 100 and print the array
	a[4] = 100
	fmt.Println(a)
	// The builtin `len` returns the length of an array.

	// TODO: Print the length of the array
	fmt.Println(len(a))


	// Use this syntax to declare and initialize an array
	// in one line.

	// TODO: Declare and initialize an array `b` with values [1, 2, 3, 4, 5] and print it
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	// You can also have the compiler count the number of
	// elements for you with `...`

	// TODO: Intialize an array `b` with using [...] syntax with values [1, 2, 3, 4, 5] and print it
	b = [...]int{1, 2, 3, 4, 5}
	fmt.Println(b)
	// If you specify the index with `:`, the elements in
	// between will be zeroed.

	// TODO: Intialize an array `b` with using [...] syntax with values [100, 3: 400, 500] and print it
	b = [...]int{100, 3: 400, 500}
	fmt.Println(b)
	// Array types are one-dimensional, but you can
	// compose types to build multi-dimensional data
	// structures.

	// TODO: Create a two-dimensional array `twoD` of size [2][3]int and print it
	// Use nested loops (range 2, range 3) to populate twoD[i][j] = i + j
	twoD := [2][3]int{}
	for i := range 2 {
		for j := range 3 {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD)

	// You can create and initialize multi-dimensional
	// arrays at once too.

	// TODO: Create and initialize a two-dimensional array `twoD2` with values {{1, 2, 3}, {1, 2, 3}} and print it
	twoD2 := [2][3]int{{1, 2, 3}, {1, 2, 3}}
	fmt.Println(twoD2)
}