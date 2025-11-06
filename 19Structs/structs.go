// Go's _structs_ are typed collections of fields.
// They're useful for grouping data together to form
// records.

package main

import "fmt"

// This `person` struct type has `name` and `age` fields.

type person struct {
	name string
	age int
}

// `newPerson` constructs a new person struct with the given name.

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {
	p := person{"Bob", 20}
	fmt.Println(p)

	// An `&` prefix yields a pointer to the struct.

	fmt.Println(&person{name: "Ann", age: 40})

	// It's idiomatic to encapsulate new struct creation in constructor functions

	fmt.Println(newPerson("Jon"))

	// Access struct fields with a dot.

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	// You can also use dots with struct pointers - the
	// pointers are automatically dereferenced.

	sp := &s
	fmt.Println(sp.age)

	// Structs are mutable.

	sp.age = 51
	fmt.Println(sp.age)

	// If a struct type is only used for a single value, we don't
	// have to give it a name. The value can have an anonymous
	// struct type. This technique is commonly used for
	// [table-driven tests](testing-and-benchmarking).

	dog := struct {
		name string
		isGood bool
	}{
		name: "Rex",
		isGood: true,
	}
	fmt.Println(dog)
}