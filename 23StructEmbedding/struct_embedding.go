// Go supports _embedding_ of structs and interfaces
// to express a more seamless _composition_ of types.
// This is not to be confused with [`//go:embed`](embed-directive) which is
// a go directive introduced in Go version 1.16+ to embed
// files and folders into the application binary.

package main

import "fmt"

// TODO: Define struct base with num int field

type base struct {
	num int
}

// TODO: Create method describe() string on base that returns fmt.Sprintf("base with num=%v", b.num)

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}
// A `container` _embeds_ a `base`. An embedding looks
// like a field without a name.

// TODO: Define struct container that embeds base and has str string field

type container struct {
	base
	str string
}

func main() {

	// When creating structs with literals, we have to
	// initialize the embedding explicitly; here the
	// embedded type serves as the field name.

	// TODO: Create co := container with base: base{num: 1} and str: "some name"
	co := container{base: base{num: 1}, str: "some name"}
	// We can access the base's fields directly on `co`,
	// e.g. `co.num`.

	// TODO: Print "co={num: %v, str: %v}" with co.num and co.str
	fmt.Println("co={num: %v, str: %v}", co.num, co.str)
	// Alternatively, we can spell out the full path using
	// the embedded type name.

	// TODO: Print "also num:" followed by co.base.num
	fmt.Println("also num:", co.base.num)
	// Since `container` embeds `base`, the methods of
	// `base` also become methods of a `container`. Here
	// we invoke a method that was embedded from `base`
	// directly on `co`.

	// TODO: Print "describe:" followed by co.describe()
	fmt.Println("describe:", co.describe())
	// TODO: Define interface describer with describe() string method
	type describer interface {
		describe() string
	}
	// Embedding structs with methods may be used to bestow
	// interface implementations onto other structs. Here
	// we see that a `container` now implements the
	// `describer` interface because it embeds `base`.

	// TODO: Create var d describer = co and print "describer:" followed by d.describe()
	d := co
	fmt.Println("describer:", d.describe())
}