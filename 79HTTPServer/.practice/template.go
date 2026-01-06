// Writing a basic HTTP server is easy using the
// `net/http` package.
package main

import (
	"fmt"
	"net/http"
)

// A fundamental concept in `net/http` servers is
// *handlers*. A handler is an object implementing the
// `http.Handler` interface. A common way to write
// a handler is by using the `http.HandlerFunc` adapter
// on functions with the appropriate signature.

// TODO: Create function hello that takes a http.ResponseWriter and a http.Request
// Functions serving as handlers take a
// `http.ResponseWriter` and a `http.Request` as
// arguments. The response writer is used to fill in the
// HTTP response. Here our simple response is just
// "hello".


// TODO: Create function headers that takes a http.ResponseWriter and a http.Request
// Inside the function, iterate over the request headers with range and print the name and value



func main() {

	// We register our handlers on server routes using the
	// `http.HandleFunc` convenience function. It sets up
	// the *default router* in the `net/http` package and
	// takes a function as an argument.

	// TODO: Create http.HandleFunc("/hello", hello)
	// TODO: Create http.HandleFunc("/headers", headers)

	// Finally, we call the `ListenAndServe` with the port
	// and a handler. `nil` tells it to use the default
	// router we've just set up.

	// TODO: Create http.ListenAndServe(":8090", nil)
}