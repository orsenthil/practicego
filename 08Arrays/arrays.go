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

	a := [5]int{}
	fmt.Println("uninitialized:", a)

	// We can set a value at an index using the
	// `array[index] = value` syntax, and get a value with
	// `array[index]`.

	a[4] = 100
	fmt.Println("a:", a)

	// The builtin `len` returns the length of an array.

	fmt.Println("length of a:", len(a))


	// Use this syntax to declare and initialize an array
	// in one line.

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("b:", b)

	// You can also have the compiler count the number of
	// elements for you with `...`

	c := [...]int{1, 2, 3, 4, 5}
	fmt.Println("c:", c)

	// If you specify the index with `:`, the elements in
	// between will be zeroed.

	d := [5]int{100, 3: 400, 500}
	fmt.Println("d:", d)

	// Array types are one-dimensional, but you can
	// compose types to build multi-dimensional data
	// structures.

	twoD := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println("twoD:", twoD)

	// You can create and initialize multi-dimensional
	// arrays at once too.

	var twoD2 [2][3]int
	twoD2[0][0] = 1
	twoD2[0][1] = 2
	twoD2[0][2] = 3
	twoD2[1][0] = 1
	twoD2[1][1] = 2
	twoD2[1][2] = 3
	fmt.Println("twoD2:", twoD2)

}