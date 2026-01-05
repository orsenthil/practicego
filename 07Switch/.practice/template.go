// _Switch statements_ express conditionals across many
// branches.

package main

import (
	"fmt"
	"time"
)

func main() {

	// Here's a basic `switch`.

	// TODO: Set i := 2, switch on i with cases for 1, 2, 3
	i := 2
	fmt.Print("Write ", i, " as ")

	// You can use commas to separate multiple expressions
	// in the same `case` statement. We use the optional
	// `default` case in this example as well.

	// TODO: Switch on time.Now().Weekday() with cases for Saturday/Sunday and default


	// `switch` without an expression is an alternate way
	// to express if/else logic. Here we also show how the
	// `case` expressions can be non-constants.

	// TODO: Create a switch with no expression, check time conditions. Set t := time.Now()
	// If t.Hour() < 12, print "It's before noon"
	// Otherwise print "It's after noon"


	// A type `switch` compares types instead of values.  You
	// can use this to discover the type of an interface
	// value.  In this example, the variable `t` will have the
	// type corresponding to its clause.

	// TODO: Create a function that uses type switch on interface{}
	// Function WhatAmI takes an argument i of type interface{} and extracts the type of i.
	// Uses switch to print the type of i

}