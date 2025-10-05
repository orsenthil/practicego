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
// use a _wait group_.

package main

import "fmt"

func main() {
	// TODO: Implement waitgroups concepts
	fmt.Println("Practicing: WaitGroups")
}"""
    },
    {
        "key": "rate-limiting",
        "display_name": "Rate Limiting",
        "template": """// _Rate limiting_ is an important mechanism for
// controlling resource utilization and maintaining
// quality of service.

package main

import "fmt"

func main() {
	// TODO: Implement rate limiting concepts
	fmt.Println("Practicing: Rate Limiting")
}"""
    },
    {
        "key": "atomic-counters",
        "display_name": "Atomic Counters",
        "template": """// The primary mechanism for managing state in Go is
// communication over channels. We saw this for example
// with worker pools. There are a few other options
// for managing state though. Here we'll look at using
// the sync/atomic package for _atomic counters_
// accessed by multiple goroutines.

package main

import "fmt"

func main() {
	// TODO: Implement atomic counters concepts
	fmt.Println("Practicing: Atomic Counters")
}"""
    },
    {
        "key": "mutexes",
        "display_name": "Mutexes",
        "template": """// In the previous example we saw how to manage simple
// counter state using atomic operations. For more complex
// state we can use a _mutex_ to safely access data
// across multiple goroutines.

package main

import "fmt"

func main() {
	// TODO: Implement mutexes concepts
	fmt.Println("Practicing: Mutexes")
}"""
    },
    {
        "key": "stateful-goroutines",
        "display_name": "Stateful Goroutines",
        "template": """// In the previous example we used explicit locking with
// mutexes to synchronize access to shared state across
// multiple goroutines. Another option is to use the
// built-in synchronization features of goroutines and
// channels to achieve the same result.

package main

import "fmt"

func main() {
	// TODO: Implement stateful goroutines concepts
	fmt.Println("Practicing: Stateful Goroutines")
}"""
    },
    {
        "key": "sorting",
        "display_name": "Sorting",
        "template": """// Go's sort package implements sorting for builtins
// and user-defined types.

package main

import "fmt"

func main() {
	// TODO: Implement sorting concepts
	fmt.Println("Practicing: Sorting")
}"""
    },
    {
        "key": "sorting-by-functions",
        "display_name": "Sorting by Functions",
        "template": """// Sometimes we'll want to sort a collection by something
// other than its natural order.

package main

import "fmt"

func main() {
	// TODO: Implement sorting by functions concepts
	fmt.Println("Practicing: Sorting by Functions")
}"""
    },
    {
        "key": "panic",
        "display_name": "Panic",
        "template": """// A _panic_ typically means something went unexpectedly
// wrong. Mostly we use it to fail fast on errors that
// shouldn't occur during normal operation, or that we
// aren't prepared to handle gracefully.

package main

import "fmt"

func main() {
	// TODO: Implement panic concepts
	fmt.Println("Practicing: Panic")
}"""
    },
    {
        "key": "defer",
        "display_name": "Defer",
        "template": """// _Defer_ is used to ensure that a function call is
// performed later in a program's execution, usually for
// purposes of cleanup.

package main

import "fmt"

func main() {
	// TODO: Implement defer concepts
	fmt.Println("Practicing: Defer")
}"""
    },
    {
        "key": "recover",
        "display_name": "Recover",
        "template": """// Go makes it possible to _recover_ from a panic, by
// using the recover builtin function.

package main

import "fmt"

func main() {
	// TODO: Implement recover concepts
	fmt.Println("Practicing: Recover")
}"""
    },
    {
        "key": "string-functions",
        "display_name": "String Functions",
        "template": """// The standard library's strings package provides many
// useful string-related functions.

package main

import "fmt"

func main() {
	// TODO: Implement string functions concepts
	fmt.Println("Practicing: String Functions")
}"""
    },
    {
        "key": "string-formatting",
        "display_name": "String Formatting",
        "template": """// Go offers excellent support for string formatting in
// the printf tradition.

package main

import "fmt"

func main() {
	// TODO: Implement string formatting concepts
	fmt.Println("Practicing: String Formatting")
}"""
    },
    {
        "key": "text-templates",
        "display_name": "Text Templates",
        "template": """// Go provides built-in support for creating dynamic content
// or showing customized output to the user with the text/template
// package.

package main

import "fmt"

func main() {
	// TODO: Implement text templates concepts
	fmt.Println("Practicing: Text Templates")
}"""
    },
    {
        "key": "regular-expressions",
        "display_name": "Regular Expressions",
        "template": """// Go offers built-in support for regular expressions.

package main

import "fmt"

