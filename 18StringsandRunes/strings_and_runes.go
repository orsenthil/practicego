// A Go string is a read-only slice of bytes. The language
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
	fmt.Println("length of s:", len(s))

	// Since strings are equivalent to `[]byte`, this
	// will produce the length of the raw bytes stored within.

	// TODO: Print length of s
	for i := range s {
		fmt.Println("hex value:", s[i])
	}
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
	fmt.Println("rune count of s:", utf8.RuneCountInString(s))
	// A `range` loop handles strings specially and decodes
	// each `rune` along with its offset in the string.

	// TODO: Loop through s and print index and rune value
	for i, r := range s {
		fmt.Println("index:", i, "rune:", r)
	}
	// We can achieve the same iteration by using the
	// `utf8.DecodeRuneInString` function explicitly.


	// TODO: Loop through s and print index and rune value using utf8.DecodeRuneInString
	// Use for loop and i, w := 0, 0; i < len(s); i += w, where w is the width of the rune
	for i, w := 0, 0; i < len(s); i += w {
		r, width := utf8.DecodeRuneInString(s[i:])
		fmt.Println("index:", i, "rune:", r)
		w= width
	}	
	// Also demonstrate passing a `rune` value to a function, examineRune
	examineRune('t')
	examineRune('ส')
	fmt.Println("\nUsing DecodeRuneInString")

}

// TODO: Create function examineRune(r rune) that prints "found tee" if r is 't' and "found so sua" if r is 'ส'
func examineRune(r rune) {
	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}