// The Uber FX library provides a dependency injection framework
// for Go applications. It helps manage application lifecycle and
// dependencies in a structured way.

package main

import (
	"fmt"

	"go.uber.org/fx"
)

// A simple greeter service that will be injected
type Greeter struct {
	message string
}

// NewGreeter is a constructor function that FX will use
// to create a Greeter instance
func NewGreeter() *Greeter {
	return &Greeter{message: "Hello from FX!"}
}

// Greet prints a greeting message
func (g *Greeter) Greet(name string) {
	fmt.Printf("%s Welcome, %s!\n", g.message, name)
}

// A printer service that depends on Greeter
type Printer struct {
	greeter *Greeter
}

// NewPrinter is a constructor that depends on Greeter.
// FX will automatically inject the Greeter dependency.
func NewPrinter(g *Greeter) *Printer {
	return &Printer{greeter: g}
}

// Print uses the greeter to print a welcome message
func (p *Printer) Print() {
	p.greeter.Greet("FX User")
}

func main() {
	fx.New(
		fx.Provide(NewGreeter, NewPrinter),
		fx.Invoke(func(p *Printer) {
			p.Print()
		}),
	).Run()
}

// Notes:
// - fx.Provide() registers constructor functions
// - fx.Invoke() runs functions with automatic dependency injection
// - FX automatically resolves and injects dependencies based on function signatures
// - Run() starts the application and blocks until it's shut down

