// _Functions_ are central in Go. We'll learn about
// functions with a few different examples.

package main

import "fmt"

// Here's a function that takes two `int`s and returns
// their sum as an `int`.

// Go requires explicit returns, i.e. it won't
// automatically return the value of the last
// expression.


// TODO: Create function plus(a, b int) int that returns a + b

func plus(a, b int) int {
	return a + b
}


// When you have multiple consecutive parameters of
// the same type, you may omit the type name for the
// like-typed parameters up to the final parameter that
// declares the type.

// TODO: Create function plusPlus(a, b, c int) int that returns a + b + c

func plusPlus(a, b, c int) int {
	return a + b + c
}


func main() {

	// Call a function just as you'd expect, with
	// `name(args)`.

	// TODO: Call plus(1, 2) and store in res and print result

	res := plus(1, 2)
	fmt.Println("res:", res)	
	// TODO: Call plusPlus(1, 2, 3) and store in res and print result
	res = plusPlus(1, 2, 3)
	fmt.Println("res:", res)
}