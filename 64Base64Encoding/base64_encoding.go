// Go provides built-in support for [base64
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
	data := "abc123!?$*&()'-=@~"

	// Go supports both standard and URL-compatible
	// base64. Here's how to encode using the standard
	// encoder. The encoder requires a `[]byte` so we
	// convert our `string` to that type.

	// TODO: Create sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)
	// TODO: Print sEnc

	// Decoding may return an error, which you can check
	// if you don't already know the input to be
	// well-formed.

	// TODO: Create sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))
	// TODO: Print sDec
	fmt.Println(string(sDec))

	// This encodes/decodes using a URL-compatible base64
	// format.

	// TODO: Create uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)
	// TODO: Create uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))
	// TODO: Print fmt.Println()
}