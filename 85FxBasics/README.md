# 85FxBasics - Introduction to Uber FX Dependency Injection

## Overview

This practice module introduces the Uber FX library, a powerful dependency injection framework for Go applications. FX helps manage application structure, dependencies, and lifecycle in a clean and maintainable way.

## Concepts Covered

- **Dependency Injection**: Automatic wiring of dependencies based on constructor signatures
- **fx.Provide()**: Registering constructor functions
- **fx.Invoke()**: Running functions with automatic dependency resolution
- **Constructor Pattern**: Creating services through constructor functions

## Key Points

1. **Constructor Functions**: Functions that return pointers to structs become providers
2. **Automatic Resolution**: FX looks at function signatures and automatically provides dependencies
3. **Type-Based Injection**: Dependencies are matched by type
4. **Inversion of Control**: FX manages object creation and lifecycle

## How to Practice

1. Navigate to the `.practice` directory
2. Open `template.go` and complete the TODOs
3. Run the code: `go run template.go`
4. Compare with `solution.go` if you get stuck

## Expected Output

```
Hello from FX! Welcome, FX User!
```

## Learning Resources

- [Uber FX Documentation](https://uber-go.github.io/fx/)
- [Dependency Injection in Go](https://blog.drewolson.org/dependency-injection-in-go)

## Next Steps

After completing this module, move on to:
- **86FxLifecycle** - Learn about application lifecycle management

