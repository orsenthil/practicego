// Go's _structs_ are typed collections of fields.
// They're useful for grouping data together to form
// records.

package main

import "fmt"

// This `person` struct type has `name` and `age` fields.

// TODO: Define struct person with name (string) and age (int) fields

// `newPerson` constructs a new person struct with the given name.

// TODO: Create function newPerson(name string) *person
// Go is a garbage collected language; you can safely
// return a pointer to a local variable - it will only
// be cleaned up by the garbage collector when there
// are no active references to it.
// Create person p with the given name, set p.age = 42, return &p

func main() {

	// This syntax creates a new struct.

	// TODO: Print person{"Bob", 20}

	// You can name the fields when initializing a struct.

	// TODO: Print person{name: "Alice", age: 30}

	// Omitted fields will be zero-valued.

	// TODO: Print person{name: "Fred"}

	// An `&` prefix yields a pointer to the struct.

	// TODO: Print &person{name: "Ann", age: 40}

	// It's idiomatic to encapsulate new struct creation in constructor functions

	// TODO: Print newPerson("Jon")

	// Access struct fields with a dot.

	// TODO: Create s := person{name: "Sean", age: 50}
	// TODO: Print s.name

	// You can also use dots with struct pointers - the
	// pointers are automatically dereferenced.

	// TODO: Create sp := &s
	// TODO: Print sp.age

	// Structs are mutable.

	// TODO: Set sp.age = 51
	// TODO: Print sp.age

	// If a struct type is only used for a single value, we don't
	// have to give it a name. The value can have an anonymous
	// struct type. This technique is commonly used for
	// [table-driven tests](testing-and-benchmarking).

	// TODO: Create anonymous struct dog with name (string) and isGood (bool) fields
	// Initialize with "Rex" and true, then print it

}