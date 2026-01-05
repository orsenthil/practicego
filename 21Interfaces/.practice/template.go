// _Interfaces_ are named collections of method
// signatures.

package main

import (
	"fmt"
	"math"
)

// Here's a basic interface for geometric shapes.

// TODO: Define interface geometry with area() float64 and perim() float64 methods

// For our example we'll implement this interface on
// `rect` and `circle` types.

// TODO: Define struct rect with width, height float64
// TODO: Define struct circle with radius float64

// To implement an interface in Go, we just need to
// implement all the methods in the interface. Here we
// implement `geometry` on `rect`s.

// TODO: Implement area() method on rect that returns width * height
// TODO: Implement perim() method on rect that returns 2*width + 2*height

// The implementation for `circle`s.

// TODO: Implement area() method on circle that returns math.Pi * radius * radius
// TODO: Implement perim() method on circle that returns 2 * math.Pi * radius

// If a variable has an interface type, then we can call
// methods that are in the named interface. Here's a
// generic `measure` function taking advantage of this
// to work on any `geometry`.

// TODO: Create function measure(g geometry) that prints g, g.area(), and g.perim()

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	// The `circle` and `rect` struct types both
	// implement the `geometry` interface so we can use
	// instances of these structs as arguments to `measure`.

	// TODO: Call measure(r) and measure(c)

	// Sometimes it's useful to know the runtime type of an
	// interface value. One option is using a *type assertion*
	// as shown here; another is a [type `switch`](switch).

	// TODO: Create function describe(i interface{}) that uses type assertion
	// Check if i is a circle, if so print "Circle with radius" and the radius
	// Check if i is a rect, if so print "Rectangle" and dimensions
	// Otherwise print "Unknown type"

	// TODO: Call describe(r) and describe(c)
}