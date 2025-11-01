// Branching with `if` and `else` in Go is
// straight-forward.

// Note that you don't need parentheses around conditions
// in Go, but that the braces are required.
package main

import "fmt"

func main() {

	// Here's a basic example.

	// TODO: Check if 7%2 == 0, print "7 is even" or "7 is odd"
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	

	// You can have an `if` statement without an else.

	// TODO: Check if 8%4 == 0, print "8 is divisible by 4"
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	// Logical operators like `&&` and `||` are often
	// useful in conditions.

	// TODO: Check if 8%2 == 0 || 7%2 == 0, print "either 8 or 7 are even"

	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("either 8 or 7 are even")
	}

	// A statement can precede conditionals; any variables
	// declared in this statement are available in the current
	// and all subsequent branches.

	// TODO: Assign num := 9 and check if num < 0, print num "is negative"
	// otherwise if num < 10, print num "has 1 digit"
	// otherwise print num "has multiple digits"
	num := 9
	if num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

}