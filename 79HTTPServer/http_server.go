// Writing a basic HTTP server is easy using the
// `net/http` package.
package main

import (
	"fmt"
	"net/http"
	"time"
)

// A fundamental concept in `net/http` servers is
// *handlers*. A handler is an object implementing the
// `http.Handler` interface. A common way to write
// a handler is by using the `http.HandlerFunc` adapter
// on functions with the appropriate signature.

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Println("server: hello handler started")
	defer fmt.Println("server: hello handler ended")

	time.Sleep(10 * time.Second)
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	fmt.Println("server: headers handler started")
	defer fmt.Println("server: headers handler ended")

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

// Functions serving as handlers take a
// `http.ResponseWriter` and a `http.Request` as
// arguments. The response writer is used to fill in the
// HTTP response. Here our simple response is just
// "hello ".


func main() {

	// We register our handlers on server routes using the
	// `http.HandleFunc` convenience function. It sets up
	// the *default router* in the `net/http` package and
	// takes a function as an argument.

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	// Finally, we call the `ListenAndServe` with the port
	// and a handler. `nil` tells it to use the default
	// router we've just set up.

	http.ListenAndServe(":8090", nil)
}