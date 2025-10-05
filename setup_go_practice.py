#!/usr/bin/env python3
"""
Script to set up Go practice modules for learning Go programming concepts.
Creates directories with Go modules, each containing a practice template .go file and go.mod file.
Templates are based on Go by Example with implementation removed for practice.
"""

import os
import re
import sys
import argparse
import shutil

# Practice templates based on Go by Example
# Each template contains practice instructions and code structure without implementation
# Ordered list for progressive learning
PRACTICE_TEMPLATES = [
    {
        "key": "hello-world",
        "display_name": "Hello World",
        "template": """// Our first program will print the classic "hello world"
// message. Here's the full source code.

package main

import "fmt"

func main() {
	// Print "hello world" to the console
	
}"""
    },
    {
        "key": "values",
        "display_name": "Values",
        "template": """// Go has various value types including strings,
// integers, floats, booleans, etc. 

package main

import "fmt"

func main() {

	// Strings, which can be added together with `+`.
	// Show the result of concatenating "go" and "lang"

	// Integers and floats.
	// Show the result of 1+1 and 7.0/3.0

	// Booleans, with boolean operators as you'd expect.
	// Show the result of true && false, true || false, and !true

}"""
    },
    {
        "key": "variables",
        "display_name": "Variables",
        "template": """// In Go, _variables_ are explicitly declared and used by
// the compiler to e.g. check type-correctness of function
// calls.

package main

import "fmt"

func main() {

	// `var` declares 1 or more variables.

	// TODO: Declare variable a with initial value "initial" and print it

	// You can declare multiple variables at once.

	// TODO: Declare variables b and c as int with values 1 and 2 and print them

	// Go will infer the type of initialized variables.

	// TODO: Declare variable d with value true and print it

	// Variables declared without a corresponding
	// initialization are _zero-valued_. For example, the
	// zero value for an `int` is `0`.

	// TODO: Declare variable e as int without initialization and print it

	// The `:=` syntax is shorthand for declaring and
	// initializing a variable, e.g. for
	// `var f string = "apple"` in this case.
	// This syntax is only available inside functions.

	// TODO: Declare and initialize f with value "apple" using := syntax and print it
}"""
    },
    {
        "key": "constants",
        "display_name": "Constants",
        "template": """// Go supports _constants_ of character, string, boolean,
// and numeric values.

package main

import (
	"fmt"
	"math"
)

// `const` declares a constant value.
const s string = "constant"

func main() {
	// TODO: Print constant s from declaration

	// A `const` statement can appear anywhere a `var`
	// statement can.

	// TODO: Declare constant n with value 500000000

	// Constant expressions perform arithmetic with
	// arbitrary precision.

	// TODO: Declare constant d as 3e20 / n and print it

	// A numeric constant has no type until it's given
	// one, such as by an explicit conversion.

	// TODO: Print d converted to int64

	// A number can be given a type by using it in a
	// context that requires one, such as a variable
	// assignment or function call. For example, here
	// `math.Sin` expects a `float64`.

	// TODO: Print the result of math.Sin(n)
}"""
    },
    {
        "key": "for", 
        "display_name": "For",
        "template": """// `for` is Go's only looping construct. Here are
// some basic types of `for` loops.

package main

import "fmt"

func main() {

	// The most basic type, with a single condition.

	// TODO: Initialize i := 1 and create a for loop that runs for i <= 3 and prints i in each iteration

	// A classic initial/condition/after `for` loop.

	// TODO: Create a for loop with j := 0; j < 3; j++ and print j in each iteration

	// Another way of accomplishing the basic "do this
	// N times" iteration is `range` over an integer.

	// TODO: Use range 3 to iterate, printing "range" and the index

	// `for` without a condition will loop repeatedly
	// until you `break` out of the loop or `return` from
	// the enclosing function.

	// TODO: Create an infinite for loop that prints "loop" then breaks

	// You can also `continue` to the next iteration of
	// the loop.

	// TODO: Use range 6 to iterate through numbers 0-5, if the number is even, continue to next iteration, 
	// otherwise print the number
	
}"""
    },
    {
        "key": "if-else",
        "display_name": "If/Else",
        "template": """// Branching with `if` and `else` in Go is
// straight-forward.

// Note that you don't need parentheses around conditions
// in Go, but that the braces are required.
package main

import "fmt"

func main() {

	// Here's a basic example.

	// TODO: Check if 7%2 == 0, print "7 is even" or "7 is odd"

	

	// You can have an `if` statement without an else.

	// TODO: Check if 8%4 == 0, print "8 is divisible by 4"


	// Logical operators like `&&` and `||` are often
	// useful in conditions.

	// TODO: Check if 8%2 == 0 || 7%2 == 0, print "either 8 or 7 are even"

	// A statement can precede conditionals; any variables
	// declared in this statement are available in the current
	// and all subsequent branches.

	// TODO: Assign num := 9 and check if num < 0, print num "is negative"
	// otherwise if num < 10, print num "has 1 digit"
	// otherwise print num "has multiple digits"


}"""
    },
    {
        "key": "switch",
        "display_name": "Switch",
        "template": """// _Switch statements_ express conditionals across many
// branches.

package main

import (
	"fmt"
	"time"
)

func main() {

	// Here's a basic `switch`.

	// TODO: Set i := 2, switch on i with cases for 1, 2, 3
	i := 2
	fmt.Print("Write ", i, " as ")

	// You can use commas to separate multiple expressions
	// in the same `case` statement. We use the optional
	// `default` case in this example as well.

	// TODO: Switch on time.Now().Weekday() with cases for Saturday/Sunday and default


	// `switch` without an expression is an alternate way
	// to express if/else logic. Here we also show how the
	// `case` expressions can be non-constants.

	// TODO: Create a switch with no expression, check time conditions. Set t := time.Now()
	// If t.Hour() < 12, print "It's before noon"
	// Otherwise print "It's after noon"


	// A type `switch` compares types instead of values.  You
	// can use this to discover the type of an interface
	// value.  In this example, the variable `t` will have the
	// type corresponding to its clause.

	// TODO: Create a function that uses type switch on interface{}
	// Function WhatAmI takes an argument i of type interface{} and extracts the type of i.
	// Uses switch to print the type of i

}"""
    },
    {
        "key": "arrays",
        "display_name": "Arrays",
        "template": """// In Go, an _array_ is a numbered sequence of elements of a
// specific length. In typical Go code, [slices](slices) are
// much more common; arrays are useful in some special
// scenarios.

package main

import "fmt"

func main() {

	// Here we create an array `a` that will hold exactly
	// 5 `int`s. The type of elements and length are both
	// part of the array's type. By default an array is
	// zero-valued, which for `int`s means `0`s.

	// TODO: Create an array `a` that will hold exactly 5 `int`s and print it

	// We can set a value at an index using the
	// `array[index] = value` syntax, and get a value with
	// `array[index]`.

	// TODO: Set a[4] to 100 and print the array

	// The builtin `len` returns the length of an array.

	// TODO: Print the length of the array


	// Use this syntax to declare and initialize an array
	// in one line.

	// TODO: Declare and initialize an array `b` with values [1, 2, 3, 4, 5] and print it

	// You can also have the compiler count the number of
	// elements for you with `...`

	// TODO: Intialize an array `b` with using [...] syntax with values [1, 2, 3, 4, 5] and print it

	// If you specify the index with `:`, the elements in
	// between will be zeroed.

	// TODO: Intialize an array `b` with using [...] syntax with values [100, 3: 400, 500] and print it

	// Array types are one-dimensional, but you can
	// compose types to build multi-dimensional data
	// structures.

	// TODO: Create a two-dimensional array `twoD` of size [2][3]int and print it
	// Use nested loops (range 2, range 3) to populate twoD[i][j] = i + j


	// You can create and initialize multi-dimensional
	// arrays at once too.

	// TODO: Create and initialize a two-dimensional array `twoD2` with values {{1, 2, 3}, {1, 2, 3}} and print it

}"""
    },
    {
        "key": "slices",
        "display_name": "Slices",
        "template": """// _Slices_ are an important data type in Go, giving
// a more powerful interface to sequences than arrays.

package main

import (
	"fmt"
	"slices"
)

func main() {

	// Unlike arrays, slices are typed only by the
	// elements they contain (not the number of elements).
	// An uninitialized slice equals to nil and has
	// length 0.

	// TODO: Declare slice s of strings
	// Print uninit: s, s == nil, len(s) == 0

	// To create a slice with non-zero length, use
	// the builtin `make`. Here we make a slice of
	// `string`s of length `3` (initially zero-valued).
	// By default a new slice's capacity is equal to its
	// length; if we know the slice is going to grow ahead
	// of time, it's possible to pass a capacity explicitly
	// as an additional parameter to `make`.

	// TODO: Create slice s with make, length 3
	// Print emp: s, len: len(s), cap: cap(s)

	// We can set and get just like with arrays.

	// TODO: Set s[0] = "a", s[1] = "b", s[2] = "c"
	// Print set: s, get: s[2]
	


	// `len` returns the length of the slice as expected.
	// Print len: len(s)

	// In addition to these basic operations, slices
	// support several more that make them richer than
	// arrays. One is the builtin `append`, which
	// returns a slice containing one or more new values.
	// Note that we need to accept a return value from
	// `append` as we may get a new slice value.

	// TODO: Append "d" to s.
	// TODO: Then append "e" and "f" to s and print slice



	// Slices can also be `copy`'d. Here we create an
	// empty slice `c` of the same length as `s` and copy
	// into `c` from `s`.

	// TODO: Create slice c with make, same length as s
	// TODO: Copy s into c
	// Print cpy: c


	// Slices support a "slice" operator with the syntax
	// `slice[low:high]`. For example, this gets a slice
	// of the elements `s[2]`, `s[3]`, and `s[4]`.

	// TODO: Create slice l := s[2:5]
	// Print sl1: l

	// This slices up to (but excluding) `s[5]`.

	// TODO: Create slice l := s[:5]
	// Print sl2: l
	

	// And this slices up from (and including) `s[2]`.

	// TODO: Create slice l := s[2:]
	// Print sl3: l
	

	// We can declare and initialize a variable for slice
	// in a single line as well.

	// TODO: Create slice t := []string{"g", "h", "i"}
	// Print dcl: t

	// The `slices` package contains a number of useful
	// utility functions for slices.

	// TODO: Create slice t2 := []string{"g", "h", "i"}
	// TODO: Use slices.Equal to compare t and t2
	// Print t == t2


	// Slices can be composed into multi-dimensional data
	// structures. The length of the inner slices can
	// vary, unlike with multi-dimensional arrays.

	// TODO: Create 2D slice twoD := make([][]int, 3)
	// TODO: Use loop to populate each inner slice with different lengths
	// Print 2d: twoD
	
}"""
    },
    {
        "key": "maps",
        "display_name": "Maps",
        "template": """// _Maps_ are Go's built-in [associative data type](https://en.wikipedia.org/wiki/Associative_array)
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

}"""
    },
    {
        "key": "functions",
        "display_name": "Functions",
        "template": """// _Functions_ are central in Go. We'll learn about
// functions with a few different examples.

package main

import "fmt"

// Here's a function that takes two `int`s and returns
// their sum as an `int`.

// Go requires explicit returns, i.e. it won't
// automatically return the value of the last
// expression.


// TODO: Create function plus(a, b int) int that returns a + b


// When you have multiple consecutive parameters of
// the same type, you may omit the type name for the
// like-typed parameters up to the final parameter that
// declares the type.

// TODO: Create function plusPlus(a, b, c int) int that returns a + b + c


func main() {

	// Call a function just as you'd expect, with
	// `name(args)`.

	// TODO: Call plus(1, 2) and store in res and print result

	// TODO: Call plusPlus(1, 2, 3) and store in res and print result

}"""
    },
    {
        "key": "multiple-return-values",
        "display_name": "Multiple Return Values",
        "template": """// Go has built-in support for _multiple return values_.
// This feature is used often in idiomatic Go, for example
// to return both result and error values from a function.

package main

import "fmt"

// TODO: Create function vals() (int, int) that returns 3, 7


func main() {

	// Here we use the 2 different return values from the
	// call with _multiple assignment_.

	// TODO: Call vals() and assign to a, b
	// TODO: Print a and b

	// If you only want a subset of the returned values,
	// use the blank identifier `_`.
	// TODO: Call vals() and only use the second return value
	// TODO: Print c
}"""
    },
    {
        "key": "variadic-functions",
        "display_name": "Variadic Functions",
        "template": """// [_Variadic functions_](https://en.wikipedia.org/wiki/Variadic_function)
// can be called with any number of trailing arguments.
// For example, `fmt.Println` is a common variadic
// function.

package main

import "fmt"

// Here's a function that will take an arbitrary number
// of `int`s as arguments.

// TODO: Create function sum(nums ...int) that calculates sum of all nums
// Within the function, the type of `nums` is equivalent to `[]int`. We can call `len(nums)`,
// We can iterate over it with `range`, etc.


func main() {

	// Variadic functions can be called in the usual way
	// with individual arguments.

	// TODO: Call sum(1, 2) and sum(1, 2, 3)

	// If you already have multiple args in a slice,
	// apply them to a variadic function using
	// `func(slice...)` like this.

	// TODO: Create slice nums := []int{1, 2, 3, 4}
	// TODO: Call sum(nums...)

}"""
    },
    {
        "key": "closures",
        "display_name": "Closures",
        "template": """// Go supports [_anonymous functions_](https://en.wikipedia.org/wiki/Anonymous_function),
// which can form <a href="https://en.wikipedia.org/wiki/Closure_(computer_science)"><em>closures</em></a>.
// Anonymous functions are useful when you want to define
// a function inline without having to name it.

package main

import "fmt"

// This function `intSeq` returns another function, which
// we define anonymously in the body of `intSeq`. The
// returned function _closes over_ the variable `i` to
// form a closure.

// TODO: Create function intSeq() func() int
// Inside, create variable i := 0
// Return anonymous function that increments i and returns it


func main() {

	// We call `intSeq`, assigning the result (a function)
	// to `nextInt`. This function value captures its
	// own `i` value, which will be updated each time
	// we call `nextInt`.

	// TODO: Call intSeq() and assign to nextInt

	// See the effect of the closure by calling `nextInt`
	// a few times.

	// TODO: Call nextInt() multiple times and print results for each call

	// To confirm that the state is unique to that
	// particular function, create and test a new one.

	// TODO: Create newInts := intSeq() and call it to show separate state
}"""
    },
    {
        "key": "recursion",
        "display_name": "Recursion",
        "template": """// Go supports
// <a href="https://en.wikipedia.org/wiki/Recursion_(computer_science)"><em>recursive functions</em></a>.
// Here's a classic example.

package main

import "fmt"

// This `fact` function calls itself until it reaches the
// base case of `fact(0)`.
// TODO: Create recursive function fact(n int) int
// Base case: if n == 0 return 1
// Recursive case: return n * fact(n-1)

func main() {
	
	// TODO: Call fact(7) and print result

	// Anonymous functions can also be recursive, but this requires
	// explicitly declaring a variable with `var` to store
	// the function before it's defined.

	// TODO: Create variable fib of type func(int) int
	// Assign anonymous function that calculates fibonacci recursively

	// Since `fib` was previously declared, you can call that with the anonymous function


	// TODO: Call fib(7) and print result
}"""
    },
    {
        "key": "range",
        "display_name": "Range over Built-in Types",
        "template": """// _range_ iterates over elements in a variety of
// built-in data structures. Let's see how to
// use `range` with some of the data structures
// we've already learned.

package main

import "fmt"

func main() {

	// Here we use `range` to sum the numbers in a slice.
	// Arrays work like this too.

	// TODO: Create slice nums := []int{2, 3, 4}

	// TODO: Use range to sum all numbers in nums
	// TODO: Print sum

	// `range` on arrays and slices provides both the
	// index and value for each entry. Above we didn't
	// need the index, so we ignored it with the
	// blank identifier `_`. Sometimes we actually want
	// the indexes though.

	// TODO: Use range over nums to print index and value

	// `range` on map iterates over key/value pairs.

	// TODO: Create map kvs := map[string]string{"a": "apple", "b": "banana"}

	// TODO: Use range to iterate over kvs to print key and value


	// `range` can also iterate over just the keys of a map.
	// TODO: Use range to iterate over just keys of kvs


	// `range` on strings iterates over Unicode code
	// points. The first value is the starting byte index
	// of the `rune` and the second the `rune` itself.
	// See [Strings and Runes](strings-and-runes) for more
	// details.

	// TODO: Use range over string "go" to print index and rune value

}"""
    },
    {
        "key": "pointers",
        "display_name": "Pointers", 
        "template": """// Go supports <em><a href="https://en.wikipedia.org/wiki/Pointer_(computer_programming)">pointers</a></em>,
// allowing you to pass references to values and records
// within your program.

package main

import "fmt"

// We'll show how pointers work in contrast to values with
// 2 functions: `zeroval` and `zeroptr`. `zeroval` has an
// `int` parameter, so arguments will be passed to it by
// value. `zeroval` will get a copy of `ival` distinct
// from the one in the calling function.

// TODO: Create function zeroval(ival int) that sets ival = 0

// `zeroptr` in contrast has an `*int` parameter, meaning
// that it takes an `int` pointer. The `*iptr` code in the
// function body then _dereferences_ the pointer from its
// memory address to the current value at that address.
// Assigning a value to a dereferenced pointer changes the
// value at the referenced address.

// TODO: Create function zeroptr(iptr *int) that sets *iptr = 0


func main() {

	// TODO: Create variable i := 1

	// TODO: Call zeroval(i) and print result

	// The `&i` syntax gives the memory address of `i`,
	// i.e. a pointer to `i`.

	// TODO: Call zeroptr(&i) and print result

	// Pointers can be printed too
	// TODO: Print pointer of i using &i
}"""
    },
    {
        "key": "strings-and-runes",
        "display_name": "Strings and Runes",
        "template": """// A Go string is a read-only slice of bytes. The language
// and the standard library treat strings specially - as
// containers of text encoded in [UTF-8](https://en.wikipedia.org/wiki/UTF-8).
// In other languages, strings are made of "characters".
// In Go, the concept of a character is called a `rune` - it's
// an integer that represents a Unicode code point.
// [This Go blog post](https://go.dev/blog/strings) is a good
// introduction to the topic.

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	// `s` is a `string` assigned a literal value
	// representing the word "hello" in the Thai
	// language. Go string literals are UTF-8
	// encoded text.
	const s = "สวัสดี"

	// Since strings are equivalent to `[]byte`, this
	// will produce the length of the raw bytes stored within.

	// TODO: Print length of s

	// Indexing into a string produces the raw byte values at
	// each index. This loop generates the hex values of all
	// the bytes that constitute the code points in `s`.

	// TODO: Loop through s and print hex values

	// To count how many _runes_ are in a string, we can use
	// the `utf8` package. Note that the run-time of
	// `RuneCountInString` depends on the size of the string,
	// because it has to decode each UTF-8 rune sequentially.
	// Some Thai characters are represented by UTF-8 code points
	// that can span multiple bytes, so the result of this count
	// may be surprising.

	// TODO: Print rune count of s. Use utf8.RuneCountInString

	// A `range` loop handles strings specially and decodes
	// each `rune` along with its offset in the string.

	// TODO: Loop through s and print index and rune value

	// We can achieve the same iteration by using the
	// `utf8.DecodeRuneInString` function explicitly.


	// TODO: Loop through s and print index and rune value using utf8.DecodeRuneInString
	// Use for loop and i, w := 0, 0; i < len(s); i += w, where w is the width of the rune

	// Also demonstrate passing a `rune` value to a function, examineRune

	fmt.Println("\\nUsing DecodeRuneInString")

}

// TODO: Create function examineRune(r rune) that prints "found tee" if r is 't' and "found so sua" if r is 'ส'"""
    },
    {
        "key": "structs",
        "display_name": "Structs",
        "template": """// Go's _structs_ are typed collections of fields.
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

}"""
    },
    {
        "key": "methods",
        "display_name": "Methods",
        "template": """// Go supports _methods_ defined on struct types.

package main

import "fmt"

// TODO: Define struct rect with width, height float64

// Here we define an `area` method which has a _receiver type_ of `*rect`.

// TODO: Create method area() float64 on rect that returns width * height

// Methods can be defined for either pointer or value receiver types.
// Here's an example of a value receiver.

// TODO: Create method perim() float64 on rect that returns 2*width + 2*height

func main() {
	r := rect{width: 10, height: 5}

	// Here we call the 2 methods defined for our struct.

	// TODO: Call r.area() and print the result
	// TODO: Call r.perim() and print the result

	// Go automatically handles conversion between values
	// and pointers for method calls. You may want to use
	// a pointer receiver type to avoid copying on method
	// calls or to allow the method to mutate the
	// receiving struct.

	// TODO: Create rp := &r
	// TODO: Call rp.area() and rp.perim() and print results
}"""
    },
    {
        "key": "interfaces",
        "display_name": "Interfaces",
        "template": """// _Interfaces_ are named collections of method
// signatures.

package main

import (
	"fmt"
	"math"
)

// Here's a basic interface for geometric shapes.

// TODO: Define interface geometry with area() float64 and perim() float64 methods

// For our example we'll implement this interface on
// `rect` and `circle` types.

// TODO: Define struct rect with width, height float64
// TODO: Define struct circle with radius float64

// To implement an interface in Go, we just need to
// implement all the methods in the interface. Here we
// implement `geometry` on `rect`s.

// TODO: Implement area() method on rect that returns width * height
// TODO: Implement perim() method on rect that returns 2*width + 2*height

// The implementation for `circle`s.

// TODO: Implement area() method on circle that returns math.Pi * radius * radius
// TODO: Implement perim() method on circle that returns 2 * math.Pi * radius

// If a variable has an interface type, then we can call
// methods that are in the named interface. Here's a
// generic `measure` function taking advantage of this
// to work on any `geometry`.

// TODO: Create function measure(g geometry) that prints g, g.area(), and g.perim()

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	// The `circle` and `rect` struct types both
	// implement the `geometry` interface so we can use
	// instances of these structs as arguments to `measure`.

	// TODO: Call measure(r) and measure(c)

	// Sometimes it's useful to know the runtime type of an
	// interface value. One option is using a *type assertion*
	// as shown here; another is a [type `switch`](switch).

	// TODO: Create function describe(i interface{}) that uses type assertion
	// Check if i is a circle, if so print "Circle with radius" and the radius
	// Check if i is a rect, if so print "Rectangle" and dimensions
	// Otherwise print "Unknown type"

	// TODO: Call describe(r) and describe(c)
}"""
    },
    {
        "key": "enums",
        "display_name": "Enums",
        "template": """// _Enumerated types_ (enums) are a special case of
// [sum types](https://en.wikipedia.org/wiki/Algebraic_data_type).
// An enum is a type that has a fixed number of possible
// values, each with a distinct name. Go doesn't have an
// enum type as a distinct language feature, but enums
// are simple to implement using existing language idioms.

package main

import "fmt"

// Our enum type `ServerState` has an underlying `int` type.

// TODO: Define struct ServerState with underlying int type

// The possible values for `ServerState` are defined as
// constants. The special keyword [iota](https://go.dev/ref/spec#Iota)
// generates successive constant values automatically; in this
// case 0, 1, 2 and so on.

// TODO: Define constants for StateIdle, StateConnected, StateError, StateRetrying

// By implementing the [fmt.Stringer](https://pkg.go.dev/fmt#Stringer)
// interface, values of `ServerState` can be printed out or converted
// to strings.
//
// This can get cumbersome if there are many possible values. In such
// cases the [stringer tool](https://pkg.go.dev/golang.org/x/tools/cmd/stringer)
// can be used in conjunction with `go:generate` to automate the
// process. See [this post](https://eli.thegreenplace.net/2021/a-comprehensive-guide-to-go-generate)
// for a longer explanation.

// TODO: Define map stateName with ServerState keys and string values
// StateIdle: "idle",
// StateConnected: "connected",
// StateError: "error",
// StateRetrying: "retrying",

// TODO: Define method String() string on ServerState that returns stateName[ss]

func main() {
	
	// TODO: Create ns := transition(StateIdle) and print it

	// If we have a value of type `int`, we cannot pass it to `transition` - the
	// compiler will complain about type mismatch. This provides some degree of
	// compile-time type safety for enums.

	// TODO: Create ns2 := transition(ns) and print it

}

// transition emulates a state transition for a
// server; it takes the existing state and returns
// a new state.

// TODO: Create function transition(s ServerState) ServerState that 
// returns StateConnected if s is StateIdle, 
// StateIdle if s is StateConnected or StateRetrying, 
// StateError if s is StateError"""
    },
    {
        "key": "struct-embedding",
        "display_name": "Struct Embedding",
        "template": """// Go supports _embedding_ of structs and interfaces
// to express a more seamless _composition_ of types.
// This is not to be confused with [`//go:embed`](embed-directive) which is
// a go directive introduced in Go version 1.16+ to embed
// files and folders into the application binary.

package main

import "fmt"

// TODO: Define struct base with num int field

// TODO: Create method describe() string on base that returns fmt.Sprintf("base with num=%v", b.num)

// A `container` _embeds_ a `base`. An embedding looks
// like a field without a name.

// TODO: Define struct container that embeds base and has str string field

func main() {

	// When creating structs with literals, we have to
	// initialize the embedding explicitly; here the
	// embedded type serves as the field name.

	// TODO: Create co := container with base: base{num: 1} and str: "some name"

	// We can access the base's fields directly on `co`,
	// e.g. `co.num`.

	// TODO: Print "co={num: %v, str: %v}" with co.num and co.str

	// Alternatively, we can spell out the full path using
	// the embedded type name.

	// TODO: Print "also num:" followed by co.base.num

	// Since `container` embeds `base`, the methods of
	// `base` also become methods of a `container`. Here
	// we invoke a method that was embedded from `base`
	// directly on `co`.

	// TODO: Print "describe:" followed by co.describe()

	// TODO: Define interface describer with describe() string method

	// Embedding structs with methods may be used to bestow
	// interface implementations onto other structs. Here
	// we see that a `container` now implements the
	// `describer` interface because it embeds `base`.

	// TODO: Create var d describer = co and print "describer:" followed by d.describe()
}"""
    },
    {
        "key": "generics",
        "display_name": "Generics",
        "template": """// Starting with version 1.18, Go has added support for
// _generics_, also known as _type parameters_.

package main

import "fmt"

// As an example of a generic function, `SlicesIndex` takes
// a slice of any `comparable` type and an element of that
// type and returns the index of the first occurrence of
// v in s, or -1 if not present. The `comparable` constraint
// means that we can compare values of this type with the
// `==` and `!=` operators. For a more thorough explanation
// of this type signature, see [this blog post](https://go.dev/blog/deconstructing-type-parameters).
// Note that this function exists in the standard library
// as [slices.Index](https://pkg.go.dev/slices#Index).


// TODO: Create function SlicesIndex[S ~[]E, E comparable](s S, v E) int that returns the index of the first occurrence of v in s, or -1 if not present


// As an example of a generic type, `List` is a
// singly-linked list with values of any type.


// TODO: Define struct List[T any] with head and tail pointers to element[T]


// TODO: Define struct element[T any] with next pointer to element[T] and val field of type T


// We can define methods on generic types just like we
// do on regular types, but we have to keep the type
// parameters in place. The type is `List[T]`, not `List`.


// TODO: Define method Push(v T) on List[T] that pushes a value v to the list


// AllElements returns all the List elements as a slice.
// In the next example we'll see a more idiomatic way
// of iterating over all elements of custom types.


// TODO: Define method AllElements() []T on List[T] that returns all the List elements as a slice


func main() {

	// TODO: Create var s = []string{"foo", "bar", "zoo"}


	// When invoking generic functions, we can often rely
	// on _type inference_. Note that we don't have to
	// specify the types for `S` and `E` when
	// calling `SlicesIndex` - the compiler infers them
	// automatically.

	// TODO: Print index of zoo, zoo should be 2

	// ... though we could also specify them explicitly.

	// TODO: Get index of zoo using explicit types

	// TODO: Create lst := List[int]{}
	// TODO: Push 10, 13, 23 to lst
	// TODO: Print list: lst.AllElements()

}"""
    },
    {
        "key": "range-over-iterators",
        "display_name": "Range over Iterators",
        "template": """// Starting with version 1.23, Go has added support for
// [iterators](https://go.dev/blog/range-functions),
// which lets us range over pretty much anything!

package main

import (
	"fmt"
	"iter"
	"slices"
)

// Let's look at the `List` type from the
// [previous example](generics) again. In that example
// we had an `AllElements` method that returned a slice
// of all elements in the list. With Go iterators, we
// can do it better - as shown below.

// TODO: Define struct List[T any] with head and tail pointers to element[T]

// TODO: Define struct element[T any] with next pointer to element[T] and val field of type T

// TODO: Define method Push(v T) on List[T] that pushes a value v to the list
// if tail is nil, set head and tail to the new element, otherwise set tail.next to the new element and tail to the new element

// All returns an _iterator_, which in Go is a function
// with a [special signature](https://pkg.go.dev/iter#Seq).

// TODO: Define method All() iter.Seq[T] on List[T] that returns an iterator

// The iterator function takes another function as
// a parameter, called `yield` by convention (but
// the name can be arbitrary). It will call `yield` for
// every element we want to iterate over, and note `yield`'s
// return value for a potential early termination.

// Iteration doesn't require an underlying data structure,
// and doesn't even have to be finite! Here's a function
// returning an iterator over Fibonacci numbers: it keeps
// running as long as `yield` keeps returning `true`.

// TODO: Define function genFib() iter.Seq[int] that returns an iterator over Fibonacci numbers
// return a function that takes a yield function as a parameter
// inside, create variables a, b := 1, 1
// in the for loop, call yield with a, if yield returns false, return
// otherwise, set a, b = b, a+b

func main() {

	// TODO: Create lst := List[int]{}
	// TODO: Push 10, 13, 23 to lst

	// Since `List.All` returns an iterator, we can use it
	// in a regular `range` loop.

	// TODO: Use range to iterate over lst.All() and print each element

	// Packages like [slices](https://pkg.go.dev/slices) have
	// a number of useful functions to work with iterators.
	// For example, `Collect` takes any iterator and collects
	// all its values into a slice.

	// TODO: Use slices.Collect to collect all elements of lst.All() into a slice
	// TODO: Print all

	
	// TODO: Use range to iterate over genFib() and print each element

	// Once the loop hits `break` or an early return, the `yield` function
	// passed to the iterator will return `false`.

}"""
    },
    {
        "key": "errors",
        "display_name": "Errors",
        "template": """// In Go it's idiomatic to communicate errors via an
// explicit, separate return value. This contrasts with
// the exceptions used in languages like Java, Python and
// Ruby and the overloaded single result / error value
// sometimes used in C. Go's approach makes it easy to
// see which functions return errors and to handle them
// using the same language constructs employed for other,
// non-error tasks.
//
// See the documentation of the [errors package](https://pkg.go.dev/errors)
// and [this blog post](https://go.dev/blog/go1.13-errors) for additional
// details.

package main

import (
	"errors"
	"fmt"
)

// By convention, errors are the last return value and
// have type `error`, a built-in interface.

// TODO: Create function f(arg int) (int, error) that returns -1, errors.New("can't work with 42") if arg == 42,
// otherwise returns arg + 3, nil

// A sentinel error is a predeclared variable that is used to
// signify a specific error condition.

// TODO: Define var ErrOutOfTea = fmt.Errorf("no more tea available")
// TODO: Define var ErrPower = fmt.Errorf("can't boil water")

// TODO: Define function makeTea(arg int) error that returns ErrOutOfTea if arg == 2, ErrPower if arg == 4, nil otherwise

// We can wrap errors with higher-level errors to add
// context. The simplest way to do this is with the
// `%w` verb in `fmt.Errorf`. Wrapped errors
// create a logical chain (A wraps B, which wraps C, etc.)
// that can be queried with functions like `errors.Is`
// and `errors.As`.

func main() {

	// TODO: Use range to iterate over []int{7, 42} and check if f(i) is nil
	// TODO: Print result
	// It's idiomatic to use an inline error check in the `if`
	// line.


	// TODO: Use range to iterate over 5 and check if makeTea(i) is nil
	// Check if err is ErrOutOfTea or ErrPower, print the error message

	// `errors.Is` checks that a given error (or any error in its chain)
	// matches a specific error value. This is especially useful with wrapped or
	// nested errors, allowing you to identify specific error types or sentinel
	// errors in a chain of errors.

}"""
    },
    {
        "key": "custom-errors",
        "display_name": "Custom Errors",
        "template": """// It's possible to define custom error types by
// implementing the `Error()` method on them. Here's a
// variant on the example above that uses a custom type
// to explicitly represent an argument error.

package main

import (
	"errors"
	"fmt"
)

// A custom error type usually has the suffix "Error".

// TODO: Define struct argError with arg int and message string fields


// Adding this `Error` method makes `argError` implement
// the `error` interface.

// TODO: Define method Error() string on argError that returns fmt.Sprintf("%d - %s", e.arg, e.message)


// TODO: Define function f(arg int) (int, error) that returns -1, &argError{arg, "can't work with it"} if arg == 42,
// otherwise returns arg + 3, nil


func main() {

	// `errors.As` is a more advanced version of `errors.Is`.
	// It checks that a given error (or any error in its chain)
	// matches a specific error type and converts to a value
	// of that type, returning `true`. If there's no match, it
	// returns `false`.
	_, err := f(42)
	var ae *argError

	// TODO: Use errors.As to check if err is an argError
	// TODO: Print arg and message

}"""
    },
    {
        "key": "goroutines",
        "display_name": "Goroutines",
        "template": """// A _goroutine_ is a lightweight thread of execution.

package main

import (
	"fmt"
	"time"
)

// TODO: Define function f(from string) that prints from and i in each iteration
// inside, use i:= range 3 to iterate


func main() {

	// Suppose we have a function call `f(s)`. Here's how
	// we'd call that in the usual way, running it
	// synchronously.

	// TODO: Call f("direct")


	// To invoke this function in a goroutine, use
	// `go f(s)`. This new goroutine will execute
	// concurrently with the calling one.

	// TODO: Call go f("goroutine")


	// You can also start a goroutine for an anonymous
	// function call.

	// TODO: Call go func(msg string) {
	// fmt.Println(msg)
	// }("going")

	// Our two function calls are running asynchronously in
	// separate goroutines now. Wait for them to finish
	// (for a more robust approach, use a [WaitGroup](waitgroups)).

	// TODO: Sleep for 1 second
}"""
    },
    {
        "key": "channels",
        "display_name": "Channels",
        "template": """// _Channels_ are the pipes that connect concurrent
// goroutines. You can send values into channels from one
// goroutine and receive those values into another
// goroutine.

package main

import "fmt"

func main() {

	// Create a new channel with `make(chan val-type)`.
	// Channels are typed by the values they convey.

	// TODO: Create messages channel of strings

	// _Send_ a value into a channel using the `channel <-`
	// syntax. Here we send `"ping"`  to the `messages`
	// channel we made above, from a new goroutine.

	// TODO: Send "ping" to messages channel

	// The `<-channel` syntax _receives_ a value from the
	// channel. Here we'll receive the `"ping"` message
	// we sent above and print it out.

	// TODO: Receive msg from messages channel and print it

}"""
    },
    {
        "key": "channel-buffering",
        "display_name": "Channel Buffering",
        "template": """// By default channels are _unbuffered_, meaning that they
// will only accept sends (`chan <-`) if there is a
// corresponding receive (`<- chan`) ready to receive the
// sent value. _Buffered channels_ accept a limited
// number of  values without a corresponding receiver for
// those values.

package main

import "fmt"

func main() {

	// Here we `make` a channel of strings buffering up to
	// 2 values.

	// TODO: Create messages channel of strings buffering up to 2 values

	// Because this channel is buffered, we can send these
	// values into the channel without a corresponding
	// concurrent receive.

	// TODO: Send "buffered" and "channel" to messages channel

	// Later we can receive these two values as usual.

	// TODO: Receive "buffered" and "channel" from messages channel and print them

}"""
    },
    {
        "key": "channel-synchronization",
        "display_name": "Channel Synchronization",
        "template": """// We can use channels to synchronize execution
// across goroutines. Here's an example of using a
// blocking receive to wait for a goroutine to finish.
// When waiting for multiple goroutines to finish,
// you may prefer to use a [WaitGroup](waitgroups).

package main

import (
	"fmt"
	"time"
)

// This is the function we'll run in a goroutine. The
// `done` channel will be used to notify another
// goroutine that this function's work is done.

// TODO: Define function worker(done chan bool) that prints "working..." and sleeps for 1 second, 
// then prints "done" and sends true to the done channel.


func main() {

	// Start a worker goroutine, giving it the channel to
	// notify on.

	// TODO: Create done channel of bool with buffer size 1

	// TODO: Call go worker(done)

	// Block until we receive a notification from the
	// worker on the channel.

	// TODO: Receive from done channel
}"""
    },
    {
        "key": "channel-directions",
        "display_name": "Channel Directions",
        "template": """// When using channels as function parameters, you can
// specify if a channel is meant to only send or receive
// values. This specificity increases the type-safety of
// the program.

package main

import "fmt"

// This `ping` function only accepts a channel for sending
// values. It would be a compile-time error to try to
// receive on this channel.

// TODO: Define function ping(pings chan<- string, msg string) that sends msg to pings channel


// The `pong` function accepts one channel for receives
// (`pings`) and a second for sends (`pongs`).

// TODO: Define function pong(pings <-chan string, pongs chan<- string) that 
// receives msg from pings channel and sends it to pongs channel


func main() {

	// TODO: Create pings channel of strings with buffer size 1

	// TODO: Create pongs channel of strings with buffer size 1

	// TODO: Call ping(pings, "passed message")

	// TODO: Call pong(pings, pongs)

	// TODO: Receive from pongs channel and print it
}"""
    },
    {
        "key": "select",
        "display_name": "Select",
        "template": """// Go's _select_ lets you wait on multiple channel
// operations. Combining goroutines and channels with
// select is a powerful feature of Go.

package main

import (
	"fmt"
	"time"
)

func main() {

	// For our example we'll select across two channels.

	// TODO: Create c1 channel of strings

	// TODO: Create c2 channel of strings


	// Each channel will receive a value after some amount
	// of time, to simulate e.g. blocking RPC operations
	// executing in concurrent goroutines.

	// TODO Creat a goroutine that sends "one" to c1 after 1 second

	// TODO: Creat a goroutine that sends "two" to c2 after 2 seconds


	// We'll use `select` to await both of these values
	// simultaneously, printing each one as it arrives.

	// TODO: Use for range 2 to receive from c1 and c2

}"""
    },
    {
        "key": "timeouts",
        "display_name": "Timeouts",
        "template": """// _Timeouts_ are important for programs that connect to
// external resources or that otherwise need to bound
// execution time. Implementing timeouts in Go is easy and
// elegant thanks to channels and `select`.

package main

import (
	"fmt"
	"time"
)

func main() {

	// For our example, suppose we're executing an external
	// call that returns its result on a channel `c1`
	// after 2s. Note that the channel is buffered, so the
	// send in the goroutine is nonblocking. This is a
	// common pattern to prevent goroutine leaks in case the
	// channel is never read.

	// TODO: Create c1 channel of strings with buffer size 1

	// TODO: Creat a goroutine that sends "result 1" to c1 after 2 seconds


	// Here's the `select` implementing a timeout.
	// `res := <-c1` awaits the result and `<-time.After`
	// awaits a value to be sent after the timeout of
	// 1s. Since `select` proceeds with the first
	// receive that's ready, we'll take the timeout case
	// if the operation takes more than the allowed 1s.

	// TODO: Use select to receive from c1 and print the result
	// or print "timeout 1" if the operation takes more than 1 second


	// If we allow a longer timeout of 3s, then the receive
	// from `c2` will succeed and we'll print the result.

	// TODO: Create c2 channel of strings with buffer size 1

	// TODO: Creat a goroutine that sends "result 2" to c2 after 2 seconds

	// TODO: Use select to receive from c2 and print the result
	// or print "timeout 2" if the operation takes more than 3 seconds

}"""
    },
    {
        "key": "non-blocking-channel-operations",
        "display_name": "Non-Blocking Channel Operations",
        "template": """// Basic sends and receives on channels are blocking.
// However, we can use `select` with a `default` clause to
// implement _non-blocking_ sends, receives, and even
// non-blocking multi-way `select`s.

package main

import "fmt"

func main() {
	
	// TODO: Create messages channel of strings

	// TODO: Create signals channel of bools

	// Here's a non-blocking receive. If a value is
	// available on `messages` then `select` will take
	// the `<-messages` `case` with that value. If not
	// it will immediately take the `default` case.

	// TODO: Use select to receive from messages and print the result
	// or print "no message received" if the operation takes more than 1 second
	

	// A non-blocking send works similarly. Here `msg`
	// cannot be sent to the `messages` channel, because
	// the channel has no buffer and there is no receiver.
	// Therefore the `default` case is selected.

	// TODO: Use select to send to messages and print the result
	// or print "no message sent" as default

	// We can use multiple `case`s above the `default`
	// clause to implement a multi-way non-blocking
	// select. Here we attempt non-blocking receives
	// on both `messages` and `signals`.

	// TODO: Use select to receive from messages and signals and print the result
	// or print "no activity" as default

}"""
    },
    {
        "key": "closing-channels",
        "display_name": "Closing Channels",
        "template": """// _Closing_ a channel indicates that no more values
// will be sent on it. This can be useful to communicate
// completion to the channel's receivers.

package main

import "fmt"

// In this example we'll use a `jobs` channel to
// communicate work to be done from the `main()` goroutine
// to a worker goroutine. When we have no more jobs for
// the worker we'll `close` the `jobs` channel.
func main() {
	
	// TODO: Create jobs channel of int with buffer size 5

	// TODO: Create done channel of bool


	// Here's the worker goroutine. It repeatedly receives
	// from `jobs` with `j, more := <-jobs`. In this
	// special 2-value form of receive, the `more` value
	// will be `false` if `jobs` has been `close`d and all
	// values in the channel have already been received.
	// We use this to notify on `done` when we've worked
	// all our jobs.

	// TODO: Creat a goroutine that receives from jobs and prints the result
	// or prints "received all jobs" if the operation takes more than 1 second
	

	// This sends 3 jobs to the worker over the `jobs`
	// channel, then closes it.

	// TODO: Send 3 jobs to the worker over the jobs channel

	// TODO: Close the jobs channel

	fmt.Println("sent all jobs")

	// We await the worker using the
	// [synchronization](channel-synchronization) approach
	// we saw earlier.

	// TODO: Receive from done channel

	// Reading from a closed channel succeeds immediately,
	// returning the zero value of the underlying type.
	// The optional second return value is `true` if the
	// value received was delivered by a successful send
	// operation to the channel, or `false` if it was a
	// zero value generated because the channel is closed
	// and empty.
	
	// TODO: Receive from jobs

	// TODO: Print the result

}"""
    },
    {
        "key": "range-over-channels",
        "display_name": "Range over Channels",
        "template": """// In a [previous](range-over-built-in-types) example we saw how `for` and
// `range` provide iteration over basic data structures.
// We can also use this syntax to iterate over
// values received from a channel.

package main

import "fmt"

func main() {

	// We'll iterate over 2 values in the `queue` channel.

	// TODO: Create queue channel of strings with buffer size 2

	// TODO: Send "one" and "two" to queue channel

	// TODO: Close the queue channel


	// This `range` iterates over each element as it's
	// received from `queue`. Because we `close`d the
	// channel above, the iteration terminates after
	// receiving the 2 elements.

	// TODO: Use range to iterate over queue and print each element

}"""
    },
    {
        "key": "timers",
        "display_name": "Timers",
        "template": """// We often want to execute Go code at some point in the
// future, or repeatedly at some interval. Go's built-in
// _timer_ and _ticker_ features make both of these tasks
// easy. We'll look first at timers and then
// at [tickers](tickers).

package main

import (
	"fmt"
	"time"
)

func main() {

	// Timers represent a single event in the future. You
	// tell the timer how long you want to wait, and it
	// provides a channel that will be notified at that
	// time. This timer will wait 2 seconds.

	// TODO: Create timer1 with 2 seconds

	// The `<-timer1.C` blocks on the timer's channel `C`
	// until it sends a value indicating that the timer
	// fired.

	// TODO: Receive from timer1.C and print the result

	// If you just wanted to wait, you could have used
	// `time.Sleep`. One reason a timer may be useful is
	// that you can cancel the timer before it fires.
	// Here's an example of that.

	// TODO: Create timer2 with 1 second using NewTimer

	// TODO: Creat a goroutine that receives from timer2.C and prints the result

	// TODO: Create stop2 with timer2.Stop()

	// TODO: If stop2 is true, print "Timer 2 stopped"

	// TODO: Give the timer2 enough time to fire, if it ever
	// was going to, to show it is in fact stopped.

}"""
    },
    {
        "key": "tickers",
        "display_name": "Tickers",
        "template": """// [Timers](timers) are for when you want to do
// something once in the future - _tickers_ are for when
// you want to do something repeatedly at regular
// intervals. Here's an example of a ticker that ticks
// periodically until we stop it.

package main

import (
	"fmt"
	"time"
)

func main() {

	// Tickers use a similar mechanism to timers: a
	// channel that is sent values. Here we'll use the
	// `select` builtin on the channel to await the
	// values as they arrive every 500ms.

	// TODO: Create ticker with 500 milliseconds using NewTicker

	// TODO: Create done channel of bool

	// TODO: Creat a goroutine that receives from ticker.C and prints the result

	// Tickers can be stopped like timers. Once a ticker
	// is stopped it won't receive any more values on its
	// channel. We'll stop ours after 1600ms.

	// TODO: Give the ticker enough time to fire, if it ever
	// was going to, to show it is in fact stopped.

	// TODO: If done is true, print "Ticker stopped"

}"""
    },
    {
        "key": "worker-pools",
        "display_name": "Worker Pools",
        "template": """// In this example we'll look at how to implement
// a _worker pool_ using goroutines and channels.

package main

import (
	"fmt"
	"time"
)

// Here's the worker, of which we'll run several
// concurrent instances. These workers will receive
// work on the `jobs` channel and send the corresponding
// results on `results`. We'll sleep a second per job to
// simulate an expensive task.

// TODO: Define function worker(id int, jobs <-chan int, results chan<- int) that receives from jobs and sends the result to results

func main() {

	// In order to use our pool of workers we need to send
	// them work and collect their results. We make 2
	// channels for this.

	// TODO: Create numJobs = 5

	// TODO: Create jobs channel of int with buffer size numJobs

	// TODO: Create results channel of int with buffer size numJobs

	// This starts up 3 workers, initially blocked
	// because there are no jobs yet.

	// TODO: Create 3 workers using worker function

	// Here we send 5 `jobs` and then `close` that
	// channel to indicate that's all the work we have.

	// TODO: Send 5 jobs to the workers
	// TODO: Close the jobs channel

	// Finally we collect all the results of the work.
	// This also ensures that the worker goroutines have
	// finished. An alternative way to wait for multiple
	// goroutines is to use a [WaitGroup](waitgroups).

	// TODO: Receive from results channel

}"""
    },
    {
        "key": "waitgroups",
        "display_name": "WaitGroups",
        "template": """// To wait for multiple goroutines to finish, we can
// use a *wait group*.

package main

import (
	"fmt"
	"sync"
	"time"
)

// This is the function we'll run in every goroutine.

// TODO: Define function worker(id int) that prints "Worker %d starting" and "Worker %d done" after sleeping for 1 second

func main() {

	// This WaitGroup is used to wait for all the
	// goroutines launched here to finish. Note: if a WaitGroup is
	// explicitly passed into functions, it should be done *by pointer*.

	// TODO: Create wg sync.WaitGroup

	// Launch several goroutines using `WaitGroup.Go`

	// TODO: Launch 5 workers using worker function

	// Block until all the goroutines started by `wg` are
	// done. A goroutine is done when the function it invokes
	// returns.

	// TODO: Wait for all the goroutines to finish
	wg.Wait()

	// Note that this approach has no straightforward way
	// to propagate errors from workers. For more
	// advanced use cases, consider using the
	// [errgroup package](https://pkg.go.dev/golang.org/x/sync/errgroup).

}"""
    },
    {
        "key": "rate-limiting",
        "display_name": "Rate Limiting",
        "template": """// [_Rate limiting_](https://en.wikipedia.org/wiki/Rate_limiting)
// is an important mechanism for controlling resource
// utilization and maintaining quality of service. Go
// elegantly supports rate limiting with goroutines,
// channels, and [tickers](tickers).

package main

import (
	"fmt"
	"time"
)

func main() {

	// First we'll look at basic rate limiting. Suppose
	// we want to limit our handling of incoming requests.
	// We'll serve these requests off a channel of the
	// same name.

	// TODO: Create requests channel of int with buffer size 5

	// TODO: Send 5 requests to the requests channel

	// TODO: Close the requests channel

	// This `limiter` channel will receive a value
	// every 200 milliseconds. This is the regulator in
	// our rate limiting scheme.

	// TODO: Create limiter channel of time.Tick 200 milliseconds

	// By blocking on a receive from the `limiter` channel
	// before serving each request, we limit ourselves to
	// 1 request every 200 milliseconds.

	// TODO: Iterate over requests channel and retrieve the limiter channel and print the request and time

	

	// We may want to allow short bursts of requests in
	// our rate limiting scheme while preserving the
	// overall rate limit. We can accomplish this by
	// buffering our limiter channel. This `burstyLimiter`
	// channel will allow bursts of up to 3 events.

	// TODO: Create burstyLimiter channel of time.Time with buffer size 3

	// Fill up the channel to represent allowed bursting.

	// TODO: Iterate over 3 and send the time to the burstyLimiter channel

	// Every 200 milliseconds we'll try to add a new
	// value to `burstyLimiter`, up to its limit of 3.

	// TODO: Creat a goroutine that sends the time to the burstyLimiter channel every 200 milliseconds
	

	// Now simulate 5 more incoming requests. The first
	// 3 of these will benefit from the burst capability
	// of `burstyLimiter`.

	// TODO: Create burstyRequests channel of int with buffer size 5

	// TODO: Send 5 requests to the burstyRequests channel

	// TODO: Close the burstyRequests channel

	// TODO: Iterate over burstyRequests channel and retrieve the burstyLimiter channel and print the request and time

}"""
    },
    {
        "key": "atomic-counters",
        "display_name": "Atomic Counters",
        "template": """// The primary mechanism for managing state in Go is
// communication over channels. We saw this for example
// with [worker pools](worker-pools). There are a few other
// options for managing state though. Here we'll
// look at using the `sync/atomic` package for _atomic
// counters_ accessed by multiple goroutines.

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	// We'll use an atomic integer type to represent our
	// (always-positive) counter.

	// TODO: Create ops atomic.Uint64

	// A WaitGroup will help us wait for all goroutines
	// to finish their work.

	// TODO: Create wg sync.WaitGroup

	// We'll start 50 goroutines that each increment the
	// counter exactly 1000 times.

	// TODO: Iterate over 50 and create a goroutine that increments the counter exactly 1000 times

	// Wait until all the goroutines are done.

	// TODO: Wait for all the goroutines to finish

	// Here no goroutines are writing to 'ops', but using
	// `Load` it's safe to atomically read a value even while
	// other goroutines are (atomically) updating it.

	// TODO: Print the result of the counter with ops.Load()
}"""
    },
    {
        "key": "mutexes",
        "display_name": "Mutexes",
        "template": """// In the previous example we saw how to manage simple
// counter state using [atomic operations](atomic-counters).
// For more complex state we can use a [_mutex_](https://en.wikipedia.org/wiki/Mutual_exclusion)
// to safely access data across multiple goroutines.

package main

import (
	"fmt"
	"sync"
)

// Container holds a map of counters; since we want to
// update it concurrently from multiple goroutines, we
// add a `Mutex` to synchronize access.
// Note that mutexes must not be copied, so if this
// `struct` is passed around, it should be done by
// pointer.

// TODO: Define struct Container with mu (sync.Mutex) and counters (map[string]int) fields

// TODO: Create method inc(name string) on Container that locks the mutex and increments the counter for the given name
// Lock the mutex before accessing `counters`; unlock
// it at the end of the function using a [defer](defer)
// statement.


func main() {
	
	// Note that the zero value of a mutex is usable as-is, so no
	// initialization is required here.

	// TODO: Create c Container with counters map[string]int{"a": 0, "b": 0}

	// TODO: Create wg sync.WaitGroup

	// This function increments a named counter
	// in a loop.

	// TODO: Define function doIncrement(name string, n int) that increments the counter for the given name in a loop

	// Run several goroutines concurrently; note
	// that they all access the same `Container`,
	// and two of them access the same counter.

	// TODO: Launch 3 goroutines using wg.Go that call doIncrement with "a" and 10000, "a" and 10000, and "b" and 10000

	// Wait for the goroutines to finish

	// TODO: Wait for all the goroutines to finish.

	// TODO: Print the result of the counters with c.counters
}"""
    },
    {
        "key": "stateful-goroutines",
        "display_name": "Stateful Goroutines",
        "template": """// In the previous example we used explicit locking with
// [mutexes](mutexes) to synchronize access to shared state
// across multiple goroutines. Another option is to use the
// built-in synchronization features of  goroutines and
// channels to achieve the same result. This channel-based
// approach aligns with Go's ideas of sharing memory by
// communicating and having each piece of data owned
// by exactly 1 goroutine.

package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

// In this example our state will be owned by a single
// goroutine. This will guarantee that the data is never
// corrupted with concurrent access. In order to read or
// write that state, other goroutines will send messages
// to the owning goroutine and receive corresponding
// replies. These `readOp` and `writeOp` `struct`s
// encapsulate those requests and a way for the owning
// goroutine to respond.

// TODO: Define struct readOp with key (int) and resp (chan int) fields

// TODO: Define struct writeOp with key (int), val (int), and resp (chan bool) fields


func main() {

	// As before we'll count how many operations we perform.

	// TODO: Create readOps uint64

	// TODO: Create writeOps uint64


	// The `reads` and `writes` channels will be used by
	// other goroutines to issue read and write requests,
	// respectively.

	// TODO: Create reads channel of readOp

	// TODO: Create writes channel of writeOp

	// Here is the goroutine that owns the `state`, which
	// is a map as in the previous example but now private
	// to the stateful goroutine. This goroutine repeatedly
	// selects on the `reads` and `writes` channels,
	// responding to requests as they arrive. A response
	// is executed by first performing the requested
	// operation and then sending a value on the response
	// channel `resp` to indicate success (and the desired
	// value in the case of `reads`).

	// TODO: Create a goroutine that owns the state, which is a map as in the previous example but now private to the stateful goroutine.
	// Inside, use for range to select on reads and writes channels and perform the requested operation and send a value on the response channel resp to indicate success (and the desired value in the case of reads).


	// This starts 100 goroutines to issue reads to the
	// state-owning goroutine via the `reads` channel.
	// Each read requires constructing a `readOp`, sending
	// it over the `reads` channel, and then receiving the
	// result over the provided `resp` channel.


	// TODO: Iterate over 100 and create a goroutine that issues reads to the state-owning goroutine via the reads channel.
	// Inside, create read readOp with key rand.Intn(5) and resp make(chan int)
	// Send read to reads channel
	// Receive the result from read.resp
	// Add 1 to readOps
	// Sleep for 1 millisecond


	// We start 10 writes as well, using a similar
	// approach.


	// TODO: Iterate over 10 and create a goroutine that issues writes to the state-owning goroutine via the writes channel.
	// Inside, create write writeOp with key rand.Intn(5) and val rand.Intn(100) and resp make(chan bool)
	// Send write to writes channel
	// Receive the result from write.resp
	// Add 1 to writeOps
	// Sleep for 1 millisecond


	// Let the goroutines work for a second.

	// TODO: Sleep for 1 second

	// Finally, capture and report the op counts.

	// TODO: Print the result of the reads with readOps

	// TODO: Print the result of the writes with writeOps

}"""
    },
    {
        "key": "sorting",
        "display_name": "Sorting",
        "template": """// Go's `slices` package implements sorting for builtins
// and user-defined types. We'll look at sorting for
// builtins first.

package main

import (
	"fmt"
	"slices"
)

func main() {

	// Sorting functions are generic, and work for any
	// _ordered_ built-in type. For a list of ordered
	// types, see [cmp.Ordered](https://pkg.go.dev/cmp#Ordered).

	// TODO: Create slice strs of strings with values "c", "a", "b"

	// TODO: Sort strs using slices.Sort

	// TODO: Print Strings: strs


	// An example of sorting `int`s.

	// TODO: Create slice ints of ints with values 7, 2, 4

	// TODO: Sort ints using slices.Sort

	// TODO: Print Ints: ints


	// We can also use the `slices` package to check if
	// a slice is already in sorted order.

	// TODO: Check if ints is sorted using slices.IsSorted

	// TODO: Print Sorted: s

}"""
    },
    {
        "key": "sorting-by-functions",
        "display_name": "Sorting by Functions",
        "template": """// Sometimes we'll want to sort a collection by something
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

}"""
    },
    {
        "key": "panic",
        "display_name": "Panic",
        "template": """// A `panic` typically means something went unexpectedly
// wrong. Mostly we use it to fail fast on errors that
// shouldn't occur during normal operation, or that we
// aren't prepared to handle gracefully.

package main

import "os"

func main() {

	// We'll use panic throughout this site to check for
	// unexpected errors. This is the only program on the
	// site designed to panic.

	// TODO: Panic with "a problem"


	// A common use of panic is to abort if a function
	// returns an error value that we don't know how to
	// (or want to) handle. Here's an example of
	// `panic`king if we get an unexpected error when creating a new file.

	// TODO: Create a new file with os.Create("/tmp/file")
	// TODO: Check if err is not nil, panic with err

}"""
    },
    {
        "key": "defer",
        "display_name": "Defer",
        "template": """// _Defer_ is used to ensure that a function call is
// performed later in a program's execution, usually for
// purposes of cleanup. `defer` is often used where e.g.
// `ensure` and `finally` would be used in other languages.

package main

import (
	"fmt"
	"os"
)

// Suppose we wanted to create a file, write to it,
// and then close when we're done. Here's how we could
// do that with `defer`.
func main() {

	// Immediately after getting a file object with
	// `createFile`, we defer the closing of that file
	// with `closeFile`. This will be executed at the end
	// of the enclosing function (`main`), after
	// `writeFile` has finished.

	// TODO: Create a file with createFile("/tmp/defer.txt")
	// TODO: Defer the closing of the file with defer closeFile(f)
	// TODO: Write to the file with writeFile(f)

}

// TODO: Create function createFile(p string) *os.File that creates a file at the given path and returns the file
// Inside, create a file with os.Create(p) and check if err is not nil, panic with err
// Return the file

// TODO: Create function writeFile(f *os.File) that writes "data" to the file
// Inside, use fmt.Fprintln(f, "data")


// TODO: Create function closeFile(f *os.File) that closes the file
// Inside, use f.Close() and check if err is not nil, panic with err

// It's important to check for errors when closing a
// file, even in a deferred function."""
    },
    {
        "key": "recover",
        "display_name": "Recover",
        "template": """// Go makes it possible to _recover_ from a panic, by
// using the `recover` built-in function. A `recover` can
// stop a `panic` from aborting the program and let it
// continue with execution instead.

// An example of where this can be useful: a server
// wouldn't want to crash if one of the client connections
// exhibits a critical error. Instead, the server would
// want to close that connection and continue serving
// other clients. In fact, this is what Go's `net/http`
// does by default for HTTP servers.

package main

import "fmt"

// This function panics.

// TODO: Create function mayPanic that panics with "a problem"

func main() {
	// `recover` must be called within a deferred function.
	// When the enclosing function panics, the defer will
	// activate and a `recover` call within it will catch
	// the panic.

	// TODO: Defer a function that recovers from a panic
	// Inside, check if recover is not nil, print the error

	// The return value of `recover` is the error raised in
	// the call to `panic`.

	// TODO: Call mayPanic

	// This code will not run, because `mayPanic` panics.
	// The execution of `main` stops at the point of the
	// panic and resumes in the deferred closure.

	// TODO: Print "After mayPanic()"
}"""
    },
    {
        "key": "string-functions",
        "display_name": "String Functions",
        "template": """// The standard library's `strings` package provides many
// useful string-related functions. Here are some examples
// to give you a sense of the package.

package main

import (
	"fmt"
	s "strings"
)

// We alias `fmt.Println` to a shorter name as we'll use
// it a lot below.
var p = fmt.Println

func main() {

	// Here's a sample of the functions available in
	// `strings`. Since these are functions from the
	// package, not methods on the string object itself,
	// we need to pass the string in question as the first
	// argument to the function. You can find more
	// functions in the [`strings`](https://pkg.go.dev/strings)
	// package docs.

	// TODO: Print Contains:  s.Contains("test", "es")
	// TODO: Print Count:     s.Count("test", "t")
	// TODO: Print HasPrefix: s.HasPrefix("test", "te")
	// TODO: Print HasSuffix: s.HasSuffix("test", "st")
	// TODO: Print Index:     s.Index("test", "e")
	// TODO: Print Join:      s.Join([]string{"a", "b"}, "-")
	// TODO: Print Repeat:    s.Repeat("a", 5)
	// TODO: Print Replace:   s.Replace("foo", "o", "0", -1)
	// TODO: Print Split:     s.Split("a-b-c-d-e", "-")
	// TODO: Print ToLower:   s.ToLower("TEST")
	// TODO: Print ToUpper:   s.ToUpper("test")

}"""
    },
    {
        "key": "string-formatting",
        "display_name": "String Formatting",
        "template": """// Go offers excellent support for string formatting in
// the `printf` tradition. Here are some examples of
// common string formatting tasks.

package main

import (
	"fmt"
	"os"
)

// TODO: Define struct point with x (int) and y (int) fields

func main() {

	// Go offers several printing "verbs" designed to
	// format general Go values. For example, this prints
	// an instance of our `point` struct.

	// TODO: Create p := point{1, 2}
	// TODO: Print p using %v

	// If the value is a struct, the `%+v` variant will
	// include the struct's field names.

	// TODO: Print p using %+v

	// The `%#v` variant prints a Go syntax representation
	// of the value, i.e. the source code snippet that
	// would produce that value.

	// TODO: Print p using %#v

	// To print the type of a value, use `%T`.

	// TODO: Print p using %T

	// Formatting booleans is straight-forward.

	// TODO: Print true using %t

	// There are many options for formatting integers.
	// Use `%d` for standard, base-10 formatting.

	// TODO: Print 123 using %d

	// This prints a binary representation.
	// TODO: Print 14 using %b

	// This prints the character corresponding to the
	// given integer.

	// TODO: Print 33 using %c

	// `%x` provides hex encoding.
	// TODO: Print 456 using %x

	// There are also several formatting options for
	// floats. For basic decimal formatting use `%f`.

	// TODO: Print 78.9 using %f

	// `%e` and `%E` format the float in (slightly
	// different versions of) scientific notation.

	// TODO: Print 123400000.0 using %e
	// TODO: Print 123400000.0 using %E

	// For basic string printing use `%s`.

	// TODO: Print "\"string\"" using %s

	// To double-quote strings as in Go source, use `%q`.

	// TODO: Print "\"string\"" using %q

	// As with integers seen earlier, `%x` renders
	// the string in base-16, with two output characters
	// per byte of input.

	// TODO: Print "hex this" using %x

	// To print a representation of a pointer, use `%p`.

	// TODO: Print pointer of p using %p

	// When formatting numbers you will often want to
	// control the width and precision of the resulting
	// figure. To specify the width of an integer, use a
	// number after the `%` in the verb. By default the
	// result will be right-justified and padded with
	// spaces.

	// TODO: Print "12" and "345" using %6d

	// You can also specify the width of printed floats,
	// though usually you'll also want to restrict the
	// decimal precision at the same time with the
	// width.precision syntax.

	// TODO: Print "1.2" and "3.45" using %6.2f

	// To left-justify, use the `-` flag.

	// TODO: Print "1.2" and "3.45" using %-6.2f

	// You may also want to control width when formatting
	// strings, especially to ensure that they align in
	// table-like output. For basic right-justified width.

	// TODO: Print "foo" and "b" using %6s

	// To left-justify use the `-` flag as with numbers.

	// TODO: Print "foo" and "b" using %-6s

	// So far we've seen `Printf`, which prints the
	// formatted string to `os.Stdout`. `Sprintf` formats
	// and returns a string without printing it anywhere.

	// TODO: Print "sprintf: a string" using %Sprintf

	// You can format+print to `io.Writers` other than
	// `os.Stdout` using `Fprintf`.

	// TODO: Print "io: an error" using %Fprintf
}"""
    },
    {
        "key": "text-templates",
        "display_name": "Text Templates",
        "template": """// Go offers built-in support for creating dynamic content or showing customized
// output to the user with the `text/template` package. A sibling package
// named `html/template` provides the same API but has additional security
// features and should be used for generating HTML.

package main

import (
	"os"
	"text/template"
)

func main() {

	// We can create a new template and parse its body from
	// a string.
	// Templates are a mix of static text and "actions" enclosed in
	// `{{...}}` that are used to dynamically insert content.

	// TODO: Create t1 := template.New("t1")
	// TODO: Parse t1 with "Value is {{.}}\n" and handle error

	// Alternatively, we can use the `template.Must` function to
	// panic in case `Parse` returns an error. This is especially
	// useful for templates initialized in the global scope.

	// TODO: Use template.Must to parse t1 with "Value: {{.}}\n" and handle error

	// By "executing" the template we generate its text with
	// specific values for its actions. The `{{.}}` action is
	// replaced by the value passed as a parameter to `Execute`.

	// TODO: Execute t1 with "some text"
	// TODO: Execute t1 with 5
	// TODO: Execute t1 with []string{"Go", "Rust", "C++", "C#"}
	

	// Helper function we'll use below.

	// TODO: Create Create function with name and t string that returns template.Must(template.New(name).Parse(t))
	

	// If the data is a struct we can use the `{{.FieldName}}` action to access
	// its fields. The fields should be exported to be accessible when a
	// template is executing.

	// TODO: Create t2 := Create("t2", "Name: {{.Name}}\n")

	// TODO: Execute t2 with struct {Name string}{"Jane Doe"}


	// The same applies to maps; with maps there is no restriction on the
	// case of key names.

	// TODO: Execute t2 with map[string]string{"Name": "Mickey Mouse"}

	// if/else provide conditional execution for templates. A value is considered
	// false if it's the default value of a type, such as 0, an empty string,
	// nil pointer, etc.
	// This sample demonstrates another
	// feature of templates: using `-` in actions to trim whitespace.


	// TODO: Create t3 := Create("t3", "{{if . -}} yes {{else -}} no {{end}}\n")
	// TODO: Execute t3 with "not empty"
	// TODO: Execute t3 with ""

	// range blocks let us loop through slices, arrays, maps or channels. Inside
	// the range block `{{.}}` is set to the current item of the iteration.

	// TODO: Create t4 := Create("t4", "Range: {{range .}}{{.}} {{end}}\n")
	// TODO: Execute t4 with []string{"Go", "Rust", "C++", "C#"}
	
}"""
    },
    {
        "key": "regular-expressions",
        "display_name": "Regular Expressions",
        "template": """// Go offers built-in support for [regular expressions](https://en.wikipedia.org/wiki/Regular_expression).
// Here are some examples of  common regexp-related tasks
// in Go.

package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {

	// This tests whether a pattern matches a string.

	// TODO: Create match, _ := regexp.MatchString("p([a-z]+)ch", "peach") and print it

	// Above we used a string pattern directly, but for
	// other regexp tasks you'll need to `Compile` an
	// optimized `Regexp` struct.

	// TODO: Create r, _ := regexp.Compile("p([a-z]+)ch")

	// Many methods are available on these structs. Here's
	// a match test like we saw earlier.

	// TODO: Print r.MatchString("peach")

	// This finds the match for the regexp.
	// TODO: Print r.FindString("peach punch")

	// This also finds the first match but returns the
	// start and end indexes for the match instead of the
	// matching text.

	// TODO: Print r.FindStringIndex("peach punch")

	// The `Submatch` variants include information about
	// both the whole-pattern matches and the submatches
	// within those matches. For example this will return
	// information for both `p([a-z]+)ch` and `([a-z]+)`.

	// TODO: Print r.FindStringSubmatch("peach punch")

	// Similarly this will return information about the
	// indexes of matches and submatches.

	// TODO: Print r.FindStringSubmatchIndex("peach punch")

	// The `All` variants of these functions apply to all
	// matches in the input, not just the first. For
	// example to find all matches for a regexp.

	// TODO: Print r.FindAllString("peach punch pinch", -1)

	// These `All` variants are available for the other
	// functions we saw above as well.

	// TODO: Print r.FindAllStringSubmatchIndex("peach punch pinch", -1)

	// Providing a non-negative integer as the second
	// argument to these functions will limit the number
	// of matches.

	// TODO: Print r.FindAllString("peach punch pinch", 2)

	// Our examples above had string arguments and used
	// names like `MatchString`. We can also provide
	// `[]byte` arguments and drop `String` from the
	// function name.

	// TODO: Print r.Match([]byte("peach"))

	// When creating global variables with regular
	// expressions you can use the `MustCompile` variation
	// of `Compile`. `MustCompile` panics instead of
	// returning an error, which makes it safer to use for
	// global variables.

	// TODO: Create r = regexp.MustCompile("p([a-z]+)ch")
	// TODO: Print r

	// The `regexp` package can also be used to replace
	// subsets of strings with other values.

	// TODO: Print r.ReplaceAllString("a peach", "<fruit>")

	// The `Func` variant allows you to transform matched
	// text with a given function.

	// TODO: Create in := []byte("a peach")
	// TODO: Create out := r.ReplaceAllFunc(in, bytes.ToUpper)
	// TODO: Print string(out)

}"""
    },
    {
        "key": "json",
        "display_name": "JSON",
        "template": """// Go offers built-in support for JSON encoding and
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

}"""
    },
    {
        "key": "xml",
        "display_name": "XML",
        "template": """// Go offers built-in support for XML and XML-like
// formats with the `encoding/xml` package.

package main

import (
	"encoding/xml"
	"fmt"
)

// Plant will be mapped to XML. Similarly to the
// JSON examples, field tags contain directives for the
// encoder and decoder. Here we use some special features
// of the XML package: the `XMLName` field name dictates
// the name of the XML element representing this struct;
// `id,attr` means that the `Id` field is an XML
// _attribute_ rather than a nested element.

// TODO: Create Plant struct with XMLName (xml.Name), Id (int), Name (string), and Origin ([]string) fields


// TODO: Create String method for Plant struct that returns a string with the Id, Name, and Origin
// TODO: Return a string with the Id, Name, and Origin

func main() {

	// TODO: Create coffee := &Plant{Id: 27, Name: "Coffee"}
	// TODO: Create coffee.Origin = []string{"Ethiopia", "Brazil"}


	// Emit XML representing our plant; using
	// `MarshalIndent` to produce a more
	// human-readable output.

	// TODO: Create out, _ := xml.MarshalIndent(coffee, " ", "  ")
	// TODO: Print string(out)

	// To add a generic XML header to the output, append
	// it explicitly.

	// TODO: Create fmt.Println(xml.Header + string(out))

	// Use `Unmarshal` to parse a stream of bytes with XML
	// into a data structure. If the XML is malformed or
	// cannot be mapped onto Plant, a descriptive error
	// will be returned.

	// TODO: Create var p Plant
	// TODO: Create if err := xml.Unmarshal(out, &p); err != nil {
	// TODO: Print err

	// TODO: Create tomato := &Plant{Id: 81, Name: "Tomato"}
	// TODO: Create tomato.Origin = []string{"Mexico", "California"}


	// The `parent>child>plant` field tag tells the encoder
	// to nest all `plant`s under `<parent><child>...`

	// TODO: Create Nesting struct with XMLName (xml.Name), Plants ([]*Plant) fields	


	// TODO: Create nesting := &Nesting{}
	// TODO: Create nesting.Plants = []*Plant{coffee, tomato}

	// TODO: Create out, _ = xml.MarshalIndent(nesting, " ", "  ")
	// TODO: Create fmt.Println(string(out))
}"""
    },
    {
        "key": "time",
        "display_name": "Time",
        "template": """// Go offers extensive support for times and durations;
// here are some examples.

package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	// We'll start by getting the current time.
	now := time.Now()
	p(now)

	// You can build a `time` struct by providing the
	// year, month, day, etc. Times are always associated
	// with a `Location`, i.e. time zone.

	// TODO: Create then := time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	// TODO: Print then

	// You can extract the various components of the time
	// value as expected.

	// TODO: Print then.Year()
	// TODO: Print then.Month()
	// TODO: Print then.Day()
	// TODO: Print then.Hour()
	// TODO: Print then.Minute()
	// TODO: Print then.Second()
	// TODO: Print then.Nanosecond()
	// TODO: Print then.Location()

	// The Monday-Sunday `Weekday` is also available.

	// TODO: Print then.Weekday()

	// These methods compare two times, testing if the
	// first occurs before, after, or at the same time
	// as the second, respectively.

	// TODO: Print then.Before(now)
	// TODO: Print then.After(now)
	// TODO: Print then.Equal(now)

	// The `Sub` methods returns a `Duration` representing
	// the interval between two times.

	// TODO: Create diff := now.Sub(then)
	// TODO: Print diff

	// We can compute the length of the duration in
	// various units.

	// TODO: Print diff.Hours()
	// TODO: Print diff.Minutes()
	// TODO: Print diff.Seconds()
	// TODO: Print diff.Nanoseconds()

	// You can use `Add` to advance a time by a given
	// duration, or with a `-` to move backwards by a
	// duration.

	// TODO: Print then.Add(diff)
	// TODO: Print then.Add(-diff)
}"""
    },
    {
        "key": "epoch",
        "display_name": "Epoch",
        "template": """// A common requirement in programs is getting the number
// of seconds, milliseconds, or nanoseconds since the
// [Unix epoch](https://en.wikipedia.org/wiki/Unix_time).
// Here's how to do it in Go.

package main

import (
	"fmt"
	"time"
)

func main() {

	// Use `time.Now` with `Unix`, `UnixMilli` or `UnixNano`
	// to get elapsed time since the Unix epoch in seconds,
	// milliseconds or nanoseconds, respectively.

	// TODO: Create now := time.Now() and print it

	// TODO: Print now.Unix()
	// TODO: Print now.UnixMilli()
	// TODO: Print now.UnixNano()

	// You can also convert integer seconds or nanoseconds
	// since the epoch into the corresponding `time`.

	// TODO: Print time.Unix(now.Unix(), 0)
	// TODO: Print time.Unix(0, now.UnixNano())
}"""
    },
    {
        "key": "time-formatting-parsing",
        "display_name": "Time Formatting / Parsing",
        "template": """// Go supports time formatting and parsing via
// pattern-based layouts.

package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	// Here's a basic example of formatting a time
	// according to RFC3339, using the corresponding layout
	// constant.

	// TODO: Create t := time.Now()
	// TODO: Print t.Format(time.RFC3339)
	

	// Time parsing uses the same layout values as `Format`.

	// TODO: Create t1, e := time.Parse(time.RFC3339, "2012-11-01T22:08:41+00:00")
	// TODO: Print t1

	// `Format` and `Parse` use example-based layouts. Usually
	// you'll use a constant from `time` for these layouts, but
	// you can also supply custom layouts. Layouts must use the
	// reference time `Mon Jan 2 15:04:05 MST 2006` to show the
	// pattern with which to format/parse a given time/string.
	// The example time must be exactly as shown: the year 2006,
	// 15 for the hour, Monday for the day of the week, etc.

	// TODO: Print t.Format("3:04PM")
	// TODO: Print t.Format("Mon Jan _2 15:04:05 2006")
	// TODO: Print t.Format("2006-01-02T15:04:05.999999-07:00")
	

	// TODO: Create form := "3 04 PM"
	// TODO: Create t2, e := time.Parse(form, "8 41 PM")
	// TODO: Print t2

	// For purely numeric representations you can also
	// use standard string formatting with the extracted
	// components of the time value.

	// TODO: Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00")
	// TODO: with t.Year(), t.Month(), t.Day(),
	// t.Hour(), t.Minute(), t.Second()

	// `Parse` will return an error on malformed input
	// explaining the parsing problem.

	// TODO: Create ansic := "Mon Jan _2 15:04:05 2006"
	// TODO: Create _, e = time.Parse(ansic, "8:41PM")
	// TODO: Print e
}"""
    },
    {
        "key": "random-numbers",
        "display_name": "Random Numbers",
        "template": """// Go's `math/rand/v2` package provides
// [pseudorandom number](https://en.wikipedia.org/wiki/Pseudorandom_number_generator)
// generation.

package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {

	// For example, `rand.IntN` returns a random `int` n,
	// `0 <= n < 100`.

	// TODO: Print rand.IntN(100), ","
	// TODO: Print rand.IntN(100)

	// `rand.Float64` returns a `float64` `f`,
	// `0.0 <= f < 1.0`.

	// TODO: Print rand.Float64()

	// This can be used to generate random floats in
	// other ranges, for example `5.0 <= f' < 10.0`.

	// TODO: Print (rand.Float64()*5)+5, ","
	// TODO: Print (rand.Float64() * 5) + 5

	// If you want a known seed, create a new
	// `rand.Source` and pass it into the `New`
	// constructor. `NewPCG` creates a new
	// [PCG](https://en.wikipedia.org/wiki/Permuted_congruential_generator)
	// source that requires a seed of two `uint64`
	// numbers.

	// TODO: Create s2 := rand.NewPCG(42, 1024)
	// TODO: Create r2 := rand.New(s2)
	// TODO: Print r2.IntN(100), ","
	// TODO: Print r2.IntN(100)

	// TODO: Create s3 := rand.NewPCG(42, 1024)
	// TODO: Create r3 := rand.New(s3)
	// TODO: Print r3.IntN(100), ","
	// TODO: Print r3.IntN(100)
}"""
    },
    {
        "key": "number-parsing",
        "display_name": "Number Parsing",
        "template": """// Parsing numbers from strings is a basic but common task
// in many programs; here's how to do it in Go.

package main

// The built-in package `strconv` provides the number
// parsing.
import (
	"fmt"
	"strconv"
)

func main() {

	// With `ParseFloat`, this `64` tells how many bits of
	// precision to parse.

	// TODO: Create f, _ := strconv.ParseFloat("1.234", 64)
	// TODO: Print f

	// For `ParseInt`, the `0` means infer the base from
	// the string. `64` requires that the result fit in 64
	// bits.

	// TODO: Create i, _ := strconv.ParseInt("123", 0, 64)
	// TODO: Print i

	// `ParseInt` will recognize hex-formatted numbers.

	// TODO: Create d, _ := strconv.ParseInt("0x1c8", 0, 64)
	// TODO: Print d

	// A `ParseUint` is also available.

	// TODO: Create u, _ := strconv.ParseUint("789", 0, 64)
	// TODO: Print u

	// `Atoi` is a convenience function for basic base-10
	// `int` parsing.

	// TODO: Create k, _ := strconv.Atoi("135")
	// TODO: Print k

	// Parse functions return an error on bad input.

	// TODO: Create _, e := strconv.Atoi("wat")
	// TODO: Print e
}"""
    },
    {
        "key": "url-parsing",
        "display_name": "URL Parsing",
        "template": """// URLs provide a [uniform way to locate resources](https://adam.herokuapp.com/past/2010/3/30/urls_are_the_uniform_way_to_locate_resources/).
// Here's how to parse URLs in Go.

package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {

	// We'll parse this example URL, which includes a
	// scheme, authentication info, host, port, path,
	// query params, and query fragment.

	// TODO: Create s := "postgres://user:pass@host.com:5432/path?k=v#f"

	// Parse the URL and ensure there are no errors.

	// TODO: Create u, err := url.Parse(s)
	// TODO: Print err

	// Accessing the scheme is straightforward.

	// TODO: Print u.Scheme

	// `User` contains all authentication info; call
	// `Username` and `Password` on this for individual
	// values.

	// TODO: Print u.User
	// TODO: Print u.User.Username()
	// TODO: Create p, _ := u.User.Password()
	// TODO: Print p

	// The `Host` contains both the hostname and the port,
	// if present. Use `SplitHostPort` to extract them.

	// TODO: Print u.Host
	// TODO: Create host, port, _ := net.SplitHostPort(u.Host)
	// TODO: Print host
	// TODO: Print port

	// Here we extract the `path` and the fragment after
	// the `#`.

	// TODO: Print u.Path
	// TODO: Print u.Fragment

	// To get query params in a string of `k=v` format,
	// use `RawQuery`. You can also parse query params
	// into a map. The parsed query param maps are from
	// strings to slices of strings, so index into `[0]`
	// if you only want the first value.

	// TODO: Print u.RawQuery
	// TODO: Create m, _ := url.ParseQuery(u.RawQuery)
	// TODO: Print m
	// TODO: Print m["k"][0]
}"""
    },
    {
        "key": "sha256-hashes",
        "display_name": "SHA256 Hashes",
        "template": """// [_SHA256 hashes_](https://en.wikipedia.org/wiki/SHA-2) are
// frequently used to compute short identities for binary
// or text blobs. For example, TLS/SSL certificates use SHA256
// to compute a certificate's signature. Here's how to compute
// SHA256 hashes in Go.

package main

// Go implements several hash functions in various
// `crypto/*` packages.
import (
	"crypto/sha256"
	"fmt"
)

func main() {
	// TODO: Create s := "sha256 this string"

	// Here we start with a new hash.
	// TODO: Create h := sha256.New()

	// `Write` expects bytes. If you have a string `s`,
	// use `[]byte(s)` to coerce it to bytes.

	// TODO: h.Write([]byte(s))

	// This gets the finalized hash result as a byte
	// slice. The argument to `Sum` can be used to append
	// to an existing byte slice: it usually isn't needed.

	// TODO: Create bs := h.Sum(nil)

	// TODO: Print s
	// TODO: Print fmt.Printf("%x\n", bs)
}"""
    },
    {
        "key": "base64-encoding",
        "display_name": "Base64 Encoding",
        "template": """// Go provides built-in support for [base64
// encoding/decoding](https://en.wikipedia.org/wiki/Base64).

package main

// This syntax imports the `encoding/base64` package with
// the `b64` name instead of the default `base64`. It'll
// save us some space below.
import (
	b64 "encoding/base64"
	"fmt"
)

func main() {

	// Here's the `string` we'll encode/decode.
	// TODO: Create data := "abc123!?$*&()'-=@~"

	// Go supports both standard and URL-compatible
	// base64. Here's how to encode using the standard
	// encoder. The encoder requires a `[]byte` so we
	// convert our `string` to that type.

	// TODO: Create sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	// TODO: Print sEnc

	// Decoding may return an error, which you can check
	// if you don't already know the input to be
	// well-formed.

	// TODO: Create sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	// TODO: Print sDec
	// TODO: Print fmt.Println()

	// This encodes/decodes using a URL-compatible base64
	// format.

	// TODO: Create uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	// TODO: Print uEnc
	// TODO: Create uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	// TODO: Print uDec
	// TODO: Print fmt.Println()
}"""
    },
    {
        "key": "reading-files",
        "display_name": "Reading Files",
        "template": """// Reading and writing files are basic tasks needed for
// many Go programs. First we'll look at some examples of
// reading files.

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// Reading files requires checking most calls for errors.
// This helper will streamline our error checks below.

// TODO: Create function check(e error) that checks if e is not nil, panic with e


func main() {

	// Perhaps the most basic file reading task is
	// slurping a file's entire contents into memory.

	// TODO: Create dat, err := os.ReadFile("/tmp/dat")
	// TODO: Print err
	// TODO: Print fmt.Print(string(dat))

	// You'll often want more control over how and what
	// parts of a file are read. For these tasks, start
	// by `Open`ing a file to obtain an `os.File` value.

	// TODO: Create f, err := os.Open("/tmp/dat")
	// TODO: Print err

	// Read some bytes from the beginning of the file.
	// Allow up to 5 to be read but also note how many
	// actually were read.

	// TODO: Create b1 := make([]byte, 5)
	// TODO: Create n1, err := f.Read(b1)
	// TODO: Print err
	// TODO: Print fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

	// You can also `Seek` to a known location in the file
	// and `Read` from there.

	// TODO: Create o2, err := f.Seek(6, io.SeekStart)
	// TODO: Print err
	// TODO: Create b2 := make([]byte, 2)
	// TODO: Create n2, err := f.Read(b2)
	// TODO: Print err
	// TODO: Print fmt.Printf("%d bytes @ %d: ", n2, o2)
	// TODO: Print fmt.Printf("%v\n", string(b2[:n2]))

	// Other methods of seeking are relative to the
	// current cursor position,

	// TODO: Create _, err = f.Seek(2, io.SeekCurrent)
	// TODO: Print err

	// and relative to the end of the file.

	// TODO: Create _, err = f.Seek(-4, io.SeekEnd)
	// TODO: Print err

	// The `io` package provides some functions that may
	// be helpful for file reading. For example, reads
	// like the ones above can be more robustly
	// implemented with `ReadAtLeast`.

	// TODO: Create o3, err := f.Seek(6, io.SeekStart)
	// TODO: Print err
	// TODO: Create b3 := make([]byte, 2)
	// TODO: Create n3, err := io.ReadAtLeast(f, b3, 2)
	// TODO: Print err
	// TODO: Print fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	// There is no built-in rewind, but
	// `Seek(0, io.SeekStart)` accomplishes this.

	// TODO: Create _, err = f.Seek(0, io.SeekStart)
	// TODO: Print err

	// The `bufio` package implements a buffered
	// reader that may be useful both for its efficiency
	// with many small reads and because of the additional
	// reading methods it provides.

	// TODO: Create r4 := bufio.NewReader(f)
	// TODO: Create b4, err := r4.Peek(5)
	// TODO: Print err
	// TODO: Print fmt.Printf("5 bytes: %s\n", string(b4))

	// Close the file when you're done (usually this would
	// be scheduled immediately after `Open`ing with
	// `defer`).

	// TODO: Create f.Close()
}"""
    },
    {
        "key": "writing-files",
        "display_name": "Writing Files",
        "template": """// Writing files in Go follows similar patterns to the
// ones we saw earlier for reading.

package main

import (
	"bufio"
	"fmt"
	"os"
)

// TODO: Create function check(e error) that checks if e is not nil, panic with e

func main() {

	// To start, here's how to dump a string (or just
	// bytes) into a file.

	// TODO: Create d1 := []byte("hello\ngo\n")
	// TODO: Create err := os.WriteFile("/tmp/dat1", d1, 0644)
	// TODO: Print err

	// For more granular writes, open a file for writing.

	// TODO: Create f, err := os.Create("/tmp/dat2")
	// TODO: Print err

	// It's idiomatic to defer a `Close` immediately
	// after opening a file.

	// TODO: Defer the closing of the file with defer f.Close()

	// You can `Write` byte slices as you'd expect.

	// TODO: Create d2 := []byte{115, 111, 109, 101, 10}
	// TODO: Create n2, err := f.Write(d2)
	// TODO: Print err
	// TODO: Print fmt.Printf("wrote %d bytes\n", n2)

	// A `WriteString` is also available.

	// TODO: Create n3, err := f.WriteString("writes\n")
	// TODO: Print err
	// TODO: Print fmt.Printf("wrote %d bytes\n", n3)

	// Issue a `Sync` to flush writes to stable storage.

	// TODO: Create f.Sync()

	// `bufio` provides buffered writers in addition
	// to the buffered readers we saw earlier.

	// TODO: Create w := bufio.NewWriter(f)
	// TODO: Print err
	// TODO: Print fmt.Printf("wrote %d bytes\n", n4)

	// Use `Flush` to ensure all buffered operations have
	// been applied to the underlying writer.

	// TODO: Create w.Flush()

}"""
    },
    {
        "key": "line-filters",
        "display_name": "Line Filters",
        "template": """// A _line filter_ is a common type of program that reads
// input on stdin, processes it, and then prints some
// derived result to stdout. `grep` and `sed` are common
// line filters.

// Here's an example line filter in Go that writes a
// capitalized version of all input text. You can use this
// pattern to write your own Go line filters.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// Wrapping the unbuffered `os.Stdin` with a buffered
	// scanner gives us a convenient `Scan` method that
	// advances the scanner to the next token; which is
	// the next line in the default scanner.

	// TODO: Create scanner := bufio.NewScanner(os.Stdin)

	// TODO: For scanner.Scan(), use strings.ToUpper(scanner.Text()) to uppercase the line
	// TODO: Write out the uppercased line


	// Check for errors during `Scan`. End of file is
	// expected and not reported by `Scan` as an error.

	// TODO: Create err := scanner.Err()
	// TODO: If err is not nil, print "error:", err and exit with 1
}"""
    },
    {
        "key": "file-paths",
        "display_name": "File Paths",
        "template": """// The `filepath` package provides functions to parse
// and construct *file paths* in a way that is portable
// between operating systems; `dir/file` on Linux vs.
// `dir\file` on Windows, for example.
package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {

	// `Join` should be used to construct paths in a
	// portable way. It takes any number of arguments
	// and constructs a hierarchical path from them.

	// TODO: Create p := filepath.Join("dir1", "dir2", "filename")
	// TODO: Print p

	// You should always use `Join` instead of
	// concatenating `/`s or `\`s manually. In addition
	// to providing portability, `Join` will also
	// normalize paths by removing superfluous separators
	// and directory changes.

	// TODO: Create fmt.Println(filepath.Join("dir1//", "filename"))
	// TODO: Create fmt.Println(filepath.Join("dir1/../dir1", "filename"))

	// `Dir` and `Base` can be used to split a path to the
	// directory and the file. Alternatively, `Split` will
	// return both in the same call.

	// TODO: Create fmt.Println("Dir(p):", filepath.Dir(p))
	// TODO: Create fmt.Println("Base(p):", filepath.Base(p))

	// We can check whether a path is absolute.

	// TODO: Create fmt.Println(filepath.IsAbs("dir/file"))
	// TODO: Create fmt.Println(filepath.IsAbs("/dir/file"))

	// TODO: Create filename := "config.json"
	// TODO: Print filename


	// Some file names have extensions following a dot. We
	// can split the extension out of such names with `Ext`.


	// TODO: Create ext := filepath.Ext(filename)
	// TODO: Print ext

	// To find the file's name with the extension removed,
	// use `strings.TrimSuffix`.

	// TODO: Create fmt.Println(strings.TrimSuffix(filename, ext))

	// `Rel` finds a relative path between a *base* and a
	// *target*. It returns an error if the target cannot
	// be made relative to base.

	// TODO: Create rel, err := filepath.Rel("a/b", "a/b/t/file")
	// TODO: Print err
	// TODO: Create rel, err = filepath.Rel("a/b", "a/c/t/file")
	// TODO: Print err
	// TODO: Print rel

}"""
    },
    {
        "key": "directories",
        "display_name": "Directories",
        "template": """// Go has several useful functions for working with
// *directories* in the file system.

package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

// TODO: Create function check(e error) that checks if e is not nil, panic with e

func main() {

	// Create a new sub-directory in the current working
	// directory.

	// TODO: Create err := os.Mkdir("subdir", 0755)
	// TODO: Print err

	// When creating temporary directories, it's good
	// practice to `defer` their removal. `os.RemoveAll`
	// will delete a whole directory tree (similarly to
	// `rm -rf`).

	// TODO: Defer the removal of the directory with defer os.RemoveAll("subdir")

	// Helper function to create a new empty file.

	// TODO: Create createEmptyFile := func(name string) {
	// That gets d := []byte("")
	// and calls check(os.WriteFile(name, d, 0644))


	// TODO: call createEmptyFile("subdir/file1")

	// We can create a hierarchy of directories, including
	// parents with `MkdirAll`. This is similar to the
	// command-line `mkdir -p`.


	// TODO: Create err = os.MkdirAll("subdir/parent/child", 0755)
	// check(err)
	// TODO: Print err
	// TODO: call createEmptyFile("subdir/parent/file2")
	// TODO: call createEmptyFile("subdir/parent/file3")
	// TODO: call createEmptyFile("subdir/parent/child/file4")

	// `ReadDir` lists directory contents, returning a
	// slice of `os.DirEntry` objects.

	// TODO: Create c, err := os.ReadDir("subdir/parent")
	// TODO: Print err

	// TODO: For _, entry := range c, print " ", entry.Name(), entry.IsDir()

	// `Chdir` lets us change the current working directory,
	// similarly to `cd`.

	// TODO: Create err = os.Chdir("subdir/parent/child")
	// TODO: Print err

	// Now we'll see the contents of `subdir/parent/child`
	// when listing the *current* directory.

	// TODO: Create c, err := os.ReadDir(".")
	// TODO: Print err

	// TODO: For _, entry := range c, print " ", entry.Name(), entry.IsDir()


	// `cd` back to where we started.

	// TODO: err = os.Chdir("../../..")

	// We can also visit a directory *recursively*,
	// including all its sub-directories. `WalkDir` accepts
	// a callback function to handle every file or
	// directory visited.
	
	// TODO: err = filepath.WalkDir("subdir", visit)
}

// `visit` is called for every file or directory found
// recursively by `filepath.WalkDir`.

// TODO: Create function visit(path string, d fs.DirEntry, err error) error"""
    },
    {
        "key": "temporary-files-and-directories",
        "display_name": "Temporary Files and Directories",
        "template": """// Throughout program execution, we often want to create
// data that isn't needed after the program exits.
// *Temporary files and directories* are useful for this
// purpose since they don't pollute the file system over
// time.

package main

import (
	"fmt"
	"os"
	"path/filepath"
)


// TODO: Create function check(e error) that checks if e is not nil, panic with e

func main() {

	// The easiest way to create a temporary file is by
	// calling `os.CreateTemp`. It creates a file *and*
	// opens it for reading and writing. We provide `""`
	// as the first argument, so `os.CreateTemp` will
	// create the file in the default location for our OS.


	// TODO: Create f, err := os.CreateTemp("", "sample")
	// TODO: Print err

	// Display the name of the temporary file. On
	// Unix-based OSes the directory will likely be `/tmp`.
	// The file name starts with the prefix given as the
	// second argument to `os.CreateTemp` and the rest
	// is chosen automatically to ensure that concurrent
	// calls will always create different file names.

	// TODO: Print fmt.Println("Temp file name:", f.Name())

	// Clean up the file after we're done. The OS is
	// likely to clean up temporary files by itself after
	// some time, but it's good practice to do this
	// explicitly.

	// TODO: Defer the removal of the file with defer os.Remove(f.Name())

	// We can write some data to the file.

	// TODO: Create _, err = f.Write([]byte{1, 2, 3, 4})
	// TODO: Print err

	// If we intend to write many temporary files, we may
	// prefer to create a temporary *directory*.
	// `os.MkdirTemp`'s arguments are the same as
	// `CreateTemp`'s, but it returns a directory *name*
	// rather than an open file.

	// TODO: Create dname, err := os.MkdirTemp("", "sampledir")
	// TODO: Print err

	// TODO: Defer the removal of the directory with defer os.RemoveAll(dname)

	// Now we can synthesize temporary file names by
	// prefixing them with our temporary directory.

	// TODO: Create fname := filepath.Join(dname, "file1")
	// TODO: Create err = os.WriteFile(fname, []byte{1, 2}, 0666)
	// TODO: Print err
}"""
    },
    {
        "key": "embed-directive",
        "display_name": "Embed Directive",
        "template": """// `//go:embed` is a [compiler
// directive](https://pkg.go.dev/cmd/compile#hdr-Compiler_Directives) that
// allows programs to include arbitrary files and folders in the Go binary at
// build time. Read more about the embed directive
// [here](https://pkg.go.dev/embed).
package main

// Import the `embed` package; if you don't use any exported
// identifiers from this package, you can do a blank import with `_ "embed"`.
import (
	"embed"
)

// `embed` directives accept paths relative to the directory containing the
// Go source file. This directive embeds the contents of the file into the
// `string` variable immediately following it.
//
//go:embed folder/single_file.txt
var fileString string

// Or embed the contents of the file into a `[]byte`.
//
//go:embed folder/single_file.txt
var fileByte []byte

// We can also embed multiple files or even folders with wildcards. This uses
// a variable of the [embed.FS type](https://pkg.go.dev/embed#FS), which
// implements a simple virtual file system.
//
//go:embed folder/single_file.txt
//go:embed folder/*.hash
var folder embed.FS

func main() {

	// Print out the contents of `single_file.txt`.
	// TODO: Print fileString
	// TODO: Print string(fileByte)

	// Retrieve some files from the embedded folder.
	// TODO: Create content1, _ := folder.ReadFile("folder/file1.hash")
	// TODO: Print string(content1)
	// TODO: Create content2, _ := folder.ReadFile("folder/file2.hash")
	// TODO: Print string(content2)
}"""
    },
    {
        "key": "testing-and-benchmarking",
        "display_name": "Testing and Benchmarking",
        "template": """// Unit testing is an important part of writing
// principled Go programs. The `testing` package
// provides the tools we need to write unit tests
// and the `go test` command runs tests.

// For the sake of demonstration, this code is in package
// `main`, but it could be any package. Testing code
// typically lives in the same package as the code it tests.
package main

import (
	"fmt"
	"testing"
)

// We'll be testing this simple implementation of an
// integer minimum. Typically, the code we're testing
// would be in a source file named something like
// `intutils.go`, and the test file for it would then
// be named `intutils_test.go`.

// TODO: Create function IntMin(a, b int) int that returns the minimum of a and b

// A test is created by writing a function with a name
// beginning with `Test`.

// TODO: Create function TestIntMinBasic(t *testing.T) that tests IntMin(2, -2) = -2

// `t.Error*` will report test failures but continue
// executing the test. `t.Fatal*` will report test
// failures and stop the test immediately.

func TestIntMinBasic(t *testing.T) {
	// TODO: Create ans := IntMin(2, -2)
	// TODO: If ans != -2, t.Errorf("IntMin(2, -2) = %d; want -2", ans)
}

// Writing tests can be repetitive, so it's idiomatic to
// use a *table-driven style*, where test inputs and
// expected outputs are listed in a table and a single loop
// walks over them and performs the test logic.
func TestIntMinTableDriven(t *testing.T) {

	// TODO: Create tests = []struct {
	// a, b int
	// want int
	//}{
	// {0, 1, 0},
	// {1, 0, 0},
	// {2, -2, -2},
	// {0, -1, -1},
	// {-1, 0, -1},
	//}

	// TODO: For _, tt := range tests, create testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
	// TODO: t.Run(testname, func(t *testing.T) {
	// TODO: Create ans := IntMin(tt.a, tt.b)
	// TODO: If ans != tt.want, t.Errorf("got %d, want %d", ans, tt.want)
	// TODO: })
}

// Benchmark tests typically go in `_test.go` files and are
// named beginning with `Benchmark`.
// Any code that's required for the benchmark to run but should
// not be measured goes before this loop.

// TODO: Create function BenchmarkIntMin(b *testing.B) that benchmarks IntMin(1, 2)
func BenchmarkIntMin(b *testing.B) {
	for b.Loop() {
		// TODO: IntMin(1, 2)
	}
}"""
    },
    {
        "key": "command-line-arguments",
        "display_name": "Command-Line Arguments",
        "template": """// [_Command-line arguments_](https://en.wikipedia.org/wiki/Command-line_interface#Arguments)
// are a common way to parameterize execution of programs.
// For example, `go run hello.go` uses `run` and
// `hello.go` arguments to the `go` program.

package main

import (
	"fmt"
	"os"
)

func main() {

	// `os.Args` provides access to raw command-line
	// arguments. Note that the first value in this slice
	// is the path to the program, and `os.Args[1:]`
	// holds the arguments to the program.

	// TODO: Create argsWithProg := os.Args
	// TODO: Create argsWithoutProg := os.Args[1:]

	// You can get individual args with normal indexing.

	// TODO: Create arg := os.Args[3]	
	// TODO: Print arg

	// TODO: Print argsWithProg
	// TODO: Print argsWithoutProg
	// TODO: Print arg
}"""
    },
    {
        "key": "command-line-flags",
        "display_name": "Command-Line Flags",
        "template": """// [_Command-line flags_](https://en.wikipedia.org/wiki/Command-line_interface#Command-line_option)
// are a common way to specify options for command-line
// programs. For example, in `wc -l` the `-l` is a
// command-line flag.

package main

// Go provides a `flag` package supporting basic
// command-line flag parsing. We'll use this package to
// implement our example command-line program.
import (
	"flag"
	"fmt"
)

func main() {

	// Basic flag declarations are available for string,
	// integer, and boolean options. Here we declare a
	// string flag `word` with a default value `"foo"`
	// and a short description. This `flag.String` function
	// returns a string pointer (not a string value);
	// we'll see how to use this pointer below.

	// TODO: Create wordPtr := flag.String("word", "foo", "a string")

	// This declares `numb` and `fork` flags, using a
	// similar approach to the `word` flag.

	// TODO: Create numbPtr := flag.Int("numb", 42, "an int")
	// TODO: Create forkPtr := flag.Bool("fork", false, "a bool")

	// It's also possible to declare an option that uses an
	// existing var declared elsewhere in the program.
	// Note that we need to pass in a pointer to the flag
	// declaration function.

	// TODO: Create var svar string
	// TODO: Create flag.StringVar(&svar, "svar", "bar", "a string var")

	// Once all flags are declared, call `flag.Parse()`
	// to execute the command-line parsing.

	// TODO: Create flag.Parse()

	// Here we'll just dump out the parsed options and
	// any trailing positional arguments. Note that we
	// need to dereference the pointers with e.g. `*wordPtr`
	// to get the actual option values.

	// TODO: Print word:", *wordPtr)
	// TODO: Print numb:", *numbPtr)
	// TODO: Print fork:", *forkPtr)
	// TODO: Print svar:", svar)
	// TODO: Print tail:", flag.Args())
}"""
    },
    {
        "key": "command-line-subcommands",
        "display_name": "Command-Line Subcommands",
        "template": """// Some command-line tools, like the `go` tool or `git`
// have many *subcommands*, each with its own set of
// flags. For example, `go build` and `go get` are two
// different subcommands of the `go` tool.
// The `flag` package lets us easily define simple
// subcommands that have their own flags.

package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	// We declare a subcommand using the `NewFlagSet`
	// function, and proceed to define new flags specific
	// for this subcommand.

	// TODO: Create fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)
	// TODO: Create fooEnable := fooCmd.Bool("enable", false, "enable")
	// TODO: Create fooName := fooCmd.String("name", "", "name")

	// For a different subcommand we can define different
	// supported flags.

	// TODO: Create barCmd := flag.NewFlagSet("bar", flag.ExitOnError)
	// TODO: Create barLevel := barCmd.Int("level", 0, "level")

	// The subcommand is expected as the first argument
	// to the program.

	// TODO: If len(os.Args) < 2, fmt.Println("expected 'foo' or 'bar' subcommands")
	// TODO: os.Exit(1)

	// Check which subcommand is invoked.

	// TODO: switch os.Args[1] {
	// If the case is "foo", parse the fooCmd flags, if the case is "bar", parse the barCmd flags.
	// default print "expected 'foo' or 'bar' subcommands" and exit with 1

	// For every subcommand, we parse its own flags and
	// have access to trailing positional arguments.
	
}"""
    },
    {
        "key": "environment-variables",
        "display_name": "Environment Variables",
        "template": """// [Environment variables](https://en.wikipedia.org/wiki/Environment_variable)
// are a universal mechanism for [conveying configuration
// information to Unix programs](https://www.12factor.net/config).
// Let's look at how to set, get, and list environment variables.

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	// To set a key/value pair, use `os.Setenv`. To get a
	// value for a key, use `os.Getenv`. This will return
	// an empty string if the key isn't present in the
	// environment.

	// TODO: Create os.Setenv("FOO", "1")
	// TODO: Create fmt.Println("FOO:", os.Getenv("FOO"))
	// TODO: Create fmt.Println("BAR:", os.Getenv("BAR"))

	// Use `os.Environ` to list all key/value pairs in the
	// environment. This returns a slice of strings in the
	// form `KEY=value`. You can `strings.SplitN` them to
	// get the key and value. Here we print all the keys.
	fmt.Println()
	// TODO: iterate over os.Environ() with range and print the key
	// Use strings.SplitN(e, "=", 2) to get the key and value
}"""
    },
    {
        "key": "logging",
        "display_name": "Logging",
        "template": """// The Go standard library provides straightforward
// tools for outputting logs from Go programs, with the
// log package for free-form output and the log/slog
// package for structured output.

package main

import "fmt"

func main() {
	// TODO: Implement logging concepts
	fmt.Println("Practicing: Logging")
}"""
    },
    {
        "key": "http-client",
        "display_name": "HTTP Client",
        "template": """// The Go standard library comes with excellent support
// for HTTP clients and servers in the `net/http`
// package. In this example we'll use it to issue simple
// HTTP requests.
package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {

	// Issue an HTTP GET request to a server. `http.Get` is a
	// convenient shortcut around creating an `http.Client`
	// object and calling its `Get` method; it uses the
	// `http.DefaultClient` object which has useful default
	// settings.

	// TODO: Create resp, err := http.Get("https://gobyexample.com")
	// TODO: Print err
	

	// TODO: Defer the closing of the response body with defer resp.Body.Close()

	// Print the HTTP response status.

	// TODO: Print fmt.Println("Response status:", resp.Status)

	// Print the first 5 lines of the response body.

	// TODO: Create scanner := bufio.NewScanner(resp.Body)
	// Iterate over the scanner with for i := 0; scanner.Scan() && i < 5; i++ {
	// TODO: Print fmt.Println(scanner.Text())
	// TODO: Print err

	// Check for scanner.Err() and if it is not nil, panic with the error
}"""
    },
    {
        "key": "http-server",
        "display_name": "HTTP Server",
        "template": """// Writing a basic HTTP server is easy using the
// `net/http` package.
package main

import (
	"fmt"
	"net/http"
)

// A fundamental concept in `net/http` servers is
// *handlers*. A handler is an object implementing the
// `http.Handler` interface. A common way to write
// a handler is by using the `http.HandlerFunc` adapter
// on functions with the appropriate signature.

// TODO: Create function hello that takes a http.ResponseWriter and a http.Request
// Functions serving as handlers take a
// `http.ResponseWriter` and a `http.Request` as
// arguments. The response writer is used to fill in the
// HTTP response. Here our simple response is just
// "hello\n".


// TODO: Create function headers that takes a http.ResponseWriter and a http.Request
// Inside the function, iterate over the request headers with range and print the name and value



func main() {

	// We register our handlers on server routes using the
	// `http.HandleFunc` convenience function. It sets up
	// the *default router* in the `net/http` package and
	// takes a function as an argument.

	// TODO: Create http.HandleFunc("/hello", hello)
	// TODO: Create http.HandleFunc("/headers", headers)

	// Finally, we call the `ListenAndServe` with the port
	// and a handler. `nil` tells it to use the default
	// router we've just set up.

	// TODO: Create http.ListenAndServe(":8090", nil)
}"""
    },
    {
        "key": "context",
        "display_name": "Context",
        "template": """// In the previous example we looked at setting up a simple
// [HTTP server](http-server). HTTP servers are useful for
// demonstrating the usage of `context.Context` for
// controlling cancellation. A `Context` carries deadlines,
// cancellation signals, and other request-scoped values
// across API boundaries and goroutines.
package main

import (
	"fmt"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request) {

	// A `context.Context` is created for each request by
	// the `net/http` machinery, and is available with
	// the `Context()` method.

	// TODO: Create ctx := req.Context()
	// TODO: Print "server: hello handler started"
	// TODO: Defer "server: hello handler ended"

	// Wait for a few seconds before sending a reply to the
	// client. This could simulate some work the server is
	// doing. While working, keep an eye on the context's
	// `Done()` channel for a signal that we should cancel
	// the work and return as soon as possible.

	// TODO: Use select to wait for 10 seconds or the context's Done() channel
	// When context is <- time.After(10 * time.Second), print "hello\n"
	// When context is <- ctx.Done(), print "server: " and the context's Err()
	// and set the status code to http.StatusInternalServerError and call http.Error(w, err.Error(), internalError)
	
	}
}

func main() {

	// As before, we register our handler on the "/hello"
	// route, and start serving.

	// TODO: Create http.HandleFunc("/hello", hello)
	// TODO: Create http.ListenAndServe(":8090", nil)
}"""
    },
    {
        "key": "spawning-processes",
        "display_name": "Spawning Processes",
        "template": """// Sometimes our Go programs need to spawn other
// processes.

package main

import (
	"errors"
	"fmt"
	"io"
	"os/exec"
)

func main() {

	// We'll start with a simple command that takes no
	// arguments or input and just prints something to
	// stdout. The `exec.Command` helper creates an object
	// to represent this external process.

	// TODO: Create dateCmd := exec.Command("date")

	// The `Output` method runs the command, waits for it
	// to finish and collects its standard output.
	//  If there were no errors, `dateOut` will hold bytes
	// with the date info.

	// TODO: Create dateOut, err := dateCmd.Output()
	// TODO: Print err
	// TODO: Print dateOut
	

	// `Output` and other methods of `Command` will return
	// `*exec.Error` if there was a problem executing the
	// command (e.g. wrong path), and `*exec.ExitError`
	// if the command ran but exited with a non-zero return
	// code.


	// TODO: Create _, err = exec.Command("date", "-x").Output()
	
	// TODO: Create if err != nil {
	// Create var execErr *exec.Error and var exitErr *exec.ExitError
	// With switch, check if err is an execErr or exitErr
	// If it is, print the error
	// If it is not, panic with the error


	// Next we'll look at a slightly more involved case
	// where we pipe data to the external process on its
	// `stdin` and collect the results from its `stdout`.

	// TODO: Create grepCmd := exec.Command("grep", "hello")

	// Here we explicitly grab input/output pipes, start
	// the process, write some input to it, read the
	// resulting output, and finally wait for the process
	// to exit.

	// TODO: Create grepIn, _ := grepCmd.StdinPipe()
	// TODO: Create grepOut, _ := grepCmd.StdoutPipe()
	// Start the process, write some input to it, read the resulting output, and finally wait for the process to exit.


	// We omitted error checks in the above example, but
	// you could use the usual `if err != nil` pattern for
	// all of them. We also only collect the `StdoutPipe`
	// results, but you could collect the `StderrPipe` in
	// exactly the same way.

	// TODO: Print "> grep hello" and the result of grepBytes

	// Note that when spawning commands we need to
	// provide an explicitly delineated command and
	// argument array, vs. being able to just pass in one
	// command-line string. If you want to spawn a full
	// command with a string, you can use `bash`'s `-c`
	// option:

	// TODO: Create lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	// Call lsCmd.Output() and print the result 
}"""
    },
    {
        "key": "execing-processes",
        "display_name": "Exec'ing Processes",
        "template": """// In the previous example we looked at
// [spawning external processes](spawning-processes). We
// do this when we need an external process accessible to
// a running Go process. Sometimes we just want to
// completely replace the current Go process with another
// (perhaps non-Go) one. To do this we'll use Go's
// implementation of the classic
// <a href="https://en.wikipedia.org/wiki/Exec_(operating_system)"><code>exec</code></a>
// function.

package main

import (
	"os"
	"os/exec"
	"syscall"
)

func main() {

	// For our example we'll exec `ls`. Go requires an
	// absolute path to the binary we want to execute, so
	// we'll use `exec.LookPath` to find it (probably
	// `/bin/ls`).

	// TODO: Create binary, lookErr := exec.LookPath("ls")
	// TODO: Print lookErr

	// `Exec` requires arguments in slice form (as
	// opposed to one big string). We'll give `ls` a few
	// common arguments. Note that the first argument should
	// be the program name.

	// TODO: Create args := []string{"ls", "-a", "-l", "-h"}

	// `Exec` also needs a set of [environment variables](environment-variables)
	// to use. Here we just provide our current
	// environment.

	// TODO: Create env := os.Environ()

	// Here's the actual `syscall.Exec` call. If this call is
	// successful, the execution of our process will end
	// here and be replaced by the `/bin/ls -a -l -h`
	// process. If there is an error we'll get a return
	// value.

	// TODO: Create execErr := syscall.Exec(binary, args, env)
	// TODO: Print execErr
}"""
    },
    {
        "key": "signals",
        "display_name": "Signals",
        "template": """// Sometimes we'd like our Go programs to intelligently
// handle [Unix signals](https://en.wikipedia.org/wiki/Unix_signal).
// For example, we might want a server to gracefully
// shutdown when it receives a `SIGTERM`, or a command-line
// tool to stop processing input if it receives a `SIGINT`.
// Here's how to handle signals in Go with channels.

package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	// Go signal notification works by sending `os.Signal`
	// values on a channel. We'll create a channel to
	// receive these notifications. Note that this channel
	// should be buffered.

	// TODO: Create sigs channel of os.Signal with buffer size 1


	// `signal.Notify` registers the given channel to
	// receive notifications of the specified signals.

	// TODO: Call signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// We could receive from `sigs` here in the main
	// function, but let's see how this could also be
	// done in a separate goroutine, to demonstrate
	// a more realistic scenario of graceful shutdown.

	// TODO: Create done channel of bool with buffer size 1

	go func() {
		// This goroutine executes a blocking receive for
		// signals. When it gets one it'll print it out
		// and then notify the program that it can finish.

		// TODO: Create sig := <-sigs
		// TODO: Print sig
		// TODO: Send true to done channel
	}()

	// The program will wait here until it gets the
	// expected signal (as indicated by the goroutine
	// above sending a value on `done`) and then exit.

	// TODO: Print "awaiting signal"
	// TODO: Receive from done channel
	// TODO: Print "exiting"
}"""
    },
    {
        "key": "exit",
        "display_name": "Exit",
        "template": """// Use `os.Exit` to immediately exit with a given
// status.

package main

import (
	"fmt"
	"os"
)

func main() {

	// `defer`s will _not_ be run when using `os.Exit`, so
	// this `fmt.Println` will never be called.

	// TODO: Defer fmt.Println("!")

	// Exit with status 3.

	// TODO: Create os.Exit(3)
}

// Note that unlike e.g. C, Go does not use an integer
// return value from `main` to indicate exit status. If
// you'd like to exit with a non-zero status you should
// use `os.Exit`."""
    }
]

