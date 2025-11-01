// [_Command-line arguments_](https://en.wikipedia.org/wiki/Command-line_interface#Arguments)
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
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	// You can get individual args with normal indexing.

	// TODO: Create arg := os.Args[3]	
	// TODO: Print arg
	arg := os.Args[3]
	fmt.Println("arg:", arg)
	// TODO: Print argsWithProg
	// TODO: Print argsWithoutProg
	// TODO: Print arg
	fmt.Println("argsWithProg:", argsWithProg)
	fmt.Println("argsWithoutProg:", argsWithoutProg)
	fmt.Println("arg:", arg)
}