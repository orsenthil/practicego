// Complete FX HTTP Server example combining:
// - Dependency injection
// - Lifecycle management
// - Value groups for route registration
// This is the full pattern from the fxdemo application.

package main

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

// Route is an http.Handler that knows the mux pattern
// under which it will be registered.
type Route interface {
	http.Handler
	// Pattern reports the path at which this is registered.
	Pattern() string
}

// EchoHandler is an http.Handler that copies its request body
// back to the response.
type EchoHandler struct {
	log *zap.Logger
}

// NewEchoHandler builds a new EchoHandler.
func NewEchoHandler(log *zap.Logger) *EchoHandler {
	return &EchoHandler{log: log}
}

func (*EchoHandler) Pattern() string {
	return "/echo"
}

// ServeHTTP handles an HTTP request to the /echo endpoint.
func (h *EchoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if _, err := io.Copy(w, r.Body); err != nil {
		h.log.Warn("Failed to handle request", zap.Error(err))
	}
}

// HelloHandler is an HTTP handler that prints a greeting to the user.
type HelloHandler struct {
	log *zap.Logger
}

// NewHelloHandler builds a new HelloHandler.
func NewHelloHandler(log *zap.Logger) *HelloHandler {
	return &HelloHandler{log: log}
}

func (*HelloHandler) Pattern() string {
	return "/hello"
}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.log.Error("Failed to read request", zap.Error(err))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if _, err := fmt.Fprintf(w, "Hello, %s\n", body); err != nil {
		h.log.Error("Failed to write response", zap.Error(err))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

// AsRoute annotates the given constructor to state that
// it provides a route to the "routes" group.
func AsRoute(f any) any {
	// TODO: Implement AsRoute using fx.Annotate
	// Should cast result to Route interface and add to "routes" group
	// 
	// Example:
	// return fx.Annotate(
	//     f,
	//     fx.As(new(Route)),
	//     fx.ResultTags(`group:"routes"`),
	// )
	
	return f // Replace this
}

// NewServeMux builds a ServeMux that will route requests to the given routes.
func NewServeMux(routes []Route) *http.ServeMux {
	// TODO: Create a new http.ServeMux and register all routes
	// Iterate through routes and call mux.Handle(route.Pattern(), route)
	//
	// Example:
	// mux := http.NewServeMux()
	// for _, route := range routes {
	//     mux.Handle(route.Pattern(), route)
	// }
	// return mux
	
	return nil // Replace this
}

// NewHTTPServer builds an HTTP server that will begin serving requests
// when the Fx application starts.
func NewHTTPServer(lc fx.Lifecycle, mux *http.ServeMux, log *zap.Logger) *http.Server {
	srv := &http.Server{Addr: ":8080", Handler: mux}
	
	// TODO: Use lc.Append() to add lifecycle hooks for the server
	// OnStart should:
	//   1. Create a TCP listener on srv.Addr
	//   2. Log that the server is starting
	//   3. Start srv.Serve(ln) in a goroutine
	//   4. Return any error from net.Listen
	// OnStop should:
	//   1. Call srv.Shutdown(ctx) to gracefully shut down the server
	//
	// Example:
	// lc.Append(fx.Hook{
	//     OnStart: func(ctx context.Context) error {
	//         ln, err := net.Listen("tcp", srv.Addr)
	//         if err != nil {
	//             return err
	//         }
	//         log.Info("Starting HTTP server", zap.String("addr", srv.Addr))
	//         go srv.Serve(ln)
	//         return nil
	//     },
	//     OnStop: func(ctx context.Context) error {
	//         return srv.Shutdown(ctx)
	//     },
	// })
	
	return srv
}

func main() {
	// TODO: Create a complete FX application that:
	// 1. Sets up logging with fx.WithLogger and fxevent.ZapLogger
	// 2. Provides all necessary constructors:
	//    - zap.NewExample for logging
	//    - NewHTTPServer for the server
	//    - NewServeMux with fx.Annotate and fx.ParamTags(`group:"routes"`)
	//    - AsRoute(NewEchoHandler) for the /echo endpoint
	//    - AsRoute(NewHelloHandler) for the /hello endpoint
	// 3. Invokes a function with *http.Server to ensure it's created
	// 4. Calls Run() to start the application
	//
	// Example structure:
	// fx.New(
	//     fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
	//         return &fxevent.ZapLogger{Logger: log}
	//     }),
	//     fx.Provide(
	//         zap.NewExample,
	//         NewHTTPServer,
	//         fx.Annotate(
	//             NewServeMux,
	//             fx.ParamTags(`group:"routes"`),
	//         ),
	//         AsRoute(NewEchoHandler),
	//         AsRoute(NewHelloHandler),
	//     ),
	//     fx.Invoke(func(*http.Server) {}),
	// ).Run()
	//
	// After implementing, test with:
	//   curl -X POST http://localhost:8080/echo -d "test"
	//   curl -X POST http://localhost:8080/hello -d "World"
}

// Notes:
// - This combines all FX patterns: DI, Lifecycle, and Value Groups
// - The server starts automatically when the app starts
// - Routes are registered via the "routes" group pattern
// - The server shuts down gracefully when the app stops (Ctrl+C)
// - fx.WithLogger configures custom logging for FX events
// - Each handler gets the logger injected automatically

