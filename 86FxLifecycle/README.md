# 86FxLifecycle - FX Lifecycle Management

## Overview

This module teaches you how to use FX's lifecycle hooks to manage application startup and shutdown. Lifecycle hooks are essential for resources that need graceful initialization and cleanup, such as servers, database connections, and background workers.

## Concepts Covered

- **fx.Lifecycle**: Interface for managing application lifecycle
- **OnStart Hooks**: Functions called when the application starts
- **OnStop Hooks**: Functions called when the application stops
- **Graceful Shutdown**: Properly cleaning up resources on exit

## Key Points

1. **Hook Registration**: Use `lc.Append(fx.Hook{...})` to register lifecycle hooks
2. **Context Support**: Both OnStart and OnStop receive a `context.Context` for cancellation
3. **Error Handling**: If OnStart returns an error, the app won't start
4. **LIFO Order**: OnStop hooks run in reverse order (Last In, First Out)
5. **Signal Handling**: FX automatically handles SIGTERM and SIGINT for graceful shutdown

## How to Practice

1. Navigate to the `.practice` directory
2. Open `template.go` and implement the lifecycle hooks
3. Run: `go run template.go`
4. Observe the startup logs
5. Press `Ctrl+C` to trigger shutdown and see cleanup logs

## Expected Output

```
Worker registered with lifecycle
Worker starting...
Worker started successfully
^C (when you press Ctrl+C)
Worker stopping...
Worker stopped successfully
```

## Common Use Cases

- Starting/stopping HTTP servers
- Opening/closing database connections
- Initializing/shutting down background workers
- Acquiring/releasing resources

## Next Steps

After completing this module, move on to:
- **87FxGroups** - Learn about value groups and advanced dependency patterns





