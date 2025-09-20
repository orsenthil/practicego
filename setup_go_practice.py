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
PRACTICE_TEMPLATES = {
    "hello-world": {
        "display_name": "Hello World",
        "template": """// Our first program will print the classic "hello world"
// message. Here's the full source code.
package main

import "fmt"

func main() {
	// Print "hello world" to the console
}"""
    },

    "values": {
        "display_name": "Values",
        "template": """// Go has various value types including strings,
// integers, floats, booleans, etc. Here are a few
// basic examples.

package main

import "fmt"

func main() {

	// Strings, which can be added together with `+`.
	// Print the result of concatenating "go" and "lang"

	// Integers and floats.
	// Print "1+1 =" followed by the result of 1+1
	// Print "7.0/3.0 =" followed by the result of 7.0/3.0

	// Booleans, with boolean operators as you'd expect.
	// Print the result of true && false
	// Print the result of true || false
	// Print the result of !true
}"""
    },

    "variables": {
        "display_name": "Variables",
        "template": """// In Go, _variables_ are explicitly declared and used by
// the compiler to e.g. check type-correctness of function
// calls.

package main

import "fmt"

func main() {

	// `var` declares 1 or more variables.
	var a = "initial"
	// Print the value of a

	// You can declare multiple variables at once.
	var b, c int = 1, 2
	// Print b and c

	// Go will infer the type of initialized variables.
	var d = true
	// Print d

	// Variables declared without a corresponding
	// initialization are _zero-valued_. For example, the
	// zero value for an `int` is `0`.
	var e int
	// Print e

	// The `:=` syntax is shorthand for declaring and
	// initializing a variable, e.g. for
	// `var f string = "apple"` in this case.
	// This syntax is only available inside functions.
	f := "apple"
	// Print f
}"""
    },

    "constants": {
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
	// Print the constant s

	// A `const` statement can appear anywhere a `var`
	// statement can.
	const n = 500000000

	// Constant expressions perform arithmetic with
	// arbitrary precision.
	const d = 3e20 / n
	// Print d

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

    "for": {
        "display_name": "For",
        "template": """// `for` is Go's only looping construct. Here are
// some basic types of `for` loops.

package main

import "fmt"

func main() {

	// The most basic type, with a single condition.
	i := 1
	// Create a for loop that runs while i <= 3
	// In each iteration: print i, then increment i by 1

	// A classic initial/condition/after `for` loop.
	// Create a for loop with j starting at 0, continuing while j < 3, incrementing j each time
	// Print j in each iteration

	// Another way of accomplishing the basic "do this
	// N times" iteration is `range` over an integer.
	// Use range to iterate 3 times, printing "range" followed by the index

	// `for` without a condition will loop repeatedly
	// until you `break` out of the loop or `return` from
	// the enclosing function.
	// Create an infinite for loop that prints "loop" once then breaks

	// You can also `continue` to the next iteration of
	// the loop.
	// Use range to iterate through numbers 0-5
	// If the number is even, continue to next iteration
	// Otherwise print the number
}"""
    },

    "if-else": {
        "display_name": "If/Else",
        "template": """// Branching with `if` and `else` in Go is
// straight-forward.

package main

import "fmt"

func main() {

	// Write a basic example 
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	//  Write an `if` statement without an else.


	// Logical operators like `&&` and `||` are often
	// useful in conditions.

	// A statement can precede conditionals; any variables
	// declared in this statement are available in the current
	// and all subsequent branches.

        // Note that you don't need parentheses around conditions
        // in Go, but that the braces are required.

}"""
    },

    "switch": {
        "display_name": "Switch",
        "template": """// _Switch statements_ express conditionals across many
// branches.

package main

import (
	"fmt"
	"time"
)

func main() {

	//  Write  basic `switch` statement
	i := 2
	fmt.Print("Write ", i, " as ")


	// You can use commas to separate multiple expressions
	// in the same `case` statement. We use the optional
	// `default` case in this example as well.


	// `switch` without an expression is an alternate way
	// to express if/else logic. Here we also show how the
	// `case` expressions can be non-constants.


	// A type `switch` compares types instead of values.  You
	// can use this to discover the type of an interface
	// value.  In this example, the variable `t` will have the
	// type corresponding to its clause.

}"""
    },

    "arrays": {
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
	var a [5]int
	// Print "emp:" followed by the array a

	// We can set a value at an index using the
	// `array[index] = value` syntax, and get a value with
	// `array[index]`.
	// Set a[4] to 100
	// Print "set:" followed by the array a
	// Print "get:" followed by a[4]

	// The builtin `len` returns the length of an array.
	// Print "len:" followed by the length of a

	// Use this syntax to declare and initialize an array
	// in one line.
	// Create array b with values [1, 2, 3, 4, 5]
	// Print "dcl:" followed by b

	// You can also have the compiler count the number of
	// elements for you with `...`
	// Create array b using [...] syntax with same values
	// Print "dcl:" followed by b

	// If you specify the index with `:`, the elements in
	// between will be zeroed.
	// Create array b using [...] with 100 at index 0, 400 at index 3, and 500 at index 4
	// Print "idx:" followed by b

	// Array types are one-dimensional, but you can
	// compose types to build multi-dimensional data
	// structures.
	var twoD [2][3]int
	// Use nested loops to populate twoD[i][j] = i + j
	// Print "2d: " followed by twoD

	// You can create and initialize multi-dimensional
	// arrays at once too.
	// Create and initialize twoD with {{1, 2, 3}, {1, 2, 3}}
	// Print "2d: " followed by twoD
}"""
    },

    "functions": {
        "display_name": "Functions",
        "template": """// _Functions_ are central in Go. We'll learn about
// functions with a few different examples.

package main

import "fmt"

// Here's a function that takes two `int`s and returns
// their sum as an `int`.
// Create function plus that takes two ints and returns their sum

// When you have multiple consecutive parameters of
// the same type, you may omit the type name for the
// like-typed parameters up to the final parameter that
// declares the type.
// Create function plusPlus that takes three ints and returns their sum

func main() {

	// Call a function just as you'd expect, with
	// `name(args)`.
	// Call plus(1, 2) and store result in res
	// Print "1+2 =" followed by res

	// Call plusPlus(1, 2, 3) and store result in res
	// Print "1+2+3 =" followed by res
}"""
    },

    "channels": {
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
	// Create channel messages of type string

	// _Send_ a value into a channel using the `channel <-`
	// syntax. Here we send `"ping"`  to the `messages`
	// channel we made above, from a new goroutine.
	// Launch anonymous goroutine that sends "ping" to messages channel

	// The `<-channel` syntax _receives_ a value from the
	// channel. Here we'll receive the `"ping"` message
	// we sent above and print it out.
	// Receive message from messages channel and store in msg
	// Print msg
}"""
    },

    "structs": {
        "display_name": "Structs",
        "template": """// Go's _structs_ are typed collections of fields.
// They're useful for grouping data together to form
// records.

package main

import "fmt"

// This `person` struct type has `name` and `age` fields.
// Define struct person with name (string) and age (int) fields

// `newPerson` constructs a new person struct with the given name.
// Create function newPerson that takes a name string and returns *person
// Create person p with the given name
// Set p.age = 42
// Return &p

func main() {

	// This syntax creates a new struct.
	// Print person{"Bob", 20}

	// You can name the fields when initializing a struct.
	// Print person{name: "Alice", age: 30}

	// Omitted fields will be zero-valued.
	// Print person{name: "Fred"}

	// An `&` prefix yields a pointer to the struct.
	// Print &person{name: "Ann", age: 40}

	// It's idiomatic to encapsulate new struct creation in constructor functions
	// Print newPerson("Jon")

	// Access struct fields with a dot.
	s := person{name: "Sean", age: 50}
	// Print s.name

	// You can also use dots with struct pointers - the
	// pointers are automatically dereferenced.
	sp := &s
	// Print sp.age

	// Structs are mutable.
	// Set sp.age = 51
	// Print sp.age

	// If a struct type is only used for a single value, we don't
	// have to give it a name. The value can have an anonymous
	// struct type. This technique is commonly used for
	// [table-driven tests](testing-and-benchmarking).
	// Create anonymous struct dog with name (string) and isGood (bool) fields
	// Initialize with "Rex" and true
	// Print dog
}"""
    },

    "interfaces": {
        "display_name": "Interfaces",
        "template": """// _Interfaces_ are named collections of method
// signatures.

package main

import (
	"fmt"
	"math"
)

// Here's a basic interface for geometric shapes.
// Define interface geometry with area() float64 and perim() float64 methods

// For our example we'll implement this interface on
// `rect` and `circle` types.
// Define struct rect with width, height float64
// Define struct circle with radius float64

// To implement an interface in Go, we just need to
// implement all the methods in the interface. Here we
// implement `geometry` on `rect`s.
// Implement area() method on rect that returns width * height
// Implement perim() method on rect that returns 2*width + 2*height

// The implementation for `circle`s.
// Implement area() method on circle that returns Pi * radius^2
// Implement perim() method on circle that returns 2 * Pi * radius

// If a variable has an interface type, then we can call
// methods that are in the named interface. Here's a
// generic `measure` function taking advantage of this
// to work on any `geometry`.
// Create function measure that takes geometry parameter g
// Print g, g.area(), and g.perim()

// Sometimes it's useful to know the runtime type of an
// interface value. One option is using a *type assertion*
// as shown here; another is a [type `switch`](switch).
// Create function detectCircle that takes geometry parameter g
// Use type assertion to check if g is a circle
// If it is, print "circle with radius" followed by the radius

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	// The `circle` and `rect` struct types both
	// implement the `geometry` interface so we can use
	// instances of these structs as arguments to `measure`.
	// Call measure(r) and measure(c)

	// Call detectCircle(r) and detectCircle(c)
}"""
    }
}

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

