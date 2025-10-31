// Go supports _constants_ of character, string, boolean,
// and numeric values.

package main

import (
	"fmt"
	"math"
)

// `const` declares a constant value.
const s string = "constant"

func main() {
	// TODO: Print constant s from declaration
	fmt.Println(s)
	// A `const` statement can appear anywhere a `var`
	// statement can.

	// TODO: Declare constant n with value 500000000
	const n = 500000000
	// Constant expressions perform arithmetic with
	// arbitrary precision.

	// TODO: Declare constant d as 3e20 / n and print it
	const d = 3e20 / n
	fmt.Println(d)
	// A numeric constant has no type until it's given
	// one, such as by an explicit conversion.

	// TODO: Print d converted to int64
	fmt.Println(int64(d))
	// A number can be given a type by using it in a
	// context that requires one, such as a variable
	// assignment or function call. For example, here
	// `math.Sin` expects a `float64`.
	// TODO: Print the result of math.Sin(n)
	fmt.Println(math.Sin(n))
}