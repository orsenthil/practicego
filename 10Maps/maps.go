// _Maps_ are Go's built-in [associative data type](https://en.wikipedia.org/wiki/Associative_array)
// (sometimes called _hashes_ or _dicts_ in other languages).

package main

import (
	"fmt"
	"maps"
)

func main() {

	// To create an empty map, use the builtin `make`:
	// `make(map[key-type]val-type)`.

	// TODO: Create map m  with key-type string and val-type int
	m := make(map[string]int)

	// Set key/value pairs using typical `name[key] = val`
	// syntax.

	// TODO: set k1 to 7 and k2 to 13
	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("m:", m)

	// Printing a map with e.g. `fmt.Println` will show all of
	// its key/value pairs.

	// TODO: Print map
	fmt.Println("m:", m)
	// Get a value for a key with `name[key]`.

	// TODO: Get and print value for key "k1"
	fmt.Println("m[k1]:", m["k1"])
	// If the key doesn't exist, the
	// [zero value](https://go.dev/ref/spec#The_zero_value) of the
	// value type is returned.

	// TODO: Get and print value for key "k3"
	fmt.Println("m[k3]:", m["k3"])
	// The builtin `len` returns the number of key/value
	// pairs when called on a map.

	// TODO: Print the length of the map
	fmt.Println("length of m:", len(m))
	// The builtin `delete` removes key/value pairs from
	// a map.

	// TODO: Delete key "k2" from map
	// TODO: Print the map
	delete(m, "k2")
	fmt.Println("m:", m)
	// To remove *all* key/value pairs from a map, use
	// the `clear` builtin.

	// TODO: Clear the map
	// TODO: Print the map
	clear(m)
	fmt.Println("m:", m)
	// The optional second return value when getting a
	// value from a map indicates if the key was present
	// in the map. This can be used to disambiguate
	// between missing keys and keys with zero values
	// like `0` or `""`. Here we didn't need the value
	// itself, so we ignored it with the _blank identifier_
	// `_`.

	// TODO: Check if key "k2" exists in map
	// TODO: Print the result
	fmt.Println("m[k2]:", m["k2"])
	// You can also declare and initialize a new map in
	// the same line with this syntax.

	// TODO: Create map n with key-type string and val-type int with initial values {"foo": 1, "bar": 2}
	// TODO: Print the map
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("n:", n)
	// The `maps` package contains a number of useful
	// utility functions for maps.

	// TODO: Create map n2 with key-type string and val-type int with initial values {"foo": 1, "bar": 2}
	// TODO: Use maps.Equal to compare n and n2
	// TODO: Print the result
	n2 := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("n == n2:", maps.Equal(n, n2))
}