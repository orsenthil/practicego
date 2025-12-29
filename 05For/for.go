// `for` is Go's only looping construct. Here are
// some basic types of `for` loops.

package main

import "fmt"

func main() {

	// The most basic type, with a single condition.

	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	// A classic initial/condition/after `for` loop.

	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	// Another way of accomplishing the basic "do this
	// N times" iteration is `range` over an integer.

	for i := range 3 {
		fmt.Println("range", i)
	}

	// `for` without a condition will loop repeatedly
	// until you `break` out of the loop or `return` from
	// the enclosing function.

	for {
		fmt.Println("loop")
		break
	}

	// You can also `continue` to the next iteration of
	// the loop.

	for i := range 6 {
		if i % 2 == 0 {
			continue
		}
		fmt.Println(i)
	}
}