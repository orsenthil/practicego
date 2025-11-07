// URLs provide a [uniform way to locate resources](https://adam.herokuapp.com/past/2010/3/30/urls_are_the_uniform_way_to_locate_resources/).
// Here's how to parse URLs in Go.

package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {

	// We'll parse this example URL, which includes a
	// scheme, authentication info, host, port, path,
	// query params, and query fragment.

	// TODO: Create s := "postgres://user:pass@host.com:5432/path?k=v#f"
	s := "postgres://user:pass@host.com:5432/path?k=v#f"
	// Parse the URL and ensure there are no errors.

	// TODO: Create u, err := url.Parse(s)
	u, err := url.Parse(s)
	fmt.Println(err)

	// Accessing the scheme is straightforward.
	fmt.Println(u.Scheme)
	// TODO: Print u.Scheme

	// `User` contains all authentication info; call
	// `Username` and `Password` on this for individual
	// values.
	fmt.Println(u.User)
	// TODO: Print u.User
	fmt.Println(u.User.Username())
	p, _ := u.User.Password()
	fmt.Println(p)
	// TODO: Print p

	// The `Host` contains both the hostname and the port,
	// if present. Use `SplitHostPort` to extract them.
	fmt.Println(u.Host)
	// TODO: Print u.Host
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(host)
	fmt.Println(port)
	// TODO: Print host
	// TODO: Print port

	// Here we extract the `path` and the fragment after
	// the `#`.
	fmt.Println(u.Path)
	// TODO: Print u.Path
	// TODO: Print u.Fragment
	fmt.Println(u.Fragment)
	// To get query params in a string of `k=v` format,
	// use `RawQuery`. You can also parse query params
	// into a map. The parsed query param maps are from
	// strings to slices of strings, so index into `[0]`
	// if you only want the first value.
	fmt.Println(u.RawQuery)
	// TODO: Print u.RawQuery
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	// TODO: Print m
	// TODO: Print m["k"][0]
}