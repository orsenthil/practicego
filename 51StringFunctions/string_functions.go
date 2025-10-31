// The standard library's `strings` package provides many
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
	fmt.Println("Contains:", s.Contains("test", "es"))
	// TODO: Print Count:     s.Count("test", "t")
	fmt.Println("Count:", s.Count("test", "t"))
	// TODO: Print HasPrefix: s.HasPrefix("test", "te")
	fmt.Println("HasPrefix:", s.HasPrefix("test", "te"))
	// TODO: Print HasSuffix: s.HasSuffix("test", "st")
	fmt.Println("HasSuffix:", s.HasSuffix("test", "st"))
	// TODO: Print Index:     s.Index("test", "e")
	fmt.Println("Index:", s.Index("test", "e"))
	// TODO: Print Join:      s.Join([]string{"a", "b"}, "-")
	fmt.Println("Join:", s.Join([]string{"a", "b"}, "-"))
	// TODO: Print Repeat:    s.Repeat("a", 5)
	fmt.Println("Repeat:", s.Repeat("a", 5))
	// TODO: Print Replace:   s.Replace("foo", "o", "0", -1)
	fmt.Println("Replace:", s.Replace("foo", "o", "0", -1))
	// TODO: Print Split:     s.Split("a-b-c-d-e", "-")
	fmt.Println("Split:", s.Split("a-b-c-d-e", "-"))
	// TODO: Print ToLower:   s.ToLower("TEST")
	fmt.Println("ToLower:", s.ToLower("TEST"))
	// TODO: Print ToUpper:   s.ToUpper("test")
	fmt.Println("ToUpper:", s.ToUpper("test"))

}