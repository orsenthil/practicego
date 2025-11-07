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

	// Since strings are equivalent to `[]byte`, this
	// will produce the length of the raw bytes stored within.

	fmt.Println("len(s) =", len(s))

	// Indexing into a string produces the raw byte values at
	// each index. This loop generates the hex values of all
	// the bytes that constitute the code points in `s`.

	for i := range s {
		fmt.Printf("%x ", s[i])
	}

	// To count how many _runes_ are in a string, we can use
	// the `utf8` package. Note that the run-time of
	// `RuneCountInString` depends on the size of the string,
	// because it has to decode each UTF-8 rune sequentially.
	// Some Thai characters are represented by UTF-8 code points
	// that can span multiple bytes, so the result of this count
	// may be surprising.

	fmt.Println("utf8.RuneCountInString(s) =", utf8.RuneCountInString(s))

	// A `range` loop handles strings specially and decodes
	// each `rune` along with its offset in the string.

	for i, r := range s {
		fmt.Printf("index = %d, rune = %c\n", i, r)
	}

	// We can achieve the same iteration by using the
	// `utf8.DecodeRuneInString` function explicitly.


	for i, w := 0, 0; i < len(s); i += w {
		r, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("index = %d, rune = %c\n", i, r)
		w = width
	}

	fmt.Println("\nUsing DecodeRuneInString")

	examineRune('t')
	examineRune('ส')

}

func examineRune(r rune) {
	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}