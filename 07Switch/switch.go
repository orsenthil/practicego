// _Switch statements_ express conditionals across many
// branches.

package main

import (
	"fmt"
	"time"
)

func main() {

	// Here's a basic `switch`.

	i := 2
	fmt.Print("Write ", i, " as ")

	// You can use commas to separate multiple expressions
	// in the same `case` statement. We use the optional
	// `default` case in this example as well.

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}


	// `switch` without an expression is an alternate way
	// to express if/else logic. Here we also show how the
	// `case` expressions can be non-constants.

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}


	// A type `switch` compares types instead of values.  You
	// can use this to discover the type of an interface
	// value.  In this example, the variable `i` will have the
	// type corresponding to its clause.

	WhatAmI := func(i interface{}) {
		switch i.(type) {
		case int:
			fmt.Println("I'm an int")
		case string:
			fmt.Println("I'm a string")
		default:
			fmt.Println("I'm something else")
		}
	}
	WhatAmI(1)
	WhatAmI("hello")
	WhatAmI(true)


}