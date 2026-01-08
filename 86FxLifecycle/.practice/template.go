// FX Lifecycle allows you to hook into the application's
// start and stop phases. This is useful for managing resources
// like servers, database connections, or background workers.

package main

import (
	"context"
	"fmt"
	"time"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

// A Worker service that needs to start and stop gracefully
type Worker struct {
	log *zap.Logger
}

// NewWorker creates a new Worker instance
func NewWorker(log *zap.Logger) *Worker {
	return &Worker{log: log}
}

// RegisterLifecycle adds lifecycle hooks to the Worker
// The fx.Lifecycle parameter is automatically injected by FX
func RegisterWorker(lc fx.Lifecycle, worker *Worker, log *zap.Logger) {
	// TODO: Use lc.Append() to add lifecycle hooks
	// 
	// The hook should have:
	// - OnStart: A function that logs "Worker starting..." and simulates startup
	// - OnStop: A function that logs "Worker stopping..." and simulates cleanup
	//
	// Example structure:
	// lc.Append(fx.Hook{
	//     OnStart: func(ctx context.Context) error {
	//         log.Info("Worker starting...")
	//         // Simulate some startup work
	//         time.Sleep(100 * time.Millisecond)
	//         log.Info("Worker started successfully")
	//         return nil
	//     },
	//     OnStop: func(ctx context.Context) error {
	//         log.Info("Worker stopping...")
	//         // Simulate cleanup work
	//         time.Sleep(100 * time.Millisecond)
	//         log.Info("Worker stopped successfully")
	//         return nil
	//     },
	// })
	
	fmt.Println("Worker registered with lifecycle")
}

func main() {
	// TODO: Create an FX application with:
	// 1. fx.Provide(NewWorker, zap.NewExample) - provide constructors
	// 2. fx.Invoke(RegisterWorker) - register lifecycle hooks
	// 3. Call Run() to start the app
	//
	// Try running the app and observe the startup/shutdown logs
	// Press Ctrl+C to trigger shutdown and see OnStop being called
}

// Notes:
// - fx.Lifecycle manages the application lifecycle
// - OnStart hooks are called in the order they're registered
// - OnStop hooks are called in reverse order (LIFO)
// - Both hooks receive a context.Context for cancellation
// - If any OnStart returns an error, the app won't start
// - OnStop hooks should clean up resources created in OnStart