func main() {
	// TODO: Implement regular expressions concepts
	fmt.Println("Practicing: Regular Expressions")
}"""
    },
    {
        "key": "json",
        "display_name": "JSON",
        "template": """// Go offers built-in support for JSON encoding and
// decoding, including to and from built-in and custom
// data types.

package main

import "fmt"

func main() {
	// TODO: Implement JSON concepts
	fmt.Println("Practicing: JSON")
}"""
    },
    {
        "key": "xml",
        "display_name": "XML",
        "template": """// Go offers built-in support for XML and XML-like
// formats with the encoding/xml package.

package main

import "fmt"

func main() {
	// TODO: Implement XML concepts
	fmt.Println("Practicing: XML")
}"""
    },
    {
        "key": "time",
        "display_name": "Time",
        "template": """// Go offers extensive support for times and durations.

package main

import "fmt"

func main() {
	// TODO: Implement time concepts
	fmt.Println("Practicing: Time")
}"""
    },
    {
        "key": "epoch",
        "display_name": "Epoch",
        "template": """// A common requirement in programs is getting the number
// of seconds, milliseconds, or nanoseconds since the
// Unix epoch.

package main

import "fmt"

func main() {
	// TODO: Implement epoch concepts
	fmt.Println("Practicing: Epoch")
}"""
    },
    {
        "key": "time-formatting-parsing",
        "display_name": "Time Formatting / Parsing",
        "template": """// Go supports time formatting and parsing via
// pattern-based layouts.

package main

import "fmt"

func main() {
	// TODO: Implement time formatting/parsing concepts
	fmt.Println("Practicing: Time Formatting / Parsing")
}"""
    },
    {
        "key": "random-numbers",
        "display_name": "Random Numbers",
        "template": """// Go's math/rand/v2 package provides pseudorandom number
// generation.

package main

import "fmt"

func main() {
	// TODO: Implement random numbers concepts
	fmt.Println("Practicing: Random Numbers")
}"""
    },
    {
        "key": "number-parsing",
        "display_name": "Number Parsing",
        "template": """// Parsing numbers from strings is a basic but common task
// in many programs.

package main

import "fmt"

func main() {
	// TODO: Implement number parsing concepts
	fmt.Println("Practicing: Number Parsing")
}"""
    },
    {
        "key": "url-parsing",
        "display_name": "URL Parsing",
        "template": """// URLs provide a uniform way to locate resources.

package main

import "fmt"

func main() {
	// TODO: Implement URL parsing concepts
	fmt.Println("Practicing: URL Parsing")
}"""
    },
    {
        "key": "sha256-hashes",
        "display_name": "SHA256 Hashes",
        "template": """// _SHA256 hashes_ are frequently used to compute short
// identities for binary or text blobs.

package main

import "fmt"

func main() {
	// TODO: Implement SHA256 hashes concepts
	fmt.Println("Practicing: SHA256 Hashes")
}"""
    },
    {
        "key": "base64-encoding",
        "display_name": "Base64 Encoding",
        "template": """// Go provides built-in support for base64
// encoding/decoding.

package main

import "fmt"

func main() {
	// TODO: Implement base64 encoding concepts
	fmt.Println("Practicing: Base64 Encoding")
}"""
    },
    {
        "key": "reading-files",
        "display_name": "Reading Files",
        "template": """// Reading and writing files are basic tasks needed for
// many Go programs.

package main

import "fmt"

func main() {
	// TODO: Implement reading files concepts
	fmt.Println("Practicing: Reading Files")
}"""
    },
    {
        "key": "writing-files",
        "display_name": "Writing Files",
        "template": """// Writing files in Go follows similar patterns to the
// ones we saw earlier for reading.

package main

import "fmt"

func main() {
	// TODO: Implement writing files concepts
	fmt.Println("Practicing: Writing Files")
}"""
    },
    {
        "key": "line-filters",
        "display_name": "Line Filters",
        "template": """// A _line filter_ is a common type of program that reads
// input on stdin, processes it, and prints some derived
// result to stdout.

package main

import "fmt"

func main() {
	// TODO: Implement line filters concepts
	fmt.Println("Practicing: Line Filters")
}"""
    },
    {
        "key": "file-paths",
        "display_name": "File Paths",
        "template": """// The filepath package provides functions to parse
// and construct _file paths_ in a way that is portable
// between operating systems.

package main

import "fmt"

func main() {
	// TODO: Implement file paths concepts
	fmt.Println("Practicing: File Paths")
}"""
    },
    {
        "key": "directories",
        "display_name": "Directories",
        "template": """// Go has several useful functions for working with
// _directories_ in the file system.

package main

import "fmt"