def topic_to_package_name(topic):
    """
    Convert a topic name to a valid Go package name.
    - Convert to lowercase
    - Replace spaces and special characters with underscores
    - Remove consecutive underscores
    - Ensure it starts with a letter
    """
    # Convert to lowercase and replace problematic characters
    name = topic.lower()
    name = re.sub(r'[^a-z0-9]+', '_', name)
    name = re.sub(r'_+', '_', name)  # Remove consecutive underscores
    name = name.strip('_')  # Remove leading/trailing underscores
    
    # Ensure it starts with a letter (prepend 'go_' if it starts with a number)
    if name and name[0].isdigit():
        name = 'go_' + name
    
    return name

def create_practice_module(topic_index, topic_info, base_dir):
    """Create a complete practice module for a given topic."""
    # Create numbered directory name
    dir_name = f"{topic_index:02d}{topic_info['display_name'].replace(' ', '').replace('/', '')}"
    directory = os.path.join(base_dir, dir_name)
    
    # Use the topic key for the Go file name
    package_name = topic_to_package_name(topic_info['key'])
    
    print(f"\nCreating module {topic_index:02d}: '{topic_info['display_name']}' -> {dir_name}")

    # Create directory
    os.makedirs(directory, exist_ok=True)

    # Create .go file
    create_go_file(directory, package_name, topic_info)

    # Create go.mod file
    create_go_mod(directory, package_name)

