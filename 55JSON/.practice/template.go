// Go offers built-in support for JSON encoding and
// decoding, including to and from built-in and custom
// data types.

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// We'll use these two structs to demonstrate encoding and
// decoding of custom types below.

// TODO: Create response1 struct with Page (int) and Fruits ([]string) fields

// Only exported fields will be encoded/decoded in JSON.
// Fields must start with capital letters to be exported.

// TODO: Create response2 struct with Page (int) and Fruits ([]string) fields

func main() {

	// First we'll look at encoding basic data types to
	// JSON strings. Here are some examples for atomic
	// values.


	// TODO: Create bolB, _ := json.Marshal(true) and print it
	// TODO: Create intB, _ := json.Marshal(1) and print it
	// TODO: Create fltB, _ := json.Marshal(2.34) and print it
	// TODO: Create strB, _ := json.Marshal("gopher") and print it


	// And here are some for slices and maps, which encode
	// to JSON arrays and objects as you'd expect.

	// TODO: Create slcD := []string{"apple", "peach", "pear"}
	// TODO: Create slcB, _ := json.Marshal(slcD) and print it

	// TODO: Create mapD := map[string]int{"apple": 5, "lettuce": 7}
	// TODO: Create mapB, _ := json.Marshal(mapD) and print it

	// The JSON package can automatically encode your
	// custom data types. It will only include exported
	// fields in the encoded output and will by default
	// use those names as the JSON keys.


	// TODO: Create res1D := &response1{Page: 1, Fruits: []string{"apple", "peach", "pear"}}
	// TODO: Create res1B, _ := json.Marshal(res1D) and print it

	// You can use tags on struct field declarations
	// to customize the encoded JSON key names. Check the
	// definition of `response2` above to see an example
	// of such tags.

	// TODO: Create res2D := &response2{Page: 1, Fruits: []string{"apple", "peach", "pear"}}
	// TODO: Create res2B, _ := json.Marshal(res2D) and print it

	// Now let's look at decoding JSON data into Go
	// values. Here's an example for a generic data
	// structure.

	// TODO: Create byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	// We need to provide a variable where the JSON
	// package can put the decoded data. This
	// `map[string]interface{}` will hold a map of strings
	// to arbitrary data types.

	// TODO: Create var dat map[string]interface{}

	// Here's the actual decoding, and a check for
	// associated errors.

	// TODO: Create if err := json.Unmarshal(byt, &dat); err != nil {
	// TODO: Print err

	// In order to use the values in the decoded map,
	// we'll need to convert them to their appropriate type.
	// For example here we convert the value in `num` to
	// the expected `float64` type.

	// TODO: Create num := dat["num"].(float64)
	// TODO: Print num

	// Accessing nested data requires a series of
	// conversions.

	// TODO: Create strs := dat["strs"].([]interface{})
	// TODO: Create str1 := strs[0].(string)
	// TODO: Print str1

	// We can also decode JSON into custom data types.
	// This has the advantages of adding additional
	// type-safety to our programs and eliminating the
	// need for type assertions when accessing the decoded
	// data.

	// TODO: Create str := `{"page": 1, "fruits": ["apple", "peach"]}`
	// TODO: Create res := response2{}
	// TODO: Create json.Unmarshal([]byte(str), &res)
	// TODO: Print res
	// TODO: Print res.Fruits[0]

	// In the examples above we always used bytes and
	// strings as intermediates between the data and
	// JSON representation on standard out. We can also
	// stream JSON encodings directly to `os.Writer`s like
	// `os.Stdout` or even HTTP response bodies.

	// TODO: Create enc := json.NewEncoder(os.Stdout)
	// TODO: Create d := map[string]int{"apple": 5, "lettuce": 7}
	// TODO: Create enc.Encode(d)

	// Streaming reads from `os.Reader`s like `os.Stdin`
	// or HTTP request bodies is done with `json.Decoder`.

	// TODO: Create dec := json.NewDecoder(strings.NewReader(str))
	// TODO: Create res1 := response2{}
	// TODO: Create dec.Decode(&res1)
	// TODO: Print res1

}