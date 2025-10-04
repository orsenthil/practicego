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
// Declare constant s with value "constant"

func main() {
	fmt.Println(s)

	// A `const` statement can appear anywhere a `var`
	// statement can.
	// Declare constant n with value 500000000

	// Constant expressions perform arithmetic with
	// arbitrary precision.
	// Declare constant d as 3e20 / n
	fmt.Println(d)

	// A numeric constant has no type until it's given
	// one, such as by an explicit conversion.
	// Print d converted to int64

	// A number can be given a type by using it in a
	// context that requires one, such as a variable
	// assignment or function call. For example, here
	// `math.Sin` expects a `float64`.
	// Print the result of math.Sin(n)
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
	// Initialize i := 1
	// Create a for loop that runs while i <= 3
	// In each iteration: print i, then increment i

	// A classic initial/condition/after `for` loop.
	// Create a for loop with j := 0; j < 3; j++
	// Print j in each iteration

	// Another way of accomplishing the basic "do this
	// N times" iteration is `range` over an integer.
	// Use range 3 to iterate, printing "range" and the index

	// `for` without a condition will loop repeatedly
	// until you `break` out of the loop or `return` from
	// the enclosing function.
	// Create an infinite for loop that prints "loop" then breaks

	// You can also `continue` to the next iteration of
	// the loop.
	// Use range 6 to iterate through numbers 0-5
	// If the number is even, continue to next iteration
	// Otherwise print the number
}"""
    },
    {
        "key": "if-else",
        "display_name": "If/Else", 
        "template": """// Branching with `if` and `else` in Go is
// straight-forward.

package main

import "fmt"

func main() {

	// Write a basic if/else statement
	// Check if 7%2 == 0, print "7 is even" or "7 is odd"

	// Write an `if` statement without an else.
	// Check if 8%4 == 0, print "8 is divisible by 4"

	// Logical operators like `&&` and `||` are often
	// useful in conditions.
	// Check if 7%2 == 0 || 8%4 == 0, print "either 7 or 8 are even"

	// A statement can precede conditionals; any variables
	// declared in this statement are available in the current
	// and all subsequent branches.
	// Use if num := 9; num < 0 to check and print if negative

	// Note that you don't need parentheses around conditions
	// in Go, but that the braces are required.
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

	// Write a basic `switch` statement
	// Set i := 2, switch on i with cases for 1, 2, 3
	fmt.Print("Write ", i, " as ")

	// You can use commas to separate multiple expressions
	// in the same `case` statement. We use the optional
	// `default` case in this example as well.
	// Switch on time.Now().Weekday() with cases for Saturday/Sunday and default

	// `switch` without an expression is an alternate way
	// to express if/else logic. Here we also show how the
	// `case` expressions can be non-constants.
	// Create a switch with no expression, check time conditions
	t := time.Now()

	// A type `switch` compares types instead of values. You
	// can use this to discover the type of an interface
	// value. In this example, the variable `t` will have the
	// type corresponding to its clause.
	// Create a function that uses type switch on interface{}
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

	// Create an array `a` that will hold exactly 5 `int`s.
	// The type of elements and length are both part of the 
	// array's type. By default an array is zero-valued, 
	// which for `int`s means `0`s.

	fmt.Println("emp:", a)

	// Set a value at an index using the `array[index] = value` 
	// syntax, and get a value with `array[index]`.
	// Set a[4] to 100

	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	// Use the builtin `len` to get the length of an array.

	fmt.Println("len:", len(a))

	// Declare and initialize an array in one line.
	// Create array b with values [1, 2, 3, 4, 5]

	fmt.Println("dcl:", b)

	// Have the compiler count the number of elements for you with `...`
	// Create array b using [...] syntax with values [1, 2, 3, 4, 5]

	fmt.Println("dcl:", b)

	// Specify the index with `:`, the elements in between will be zeroed.
	// Create array c using [...] with 100 at index 0, 400 at index 3, and 500 at index 4

	fmt.Println("idx:", c)

	// Compose array types to build multi-dimensional data structures.
	// Declare a two-dimensional array twoD of size [2][3]int

	// Use nested loops (range 2, range 3) to populate twoD[i][j] = i + j

	fmt.Println("2d: ", twoD)

	// Create and initialize multi-dimensional arrays at once too.
	// Create and initialize twoD2 with {{1, 2, 3}, {1, 2, 3}}

	fmt.Println("2d: ", twoD2)
}"""
    },
    {
        "key": "slices",
        "display_name": "Slices",
        "template": """// Slices are an important data type in Go, giving a more
