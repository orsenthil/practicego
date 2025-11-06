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
	fruits := []string{"peach", "banana", "kiwi"}


	// We implement a comparison function for string
	// lengths. `cmp.Compare` is helpful for this.

	// TODO: Create lenCmp function that returns the comparison of the lengths of a and b
	lenCmp := func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	}
	// Now we can call `slices.SortFunc` with this custom
	// comparison function to sort `fruits` by name length.

	// TODO: Sort fruits using slices.SortFunc with lenCmp
	slices.SortFunc(fruits, lenCmp)
	// TODO: Print Fruits: fruits
	fmt.Println("Fruits:", fruits)

	// We can use the same technique to sort a slice of
	// values that aren't built-in types.

	// TODO: Define struct Person with name (string) and age (int) fields
	type Person struct {
		name string
		age int
	}
	people := []Person{
		{name: "Jax", age: 37},
		{name: "TJ", age: 25},
		{name: "Alex", age: 72},
	}
	ageCmp := func(a, b Person) int {
		return cmp.Compare(a.age, b.age)
	}
	slices.SortFunc(people, ageCmp)
	fmt.Println("People:", people)

}