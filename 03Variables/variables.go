// In Go, _variables_ are explicitly declared and used by
// the compiler to e.g. check type-correctness of function
// calls.

package main

import "fmt"

func main() {

	// `var` declares 1 or more variables.

	// TODO: Declare variable a with initial value "initial" and print it
	a := "initial"
	fmt.Println(a)
	// You can declare multiple variables at once.

	// TODO: Declare variables b and c as int with values 1 and 2 and print them
	b, c := 1, 2
	fmt.Println(b, c)
	// Go will infer the type of initialized variables.

	// TODO: Declare variable d with value true and print it
	d := true
	fmt.Println(d)
	// Variables declared without a corresponding
	// initialization are _zero-valued_. For example, the
	// zero value for an `int` is `0`.

	// TODO: Declare variable e as int without initialization and print it
	var e int
	fmt.Println(e)
	// The `:=` syntax is shorthand for declaring and
	// initializing a variable, e.g. for
	// `var f string = "apple"` in this case.
	// This syntax is only available inside functions.

	// TODO: Declare and initialize f with value "apple" using := syntax and print it
	f := "apple"
	fmt.Println(f)
}