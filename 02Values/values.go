// Go has various value types including strings,
// integers, floats, booleans, etc. 

package main

import "fmt"

func main() {

	// Strings, which can be added together with `+`.
	// Show the result of concatenating "go" and "lang"
	fmt.Println("go" + "lang")

	// Integers and floats.
	// Show the result of 1+1 and 7.0/3.0
	fmt.Println(1+1)
	fmt.Println(7.0/3.0)
	// Booleans, with boolean operators as you'd expect.
	// Show the result of true && false, true || false, and !true
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}