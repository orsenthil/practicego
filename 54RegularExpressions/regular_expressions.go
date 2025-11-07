// Go offers built-in support for [regular expressions](https://en.wikipedia.org/wiki/Regular_expression).
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
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)


	// Above we used a string pattern directly, but for
	// other regexp tasks you'll need to `Compile` an
	// optimized `Regexp` struct.

	// TODO: Create r, _ := regexp.Compile("p([a-z]+)ch")
	r, _ := regexp.Compile("p([a-z]+)ch")
	// Many methods are available on these structs. Here's
	// a match test like we saw earlier.

	// TODO: Print r.MatchString("peach")
	fmt.Println(r.MatchString("peach"))
	// This finds the match for the regexp.
	// TODO: Print r.FindString("peach punch")
	fmt.Println(r.FindString("peach punch"))
	// This also finds the first match but returns the
	// start and end indexes for the match instead of the
	// matching text.
	fmt.Println(r.FindStringIndex("peach punch"))
	// TODO: Print r.FindStringIndex("peach punch")

	// The `Submatch` variants include information about
	// both the whole-pattern matches and the submatches
	// within those matches. For example this will return
	// information for both `p([a-z]+)ch` and `([a-z]+)`.
	fmt.Println(r.FindStringSubmatch("peach punch"))				
	// TODO: Print r.FindStringSubmatch("peach punch")

	// Similarly this will return information about the
	// indexes of matches and submatches.

	// TODO: Print r.FindStringSubmatchIndex("peach punch")
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))
	// The `All` variants of these functions apply to all
	// matches in the input, not just the first. For
	// example to find all matches for a regexp.

	// TODO: Print r.FindAllString("peach punch pinch", -1)
	fmt.Println(r.FindAllString("peach punch pinch", -1))
	// These `All` variants are available for the other
	// functions we saw above as well.

	// TODO: Print r.FindAllStringSubmatchIndex("peach punch pinch", -1)
	fmt.Println(r.FindAllStringSubmatchIndex("peach punch pinch", -1))
	// Providing a non-negative integer as the second
	// argument to these functions will limit the number
	// of matches.

	// TODO: Print r.FindAllString("peach punch pinch", 2)
	fmt.Println(r.FindAllString("peach punch pinch", 2))
	// Our examples above had string arguments and used
	// names like `MatchString`. We can also provide
	// `[]byte` arguments and drop `String` from the
	// function name.

	// TODO: Print r.Match([]byte("peach"))
	fmt.Println(r.Match([]byte("peach")))
	// When creating global variables with regular
	// expressions you can use the `MustCompile` variation
	// of `Compile`. `MustCompile` panics instead of
	// returning an error, which makes it safer to use for
	// global variables.
	r = regexp.MustCompile("p([a-z]+)ch")
	// TODO: Create r = regexp.MustCompile("p([a-z]+)ch")
	// TODO: Print r

	// The `regexp` package can also be used to replace
	// subsets of strings with other values.

	// TODO: Print r.ReplaceAllString("a peach", "<fruit>")
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))
	// The `Func` variant allows you to transform matched
	// text with a given function.
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
	// TODO: Create in := []byte("a peach")
	// TODO: Create out := r.ReplaceAllFunc(in, bytes.ToUpper)
	// TODO: Print string(out)

}