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
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Info("Worker starting...")
			// Simulate some startup work
			time.Sleep(100 * time.Millisecond)
			log.Info("Worker started successfully")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Info("Worker stopping...")
			// Simulate cleanup work
			time.Sleep(100 * time.Millisecond)
			log.Info("Worker stopped successfully")
			return nil
		},
	})

	fmt.Println("Worker registered with lifecycle")
}

func main() {
	fx.New(
		fx.Provide(NewWorker, zap.NewExample),
		fx.Invoke(RegisterWorker),
	).Run()
}

// Notes:
// - fx.Lifecycle manages the application lifecycle
// - OnStart hooks are called in the order they're registered
// - OnStop hooks are called in reverse order (LIFO)
// - Both hooks receive a context.Context for cancellation
// - If any OnStart returns an error, the app won't start
// - OnStop hooks should clean up resources created in OnStart





