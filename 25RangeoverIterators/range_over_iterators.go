// Starting with version 1.23, Go has added support for
// [iterators](https://go.dev/blog/range-functions),
// which lets us range over pretty much anything!

package main

import (
	"fmt"
	"iter"
	"slices"
)

// Let's look at the `List` type from the
// [previous example](generics) again. In that example
// we had an `AllElements` method that returned a slice
// of all elements in the list. With Go iterators, we
// can do it better - as shown below.

// TODO: Define struct List[T any] with head and tail pointers to element[T]

// TODO: Define struct element[T any] with next pointer to element[T] and val field of type T

// TODO: Define method Push(v T) on List[T] that pushes a value v to the list
// if tail is nil, set head and tail to the new element, otherwise set tail.next to the new element and tail to the new element

// All returns an _iterator_, which in Go is a function
// with a [special signature](https://pkg.go.dev/iter#Seq).

// TODO: Define method All() iter.Seq[T] on List[T] that returns an iterator

// The iterator function takes another function as
// a parameter, called `yield` by convention (but
// the name can be arbitrary). It will call `yield` for
// every element we want to iterate over, and note `yield`'s
// return value for a potential early termination.

// Iteration doesn't require an underlying data structure,
// and doesn't even have to be finite! Here's a function
// returning an iterator over Fibonacci numbers: it keeps
// running as long as `yield` keeps returning `true`.

// TODO: Define function genFib() iter.Seq[int] that returns an iterator over Fibonacci numbers
// return a function that takes a yield function as a parameter
// inside, create variables a, b := 1, 1
// in the for loop, call yield with a, if yield returns false, return
// otherwise, set a, b = b, a+b

func main() {

	// TODO: Create lst := List[int]{}
	// TODO: Push 10, 13, 23 to lst

	// Since `List.All` returns an iterator, we can use it
	// in a regular `range` loop.

	// TODO: Use range to iterate over lst.All() and print each element

	// Packages like [slices](https://pkg.go.dev/slices) have
	// a number of useful functions to work with iterators.
	// For example, `Collect` takes any iterator and collects
	// all its values into a slice.

	// TODO: Use slices.Collect to collect all elements of lst.All() into a slice
	// TODO: Print all

	
	// TODO: Use range to iterate over genFib() and print each element

	// Once the loop hits `break` or an early return, the `yield` function
	// passed to the iterator will return `false`.

}