def create_go_file(directory, package_name, topic_info):
    """Create a practice template .go file based on Go by Example patterns."""
    go_filename = f"{package_name}.go"
    go_filepath = os.path.join(directory, go_filename)

    with open(go_filepath, 'w') as f:
        f.write(topic_info["template"])

    print(f"  Created {go_filename}")

def create_go_mod(directory, package_name):
    """Create a go.mod file for the module."""
    go_mod_path = os.path.join(directory, "go.mod")
    
    go_mod_content = f'''module github.com/orsenthil/gobyexample/{package_name}

go 1.25
'''
    
    with open(go_mod_path, 'w') as f:
        f.write(go_mod_content)
    
    print(f"  Created go.mod")

def create_go_workspace(base_dir):
    """Create a go.work file to manage all modules in the workspace."""
    go_work_path = os.path.join(base_dir, "go.work")

    print("\nCreating Go workspace file...")

    go_work_content = "go 1.25\n\nuse (\n"
    for i, topic_info in enumerate(PRACTICE_TEMPLATES, 1):
        dir_name = f"{i:02d}{topic_info['display_name'].replace(' ', '').replace('/', '')}"
        go_work_content += f"    ./{dir_name}\n"
    go_work_content += ")\n"

    with open(go_work_path, 'w') as f:
        f.write(go_work_content)

    print("  Created go.work (enables multi-module workspace)")

