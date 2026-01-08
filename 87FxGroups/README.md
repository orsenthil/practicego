# 87FxGroups - FX Value Groups and Annotations

## Overview

This module demonstrates FX's value groups feature, which allows you to collect multiple implementations of the same interface and inject them as a slice. This is perfect for plugin architectures, middleware chains, and extensible systems.

## Concepts Covered

- **Value Groups**: Collecting multiple values with the same group tag
- **fx.Annotate()**: Adding metadata to constructors
- **fx.As()**: Casting results to interface types
- **fx.ResultTags()**: Marking values for group inclusion
- **fx.ParamTags()**: Receiving grouped values as a slice

## Key Points

1. **Group Pattern**: Use `group:"name"` tags to organize related values
2. **Interface Casting**: `fx.As(new(Interface))` casts concrete types to interfaces
3. **Collection**: All values with the same group tag are collected into a slice
4. **Order**: Values are provided in the order they're registered
5. **Type Safety**: FX ensures type compatibility at startup

## How to Practice

1. Navigate to the `.practice` directory
2. Open `template.go` and implement the `AsPlugin` function and main
3. Run: `go run template.go`
4. Observe all plugins being executed
5. Try adding your own plugin type

## Expected Output

```
Running 3 plugins:
  [Logger]: Logging some information...
  [Greeter]: Hello from Greeter Plugin!
  [Calculator]: 2 + 2 = 4
```

## Common Use Cases

- Plugin systems (as demonstrated)
- HTTP middleware chains
- Route handlers
- Event listeners
- Validation rules
- Database migrations

## Pattern Example

```go
// Mark a constructor as providing to a group
fx.Annotate(
    NewHandler,
    fx.As(new(Handler)),           // Cast to interface
    fx.ResultTags(`group:"handlers"`), // Add to group
)

// Consume the group as a slice
fx.Annotate(
    NewServer,
    fx.ParamTags(`group:"handlers"`), // Receive all handlers
)
```

## Next Steps

After completing this module, move on to:
- **88FxHTTPServer** - Combine all patterns in a real HTTP server





