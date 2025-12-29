// Sometimes we'll want to sort a collection by something
// other than its natural order. For example, suppose we
// wanted to sort strings by their length instead of
// alphabetically. Here's an example of custom sorts
// in Go.

package main

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {

	// TODO: Create slice fruits of strings with values "peach", "banana", "kiwi"


	// We implement a comparison function for string
	// lengths. `cmp.Compare` is helpful for this.

	// TODO: Create lenCmp function that returns the comparison of the lengths of a and b

	// Now we can call `slices.SortFunc` with this custom
	// comparison function to sort `fruits` by name length.

	// TODO: Sort fruits using slices.SortFunc with lenCmp

	// TODO: Print Fruits: fruits


	// We can use the same technique to sort a slice of
	// values that aren't built-in types.

	// TODO: Define struct Person with name (string) and age (int) fields

	// TODO: Create slice people of Person with values Person{name: "Jax", age: 37}, Person{name: "TJ", age: 25}, Person{name: "Alex", age: 72}


	// Sort `people` by age using `slices.SortFunc`.
	//
	// Note: if the `Person` struct is large,
	// you may want the slice to contain `*Person` instead
	// and adjust the sorting function accordingly. If in
	// doubt, [benchmark](testing-and-benchmarking)!


	// TODO: Sort people using slices.SortFunc with func(a, b Person) int that returns the comparison of the ages of a and b

	// TODO: Print People: people

}