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
	// TODO: Print "server: hello handler started"
	// TODO: Defer "server: hello handler ended"

	// Wait for a few seconds before sending a reply to the
	// client. This could simulate some work the server is
	// doing. While working, keep an eye on the context's
	// `Done()` channel for a signal that we should cancel
	// the work and return as soon as possible.

	// TODO: Use select to wait for 10 seconds or the context's Done() channel
	// When context is <- time.After(10 * time.Second), print "hello
"
	// When context is <- ctx.Done(), print "server: " and the context's Err()
	// and set the status code to http.StatusInternalServerError and call http.Error(w, err.Error(), internalError)
	
	}
}

func main() {

	// As before, we register our handler on the "/hello"
	// route, and start serving.

	// TODO: Create http.HandleFunc("/hello", hello)
	// TODO: Create http.ListenAndServe(":8090", nil)
}