def create_go_file(directory, package_name, topic_key):
    """Create a practice template .go file based on Go by Example patterns."""
    go_filename = f"{package_name}.go"
    go_filepath = os.path.join(directory, go_filename)

    # Get template content, fallback to basic template if not found
    template_info = PRACTICE_TEMPLATES.get(topic_key, {
        "display_name": topic_key.replace("-", " ").title(),
        "template": f'''package main

// {topic_key.replace("-", " ").title()} - Practice implementation

import "fmt"

func main() {{
    // TODO: Implement {topic_key.replace("-", " ").title()} concepts
    fmt.Println("Practicing: {topic_key.replace("-", " ").title()}")
}}
'''
    })

    with open(go_filepath, 'w') as f:
        f.write(template_info["template"])

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

def create_practice_module(topic_key, base_dir):
    """Create a complete practice module for a given topic."""
    package_name = topic_to_package_name(topic_key)
    directory = os.path.join(base_dir, package_name)

    display_name = PRACTICE_TEMPLATES.get(topic_key, {"display_name": topic_key.replace("-", " ").title()})["display_name"]
    print(f"\nCreating module for '{display_name}' -> {package_name}")

    # Create directory
    os.makedirs(directory, exist_ok=True)

    # Create .go file
    create_go_file(directory, package_name, topic_key)

    # Create go.mod file
    create_go_mod(directory, package_name)

