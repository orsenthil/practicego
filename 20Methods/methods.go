// Go supports _methods_ defined on struct types.

package main

import "fmt"

// TODO: Define struct rect with width, height float64

type rect struct {
	width float64
	height float64
}
// Here we define an `area` method which has a _receiver type_ of `*rect`.

// TODO: Create method area() float64 on rect that returns width * height
func (r *rect) area() float64 {
	return r.width * r.height
}
// Methods can be defined for either pointer or value receiver types.
// Here's an example of a value receiver.

// TODO: Create method perim() float64 on rect that returns 2*width + 2*height
func (r *rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	// Here we call the 2 methods defined for our struct.

	// TODO: Call r.area() and print the result
	fmt.Println(r.area())
	// TODO: Call r.perim() and print the result
	fmt.Println(r.perim())

	// Go automatically handles conversion between values
	// and pointers for method calls. You may want to use
	// a pointer receiver type to avoid copying on method
	// calls or to allow the method to mutate the
	// receiving struct.

	// TODO: Create rp := &r
	rp := &r
	// TODO: Call rp.area() and rp.perim() and print results
	fmt.Println(rp.area())
	fmt.Println(rp.perim())
}