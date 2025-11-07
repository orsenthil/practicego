// In the previous example we looked at setting up a simple
// [HTTP server](http-server). HTTP servers are useful for
// demonstrating the usage of `context.Context` for
// controlling cancellation. A `Context` carries deadlines,
// cancellation signals, and other request-scoped values
// across API boundaries and goroutines.
package main

import (
	"fmt"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request) {

	// A `context.Context` is created for each request by
	// the `net/http` machinery, and is available with
	// the `Context()` method.

	// TODO: Create ctx := req.Context()
	ctx := req.Context()
	fmt.Println("server: hello handler started")
	defer fmt.Println("server: hello handler ended")

	// Wait for a few seconds before sending a reply to the
	// client. This could simulate some work the server is
	// doing. While working, keep an eye on the context's
	// `Done()` channel for a signal that we should cancel
	// the work and return as soon as possible.

	select {
	case <- time.After(10 * time.Second):
		fmt.Println("hello\n")
	case <- ctx.Done():
		fmt.Println("server: ", ctx.Err())
		http.Error(w, ctx.Err().Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("hello\n")
}

func main() {

	// As before, we register our handler on the "/hello"
	// route, and start serving.

	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8090", nil)
}