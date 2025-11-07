// _Interfaces_ are named collections of method
// signatures.

package main

import (
	"fmt"
	"math"
)

// Here's a basic interface for geometric shapes.

type geometry interface {
	area() float64
	perim() float64
}

// For our example we'll implement this interface on
// `rect` and `circle` types.

type rect struct {
	width float64
	height float64
}

type circle struct {
	radius float64
}

// To implement an interface in Go, we just need to
// implement all the methods in the interface. Here we
// implement `geometry` on `rect`s.

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

// The implementation for `circle`s.

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// If a variable has an interface type, then we can call
// methods that are in the named interface. Here's a
// generic `measure` function taking advantage of this
// to work on any `geometry`.

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{width: 3, height: 4}
	measure(r)
	c := circle{radius: 5}
	measure(c)

	// Sometimes it's useful to know the runtime type of an
	// interface value. One option is using a *type assertion*
	// as shown here; another is a [type `switch`](switch).

	// TODO: Create function describe(i interface{}) that uses type assertion
	// Check if i is a circle, if so print "Circle with radius" and the radius
	// Check if i is a rect, if so print "Rectangle" and dimensions
	// Otherwise print "Unknown type"

	describe := func(i interface{}) {
		fmt.Println(i)
		switch v := i.(type) {
		case circle:
			fmt.Println("Circle with radius", v.radius)
		case rect:
			fmt.Println("Rectangle with width", v.width, "and height", v.height)
		default:
			fmt.Println("Unknown type")
		}
	}

	describe(r)
	describe(c)
}