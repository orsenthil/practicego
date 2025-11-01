// _range_ iterates over elements in a variety of
// built-in data structures. Let's see how to
// use `range` with some of the data structures
// we've already learned.

package main

import "fmt"

func main() {

	// Here we use `range` to sum the numbers in a slice.
	// Arrays work like this too.

	// TODO: Create slice nums := []int{2, 3, 4}
	nums := []int{2, 3, 4}

	// TODO: Use range to sum all numbers in nums
	// TODO: Print sum
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)
	// `range` on arrays and slices provides both the
	// index and value for each entry. Above we didn't
	// need the index, so we ignored it with the
	// blank identifier `_`. Sometimes we actually want
	// the indexes though.

	// TODO: Use range over nums to print index and value
	for i, num := range nums {
		fmt.Println("index:", i, "value:", num)
	}

	// `range` on map iterates over key/value pairs.

	// TODO: Create map kvs := map[string]string{"a": "apple", "b": "banana"}
	kvs := map[string]string{"a": "apple", "b": "banana"}
	// TODO: Use range to iterate over kvs to print key and value
	for k, v := range kvs {
		fmt.Println("key:", k, "value:", v)
	}


	// `range` can also iterate over just the keys of a map.
	// TODO: Use range to iterate over just keys of kvs
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// `range` on strings iterates over Unicode code
	// points. The first value is the starting byte index
	// of the `rune` and the second the `rune` itself.
	// See [Strings and Runes](strings-and-runes) for more
	// details.

	// TODO: Use range over string "go" to print index and rune value
	s := "go"
	for i, r := range s {
		fmt.Println("index:", i, "rune:", r)
	}
}