// powerful interface to sequences than arrays.

package main

import "fmt"

func main() {

	// Unlike arrays, slices are typed only by the elements
	// they contain (not the number of elements). An uninitialized
	// slice equals to nil and has length 0.
	// Declare slice s of strings

	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	// To create an empty slice with non-zero length, use
	// the builtin `make`. Here we make a slice of
	// `string`s of length `3` (initially zero-valued).
	// Create slice s with make, length 3

	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	// We can set and get just like with arrays.
	// Set s[0] = "a", s[1] = "b", s[2] = "c"

	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	// `len` returns the length of the slice as expected.
	fmt.Println("len:", len(s))

	// In addition to these basic operations, slices
	// support several more that make them richer than
	// arrays. One is the builtin `append`, which
	// returns a slice containing one or more new values.
	// Note that we need to accept a return value from
	// `append` as we may get a new slice value.
	// Append "d" to s, then append "e" and "f"

	fmt.Println("apd:", s)

	// Slices can also be `copy`'d. Here we create an
	// empty slice `c` of the same length as `s` and copy
	// into `c` from `s`.
	// Create slice c with make, same length as s
	// Copy s into c

	fmt.Println("cpy:", c)

	// Slices support a "slice" operator with the syntax
	// `slice[low:high]`. For example, this gets a slice
	// of the elements `s[2]`, `s[3]`, and `s[4]`.
	// Create slice l := s[2:5]

	fmt.Println("sl1:", l)

	// This slices up to (but excluding) `s[5]`.
	// Create slice l := s[:5]

	fmt.Println("sl2:", l)

	// And this slices up from (and including) `s[2]`.
	// Create slice l := s[2:]

	fmt.Println("sl3:", l)

	// We can declare and initialize a variable for slice
	// in a single line as well.
	// Create slice t := []string{"g", "h", "i"}

	fmt.Println("dcl:", t)

	// The builtin copy function copies elements between slices.
	// If the slices have different lengths, copy will use the
	// smaller of the two slice lengths.
	// Create slice c2 := make([]string, len(t))
	// Copy t to c2

	fmt.Println("cpy:", c2)

	// Slices can be composed into multi-dimensional data
	// structures. The length of the inner slices can
	// vary, unlike with multi-dimensional arrays.
	// Create 2D slice twoD := make([][]int, 3)
	// Use loop to populate each inner slice with different lengths

	fmt.Println("2d: ", twoD)
}"""
    },
    {
        "key": "maps",
        "display_name": "Maps",
        "template": """// _Maps_ are Go's built-in associative data type
// (sometimes called _hashes_ or _dicts_ in other languages).

package main

import "fmt"

