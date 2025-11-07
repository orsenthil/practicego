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

	fmt.Println(s.Contains("test", "es"))
	fmt.Println(s.Count("test", "t"))
	fmt.Println(s.HasPrefix("test", "te"))
	fmt.Println(s.HasSuffix("test", "st"))
	fmt.Println(s.Index("test", "e"))
	fmt.Println(s.Join([]string{"a", "b"}, "-"))
	fmt.Println(s.Repeat("a", 5))
	fmt.Println(s.Replace("foo", "o", "0", -1))
	fmt.Println(s.Split("a-b-c-d-e", "-"))
	fmt.Println(s.ToLower("TEST"))
	fmt.Println(s.ToUpper("test"))
}