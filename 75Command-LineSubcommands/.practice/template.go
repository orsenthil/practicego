// Some command-line tools, like the `go` tool or `git`
// have many *subcommands*, each with its own set of
// flags. For example, `go build` and `go get` are two
// different subcommands of the `go` tool.
// The `flag` package lets us easily define simple
// subcommands that have their own flags.

package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	// We declare a subcommand using the `NewFlagSet`
	// function, and proceed to define new flags specific
	// for this subcommand.

	// TODO: Create fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)
	// TODO: Create fooEnable := fooCmd.Bool("enable", false, "enable")
	// TODO: Create fooName := fooCmd.String("name", "", "name")

	// For a different subcommand we can define different
	// supported flags.

	// TODO: Create barCmd := flag.NewFlagSet("bar", flag.ExitOnError)
	// TODO: Create barLevel := barCmd.Int("level", 0, "level")

	// The subcommand is expected as the first argument
	// to the program.

	// TODO: If len(os.Args) < 2, fmt.Println("expected 'foo' or 'bar' subcommands")
	// TODO: os.Exit(1)

	// Check which subcommand is invoked.

	// TODO: switch os.Args[1] {
	// If the case is "foo", parse the fooCmd flags, if the case is "bar", parse the barCmd flags.
	// default print "expected 'foo' or 'bar' subcommands" and exit with 1

	// For every subcommand, we parse its own flags and
	// have access to trailing positional arguments.
	
}