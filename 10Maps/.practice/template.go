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

	// Set key/value pairs using typical `name[key] = val`
	// syntax.

	// TODO: set k1 to 7 and k2 to 13


	// Printing a map with e.g. `fmt.Println` will show all of
	// its key/value pairs.

	// TODO: Print map

	// Get a value for a key with `name[key]`.

	// TODO: Get and print value for key "k1"

	// If the key doesn't exist, the
	// [zero value](https://go.dev/ref/spec#The_zero_value) of the
	// value type is returned.

	// TODO: Get and print value for key "k3"

	// The builtin `len` returns the number of key/value
	// pairs when called on a map.

	// TODO: Print the length of the map

	// The builtin `delete` removes key/value pairs from
	// a map.

	// TODO: Delete key "k2" from map
	// TODO: Print the map

	// To remove *all* key/value pairs from a map, use
	// the `clear` builtin.

	// TODO: Clear the map
	// TODO: Print the map

	// The optional second return value when getting a
	// value from a map indicates if the key was present
	// in the map. This can be used to disambiguate
	// between missing keys and keys with zero values
	// like `0` or `""`. Here we didn't need the value
	// itself, so we ignored it with the _blank identifier_
	// `_`.

	// TODO: Check if key "k2" exists in map
	// TODO: Print the result

	// You can also declare and initialize a new map in
	// the same line with this syntax.

	// TODO: Create map n with key-type string and val-type int with initial values {"foo": 1, "bar": 2}
	// TODO: Print the map

	// The `maps` package contains a number of useful
	// utility functions for maps.

	// TODO: Create map n2 with key-type string and val-type int with initial values {"foo": 1, "bar": 2}
	// TODO: Use maps.Equal to compare n and n2
	// TODO: Print the result

}