func main() {

	// To create an empty map, use the builtin `make`:
	// `make(map[key-type]val-type)`.
	// Create map m := make(map[string]int)

	// Set key/value pairs using typical `name[key] = val`
	// syntax.
	// Set m["k1"] = 7 and m["k2"] = 13

	// Printing a map with e.g. `fmt.Println` will show all of
	// its key/value pairs.
	fmt.Println("map:", m)

	// Get a value for a key with `name[key]`.
	// Get and print value for key "k1"

	// If the key doesn't exist, the zero value of the
	// value type is returned.
	// Get and print value for key "k3"

	// The builtin `len` returns the number of key/value
	// pairs when called on a map.
	fmt.Println("len:", len(m))

	// The builtin `delete` removes key/value pairs from
	// a map.
	// Delete key "k2" from map m

	// Clear removes all key/value pairs from a map.
	// Clear map m (Go 1.21+)

	// The optional second return value when getting a
	// value from a map indicates if the key was present
	// in the map. This can be used to disambiguate
	// between missing keys and keys with zero values
	// like `0` or `""`. Here we didn't need the value
	// itself, so we ignored it with the _blank identifier_ `_`.
	// Check if key "k2" exists in map m

	// You can also declare and initialize a new map in
	// the same line with this syntax.
	// Create map n := map[string]int{"foo": 1, "bar": 2}

	fmt.Println("map:", n)

	// The `maps` package contains a number of useful
	// utility functions for maps.
	// Use maps.Equal to compare maps m and n (need to import maps)
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
// Create function plus(a, b int) int that returns a + b

// When you have multiple consecutive parameters of
// the same type, you may omit the type name for the
// like-typed parameters up to the final parameter that
// declares the type.
// Create function plusPlus(a, b, c int) int that returns a + b + c

func main() {

	// Call a function just as you'd expect, with
	// `name(args)`.
	// Call plus(1, 2) and store in res

	fmt.Println("1+2 =", res)

	// Call plusPlus(1, 2, 3) and store in res

	fmt.Println("1+2+3 =", res)
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

// The `(int, int)` in this function signature shows that
// the function returns 2 `int`s.
// Create function vals() (int, int) that returns 3, 7

func main() {

	// Here we use the 2 different return values from the
	// call with _multiple assignment_.
	// Call vals() and assign to a, b

	fmt.Println(a)
	fmt.Println(b)

	// If you only want a subset of the returned values,
	// use the blank identifier `_`.
	// Call vals() and only use the second return value

	fmt.Println(c)
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
// Create function sum(nums ...int) that calculates sum of all nums

func main() {

	// Variadic functions can be called in the usual way
	// with individual arguments.
	// Call sum(1, 2) and sum(1, 2, 3)

	// If you already have multiple args in a slice,
	// apply them to a variadic function using
	// `func(slice...)` like this.
	// Create slice nums := []int{1, 2, 3, 4}
	// Call sum(nums...)
}"""
    },
    {
        "key": "closures",
        "display_name": "Closures",
        "template": """// Go supports [_anonymous functions_](https://en.wikipedia.org/wiki/Anonymous_function),
// which can form [_closures_](https://en.wikipedia.org/wiki/Closure_(computer_science)).
// Anonymous functions are useful when you want to define
// a function inline without having to name it.

package main

import "fmt"

// This function `intSeq` returns another function, which
// we define anonymously in the body of `intSeq`. The
// returned function _closes over_ the variable `i` to
// form a closure.
// Create function intSeq() func() int
// Inside, create variable i := 0
// Return anonymous function that increments i and returns it

func main() {

	// We call `intSeq`, assigning the result (a function)
	// to `nextInt`. This function value captures its
	// own `i` value, which will be updated each time
	// we call `nextInt`.
	// Call intSeq() and assign to nextInt

	// Call nextInt a few times to see the closure in action
	// Call nextInt() multiple times and print results

	// To confirm that the state is unique to that
	// particular function, create and test a new one.
	// Create newInts := intSeq() and call it to show separate state
}"""
    },
    {
        "key": "recursion",
        "display_name": "Recursion",
        "template": """// Go supports
// [_recursive functions_](https://en.wikipedia.org/wiki/Recursion_(computer_science)).
// Here's a classic example.

package main

import "fmt"

// This `fact` function calls itself until it reaches the
// base case of `fact(0)`.
// Create recursive function fact(n int) int
// Base case: if n == 0 return 1
// Recursive case: return n * fact(n-1)

// Closures can also be recursive, but this requires the
// closure to be declared with a typed `var` explicitly
// before it's defined.
// Create variable fib of type func(int) int
// Assign anonymous function that calculates fibonacci recursively

func main() {
	// Call fact(7) and print result

	// Call fib(7) and print result
}"""
    },
    {
        "key": "range",
        "display_name": "Range over Built-in Types",
        "template": """// _Range_ iterates over elements in a variety of data
// structures. Let's see how to use `range` with some
// of the data structures we've already learned.

package main

import "fmt"

func main() {

	// Here we use `range` to sum the numbers in a slice.
	// Arrays work like this too.
	// Create slice nums := []int{2, 3, 4}
	// Use range to sum all numbers in nums

	fmt.Println("sum:", sum)

	// `range` on arrays and slices provides both the
	// index and value for each entry. Above we didn't
	// need the index, so we ignored it with the
	// _blank identifier_ `_`. Sometimes we actually want
	// the indexes though.
	// Use range over nums to print index and value

	// `range` without a value on slices iterates over just
	// the indexes.
	// Use range to iterate over just indices of nums

	// `range` on a map iterates over key/value pairs.
	// Create map kvs := map[string]string{"a": "apple", "b": "banana"}
	// Use range to iterate over kvs

	// `range` can also iterate over just the keys of a map.
	// Use range to iterate over just keys of kvs

	// `range` on strings iterates over Unicode code
	// points. The first value is the starting byte index
	// of the `rune` and the second the `rune` itself.
	// See [Strings and Runes](strings-and-runes) for more
	// details.
	// Use range over string "go" to print index and rune
}"""
    },
    {
        "key": "pointers",
        "display_name": "Pointers", 
        "template": """// Go supports [_pointers_](https://en.wikipedia.org/wiki/Pointer_(computer_programming)),
// allowing you to pass references to values and records
// within your program.

package main

import "fmt"

// We'll show how pointers work in contrast to values with
// 2 functions: `zeroval` and `zeroptr`. `zeroval` has an
// `int` parameter, so arguments will be passed to it by
// value. `zeroval` will get a copy of `ival` distinct
// from the one in the calling function.
// Create function zeroval(ival int) that sets ival = 0

// `zeroptr` in contrast has an `*int` parameter, meaning
// that it takes an `int` pointer. The `*iptr` code in the
// function body then _dereferences_ the pointer from its
// memory address to the current value at that address.
// Assigning a value to a dereferenced pointer changes the
// value at the referenced address.
// Create function zeroptr(iptr *int) that sets *iptr = 0

func main() {
	i := 1
	fmt.Println("initial:", i)

	// Call zeroval with i

	fmt.Println("zeroval:", i)

	// The `&i` syntax gives the memory address of `i`,
	// i.e. a pointer to `i`.
	// Call zeroptr with &i

	fmt.Println("zeroptr:", i)

	// Pointers can be printed too.
	// Print the address of i using &i
}"""
    },
    {
        "key": "strings-and-runes",
        "display_name": "Strings and Runes",
        "template": """// A Go string is a read-only slice of bytes. The language
// and the standard library treat strings specially - as
// containers of text encoded in [UTF-8](https://en.wikipedia.org/wiki/UTF-8).
// In other languages, strings are made of "characters".
// In Go, the concept of a character is called a `rune` -
// it's an integer that represents a Unicode code point.
// [This Go blog post](https://blog.golang.org/strings) is a good
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
	// Print len(s)

	// Indexing into a string produces the raw byte values at
	// each index. This loop generates the hex values of all
	// the bytes that constitute the code points in `s`.
	// Loop through range len(s) and print hex values

	// To count how many _runes_ are in a string, we can use
	// the `utf8` package. Note that the run-time of
	// `RuneCountInString` depends on the size of the string,
	// because it has to decode each UTF-8 rune sequentially.
	// Some Thai characters are represented by multiple UTF-8
	// code points, so the result of this count may be surprising.
	// Print utf8.RuneCountInString(s)

	// A `range` loop handles strings specially and decodes each
	// `rune` along with its offset in the string.
	// Use range over s to print index and rune value

	// We can achieve the same iteration by using the
	// `utf8.DecodeRuneInString` function explicitly.
	// Use utf8.DecodeRuneInString to manually decode runes

	// We can also examine individual runes by converting
	// the string to a slice of runes.
	// Print individual runes by converting s to []rune

	// This demonstrates passing a `rune` value to a function.
	// Create function examineRune(r rune) and call it with runes from s
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