func main() {
	// TODO: Implement directories concepts
	fmt.Println("Practicing: Directories")
}"""
    },
    {
        "key": "temporary-files-and-directories",
        "display_name": "Temporary Files and Directories",
        "template": """// Throughout program execution, we often want to create
// data that isn't needed after the program exits.

package main

import "fmt"

func main() {
	// TODO: Implement temporary files and directories concepts
	fmt.Println("Practicing: Temporary Files and Directories")
}"""
    },
    {
        "key": "embed-directive",
        "display_name": "Embed Directive",
        "template": """// //go:embed is a compiler directive that allows programs
// to include arbitrary files and folders in the Go binary
// at build time.

package main

import "fmt"

func main() {
	// TODO: Implement embed directive concepts
	fmt.Println("Practicing: Embed Directive")
}"""
    },
    {
        "key": "testing-and-benchmarking",
        "display_name": "Testing and Benchmarking",
        "template": """// Unit testing is an important part of writing
// principled Go programs.

package main

import "fmt"

func main() {
	// TODO: Implement testing and benchmarking concepts
	fmt.Println("Practicing: Testing and Benchmarking")
}"""
    },
    {
        "key": "command-line-arguments",
        "display_name": "Command-Line Arguments",
        "template": """// _Command-line arguments_ are a common way to parameterize
// execution of programs.

package main

import "fmt"

func main() {
	// TODO: Implement command-line arguments concepts
	fmt.Println("Practicing: Command-Line Arguments")
}"""
    },
    {
        "key": "command-line-flags",
        "display_name": "Command-Line Flags",
        "template": """// _Command-line flags_ are a common way to specify options
// for command-line programs.

package main

import "fmt"

func main() {
	// TODO: Implement command-line flags concepts
	fmt.Println("Practicing: Command-Line Flags")
}"""
    },
    {
        "key": "command-line-subcommands",
        "display_name": "Command-Line Subcommands",
        "template": """// Some command-line tools, like the go tool or git
// have many _subcommands_, each with its own set of
// flags.

package main

import "fmt"

func main() {
	// TODO: Implement command-line subcommands concepts
	fmt.Println("Practicing: Command-Line Subcommands")
}"""
    },
    {
        "key": "environment-variables",
        "display_name": "Environment Variables",
        "template": """// _Environment variables_ are a universal mechanism for
// conveying configuration information to Unix programs.

package main

import "fmt"

func main() {
	// TODO: Implement environment variables concepts
	fmt.Println("Practicing: Environment Variables")
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
// for HTTP clients and servers in the net/http package.

package main

import "fmt"

func main() {
	// TODO: Implement HTTP client concepts
	fmt.Println("Practicing: HTTP Client")
}"""
    },
    {
        "key": "http-server",
        "display_name": "HTTP Server",
        "template": """// Writing a basic HTTP server is easy using the
// net/http package.

package main

import "fmt"

func main() {
	// TODO: Implement HTTP server concepts
	fmt.Println("Practicing: HTTP Server")
}"""
    },
    {
        "key": "context",
        "display_name": "Context",
        "template": """// In the previous example we looked at setting up a simple
// HTTP server. HTTP servers are useful for demonstrating
// the usage of context.Context for controlling cancellation.

package main

import "fmt"

func main() {
	// TODO: Implement context concepts
	fmt.Println("Practicing: Context")
}"""
    },
    {
        "key": "spawning-processes",
        "display_name": "Spawning Processes",
        "template": """// Sometimes our Go programs need to spawn other, non-Go
// processes.

package main

import "fmt"

func main() {
	// TODO: Implement spawning processes concepts
	fmt.Println("Practicing: Spawning Processes")
}"""
    },
    {
        "key": "execing-processes",
        "display_name": "Exec'ing Processes",
        "template": """// In the previous example we looked at spawning external
// processes. We do this when we need an external process
// accessible to a running Go process. Sometimes we just
// want to completely replace the current Go process with
// another (perhaps non-Go) one. To do this we'll use Go's
// implementation of the classic exec function.

package main

import "fmt"

func main() {
	// TODO: Implement exec'ing processes concepts
	fmt.Println("Practicing: Exec'ing Processes")
}"""
    },
    {
        "key": "signals",
        "display_name": "Signals",
        "template": """// Sometimes we'd like our Go programs to intelligently
// handle Unix signals.

package main

import "fmt"

func main() {
	// TODO: Implement signals concepts
	fmt.Println("Practicing: Signals")
}"""
    },
    {
        "key": "exit",
        "display_name": "Exit",
        "template": """// Use os.Exit to immediately exit with a given status.

package main

import "fmt"

func main() {
	// TODO: Implement exit concepts
	fmt.Println("Practicing: Exit")
}"""
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