def create_go_workspace(base_dir, topic_keys):
    """Create a go.work file to manage all modules in the workspace."""
    go_work_path = os.path.join(base_dir, "go.work")

    print("\nCreating Go workspace file...")

    go_work_content = "go 1.25\n\nuse (\n"
    for topic_key in topic_keys:
        package_name = topic_to_package_name(topic_key)
        go_work_content += f"    ./{package_name}\n"
    go_work_content += ")\n"

    with open(go_work_path, 'w') as f:
        f.write(go_work_content)

    print("  Created go.work (enables multi-module workspace)")

def clean_modules(base_dir, topic_keys):
    """Remove all practice modules and the go.work file."""
    print(f"🧹 Cleaning up Go practice modules in: {base_dir}")

    removed_count = 0

    # Remove all module directories
    for topic_key in topic_keys:
        package_name = topic_to_package_name(topic_key)
        directory = os.path.join(base_dir, package_name)

        if os.path.exists(directory):
            try:
                shutil.rmtree(directory)
                print(f"  Removed {package_name}/")
                removed_count += 1
            except Exception as e:
                print(f"  ❌ Failed to remove {package_name}/: {e}")
        else:
            print(f"  ⏭️  {package_name}/ doesn't exist")

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
    
    topic_keys = list(PRACTICE_TEMPLATES.keys())

    if args.clean:
        # Clean up modules
        clean_modules(script_dir, topic_keys)
    else:
        # Create modules (default behavior)
        print(f"🚀 Setting up Go practice modules in: {script_dir}")
        print(f"Total modules to create: {len(topic_keys)}")

        # Create all modules
        for topic_key in topic_keys:
            create_practice_module(topic_key, script_dir)

        # Create Go workspace file
        create_go_workspace(script_dir, topic_keys)

        print(f"\n✅ Successfully created {len(topic_keys)} Go practice modules!")
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
