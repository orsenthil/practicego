// The Go standard library comes with excellent support
// for HTTP clients and servers in the `net/http`
// package. In this example we'll use it to issue simple
// HTTP requests.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {

	// Issue an HTTP GET request to a server. `http.Get` is a
	// convenient shortcut around creating an `http.Client`
	// object and calling its `Get` method; it uses the
	// `http.DefaultClient` object which has useful default
	// settings.

	// TODO: Create resp, err := http.Get("https://gobyexample.com")
	// TODO: Print err

	resp, err := http.Get("https://gobyexample.com")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// TODO: Defer the closing of the response body with defer resp.Body.Close()

	// Print the HTTP response status.

	// TODO: Print fmt.Println("Response status:", resp.Status)
	fmt.Println("Response status:", resp.Status)
	// Print the first 5 lines of the response body.

	// TODO: Create scanner := bufio.NewScanner(resp.Body)
	// Iterate over the scanner with for i := 0; scanner.Scan() && i < 5; i++ {
	// TODO: Print fmt.Println(scanner.Text())
	// TODO: Print err

	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Check for scanner.Err() and if it is not nil, panic with the error
}