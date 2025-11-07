// The Go standard library comes with excellent support
// for HTTP clients and servers in the `net/http`
// package. In this example we'll use it to issue simple
// HTTP requests.
package main

import (
	"bufio"
	"fmt"
	"net/http"

)

func main() {

	// Issue an HTTP GET request to a server. `http.Get` is a
	// convenient shortcut around creating an `http.Client`
	// object and calling its `Get` method; it uses the
	// `http.DefaultClient` object which has useful default
	// settings.

	// TODO: Create resp, err := http.Get("https://gobyexample.com")
	resp, err := http.Get("https://gobyexample.com")
	fmt.Println(err)
	

	// TODO: Defer the closing of the response body with defer resp.Body.Close()
	defer resp.Body.Close()
	// Print the HTTP response status.

	// TODO: Print fmt.Println("Response status:", resp.Status)
	fmt.Println("Response status:", resp.Status)
	// Print the first 5 lines of the response body.

	// TODO: Create scanner := bufio.NewScanner(resp.Body)
	scanner := bufio.NewScanner(resp.Body)
	// Iterate over the scanner with for i := 0; scanner.Scan() && i < 5; i++ {
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}
	fmt.Println(err)
}