def clean_modules(base_dir):
    """Remove all practice modules and the go.work file."""
    print(f"🧹 Cleaning up Go practice modules in: {base_dir}")

    removed_count = 0

    # Remove all module directories
    for i, topic_info in enumerate(PRACTICE_TEMPLATES, 1):
        dir_name = f"{i:02d}{topic_info['display_name'].replace(' ', '').replace('/', '')}"
        directory = os.path.join(base_dir, dir_name)

        if os.path.exists(directory):
            try:
                shutil.rmtree(directory)
                print(f"  Removed {dir_name}/")
                removed_count += 1
            except Exception as e:
                print(f"  ❌ Failed to remove {dir_name}/: {e}")
        else:
            print(f"  ⏭️  {dir_name}/ doesn't exist")

    # Remove go.work file
    go_work_path = os.path.join(base_dir, "go.work")
    if os.path.exists(go_work_path):
        try:
            os.remove(go_work_path)
            print("  Removed go.work")
        except Exception as e:
            print(f"  ❌ Failed to remove go.work: {e}")
    else:
        print("  ⏭️  go.work doesn't exist")

    print(f"\n✅ Cleanup complete! Removed {removed_count} modules.")

def main():
    """Main function to handle command line arguments and operations."""
    parser = argparse.ArgumentParser(
        description="Set up or clean up Go practice modules for learning",
        formatter_class=argparse.RawDescriptionHelpFormatter,
        epilog="""
Examples:
  # Create all practice modules
  python3 setup_go_practice.py
  
  # Clean up all modules for fresh start
  python3 setup_go_practice.py --clean
  
  # Create modules (same as no arguments)
  python3 setup_go_practice.py --create
        """
    )
    
    group = parser.add_mutually_exclusive_group()
    group.add_argument(
        '--clean', '-c',
        action='store_true',
        help='Remove all practice modules and go.work file'
    )
    group.add_argument(
        '--create',
        action='store_true',
        help='Create practice modules (default action)'
    )
    
    args = parser.parse_args()
    script_dir = os.path.dirname(os.path.abspath(__file__))

    if args.clean:
        # Clean up modules
        clean_modules(script_dir)
    else:
        # Create modules (default behavior)
        print(f"🚀 Setting up Go practice modules in: {script_dir}")
        print(f"Total modules to create: {len(PRACTICE_TEMPLATES)}")

        # Create all modules
        for i, topic_info in enumerate(PRACTICE_TEMPLATES, 1):
            create_practice_module(i, topic_info, script_dir)

        # Create Go workspace file
        create_go_workspace(script_dir)

        print(f"\n✅ Successfully created {len(PRACTICE_TEMPLATES)} Go practice modules!")
        print("\nTo run a specific module:")
        print("  From terminal: cd <module_directory> && go run <module_name>.go")
        print("  From editor: Open any .go file and use the Run/Debug buttons")
        print("\nThe go.work file enables multi-module support in your editor.")
        print("Happy coding! 🚀")
        print(f"\n💡 Tip: Use '{sys.argv[0]} --clean' to remove all modules for fresh practice")
        print("\n📚 Practice templates are based on Go by Example with implementation removed.")
        print("   Follow the comment instructions to implement each concept from scratch.")
        print("   This hands-on approach will help you learn Go programming effectively!")

if __name__ == "__main__":
    main()
