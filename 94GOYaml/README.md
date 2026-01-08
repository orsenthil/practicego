# 94GOYaml - YAML and JSON Marshaling in Go

## Overview

This practice module teaches YAML and JSON marshaling/unmarshaling in Go, with a focus on understanding the `omitempty` tag behavior and the critical distinction (or lack thereof) between omitted and explicit zero values.

## Challenge: Configuration Parsing

Understand how Go handles YAML and JSON parsing, particularly:
- How omitted fields behave
- How explicit zero values behave  
- The difference between marshaling and unmarshaling
- The effect of the `omitempty` tag

## Concepts Covered

- **YAML Unmarshaling**: Parsing YAML strings into Go structs
- **JSON Unmarshaling**: Parsing JSON strings into Go structs
- **YAML Marshaling**: Converting Go structs to YAML
- **JSON Marshaling**: Converting Go structs to JSON
- **Struct Tags**: Using `yaml` and `json` tags
- **omitempty Tag**: Understanding its behavior
- **Zero Values**: Go's default values for types
- **Field Presence**: Detecting missing vs zero values

## Data Model

```go
type Config struct {
    Name     string  `json:"name" yaml:"name"`
    Override float64 `json:"override,omitempty" yaml:"override,omitempty"`
}
```

### Tag Explanation

- `json:"name"` - JSON field name mapping
- `yaml:"name"` - YAML field name mapping
- `omitempty` - Omit field from output if it has a zero value

## Required Functions

Implement these 10 functions:

### Unmarshal Functions (6)

1. **UnmarshalYAMLOmitted() (Config, error)**
   - Parse YAML where `override` field is omitted

2. **UnmarshalYAMLZero() (Config, error)**
   - Parse YAML where `override` is explicitly `0.0`

3. **UnmarshalYAMLNonZero() (Config, error)**
   - Parse YAML where `override` is `0.75`

4. **UnmarshalJSONOmitted() (Config, error)**
   - Parse JSON where `override` field is omitted

5. **UnmarshalJSONZero() (Config, error)**
   - Parse JSON where `override` is explicitly `0.0`

6. **UnmarshalJSONNonZero() (Config, error)**
   - Parse JSON where `override` is `0.75`

### Marshal Functions (4)

7. **MarshalYAMLZero() (string, error)**
   - Convert Config with `Override=0.0` to YAML

8. **MarshalYAMLNonZero() (string, error)**
   - Convert Config with `Override=0.75` to YAML

9. **MarshalJSONZero() (string, error)**
   - Convert Config with `Override=0.0` to JSON

10. **MarshalJSONNonZero() (string, error)**
    - Convert Config with `Override=0.75` to JSON

## Key Learning Points

### 1. Omitted vs Explicit Zero

**Critical Insight**: You CANNOT distinguish between an omitted field and an explicit zero value!

```go
// YAML with omitted field
name: test-service
# override not specified

// YAML with explicit zero
name: test-service
override: 0.0
```

**Both result in `config.Override == 0.0`**

### 2. omitempty During Marshaling

The `omitempty` tag affects **output**, not input:

```go
config := Config{Name: "test", Override: 0.0}
yaml.Marshal(config)  // "name: test\n" (override omitted!)

config := Config{Name: "test", Override: 0.75}
yaml.Marshal(config)  // "name: test\noverride: 0.75\n" (included!)
```

### 3. Zero Values in Go

Go has default "zero values" for all types:
- `float64` → `0.0`
- `int` → `0`
- `string` → `""`
- `bool` → `false`
- pointers → `nil`

### 4. Detecting Missing Fields

If you need to distinguish "not provided" from "explicitly zero", use **pointers**:

```go
type Config struct {
    Override *float64 `json:"override,omitempty" yaml:"override,omitempty"`
}

// nil = field not provided
// &0.0 = explicitly zero
// &0.75 = non-zero value
```

## How to Practice

1. Navigate to the `.practice` directory
2. Open `template.go` and complete the TODOs
3. Uncomment the main function to test
4. Run: `go run template.go`
5. Compare with `solution.go` if needed

## Expected Output

```
=== YAML Unmarshaling Tests ===
YAML Omitted - Override: 0.000000 (is zero: true)
YAML Zero - Override: 0.000000 (is zero: true)
YAML Non-Zero - Override: 0.750000 (is zero: false)

=== JSON Unmarshaling Tests ===
JSON Omitted - Override: 0.000000 (is zero: true)
JSON Zero - Override: 0.000000 (is zero: true)
JSON Non-Zero - Override: 0.750000 (is zero: false)

=== Marshaling Tests (with omitempty) ===
Marshal Zero Value:
YAML:
name: test-service

JSON: {"name":"test-service"}
Marshal Non-Zero Value:
YAML:
name: test-service
override: 0.75

JSON: {"name":"test-service","override":0.75}

=== Conclusion ===
✅ Omitted fields default to 0.0
✅ Explicitly set 0.0 also results in 0.0
✅ Cannot distinguish between omitted and explicit 0.0
✅ omitempty tag excludes 0.0 values when marshaling
```

## Testing Requirements

Your solution should pass all 12 tests:
- ✅ Unmarshal YAML with omitted field
- ✅ Unmarshal YAML with explicit zero
- ✅ Unmarshal YAML with non-zero value
- ✅ Unmarshal JSON with omitted field
- ✅ Unmarshal JSON with explicit zero
- ✅ Unmarshal JSON with non-zero value
- ✅ Marshal YAML with zero (omitted due to omitempty)
- ✅ Marshal YAML with non-zero (included)
- ✅ Marshal JSON with zero (omitted due to omitempty)
- ✅ Marshal JSON with non-zero (included)
- ✅ Demonstrate omitted equals explicit zero
- ✅ Verify omitempty marshaling behavior

## Common Use Cases

### Configuration Files

```go
// config.yaml
database:
  host: localhost
  port: 5432
  # timeout field omitted - will default to 0

// config.go
type DatabaseConfig struct {
    Host    string `yaml:"host"`
    Port    int    `yaml:"port"`
    Timeout int    `yaml:"timeout,omitempty"`
}
```

### API Responses

```go
type Response struct {
    Status  string  `json:"status"`
    Message string  `json:"message,omitempty"`
    Data    *Data   `json:"data,omitempty"`  // Omit if nil
}
```

### Feature Flags

```go
type Features struct {
    EnableNewUI     bool    `yaml:"enable_new_ui"`
    MaxConnections  int     `yaml:"max_connections,omitempty"`
    RateLimitFactor float64 `yaml:"rate_limit_factor,omitempty"`
}
```

## Best Practices

### 1. Use Pointers for Optional Fields

```go
// ❌ Bad: Can't tell if user didn't provide or set to zero
type Config struct {
    Timeout int `yaml:"timeout,omitempty"`
}

// ✅ Good: nil means not provided, &0 means explicitly zero
type Config struct {
    Timeout *int `yaml:"timeout,omitempty"`
}
```

### 2. Validate After Unmarshaling

```go
config, err := parseConfig(yamlStr)
if err != nil {
    return err
}

// Validate required fields
if config.Name == "" {
    return errors.New("name is required")
}
```

### 3. Provide Defaults

```go
config := Config{
    Name:    "default-name",
    Timeout: 30,  // default value
}
yaml.Unmarshal(data, &config)  // Overrides defaults if present
```

### 4. Document Zero Value Behavior

```go
type Config struct {
    // RetryCount is the number of retries.
    // If omitted, defaults to 0 (no retries).
    RetryCount int `yaml:"retry_count,omitempty"`
}
```

## Advanced Patterns

### Custom Unmarshaler

```go
type Duration time.Duration

func (d *Duration) UnmarshalYAML(value *yaml.Node) error {
    var s string
    if err := value.Decode(&s); err != nil {
        return err
    }
    dur, err := time.ParseDuration(s)
    *d = Duration(dur)
    return err
}
```

### Strict Decoding

```go
decoder := yaml.NewDecoder(reader)
decoder.KnownFields(true)  // Error on unknown fields
err := decoder.Decode(&config)
```

### Multiple Tag Sets

```go
type Config struct {
    Name string `json:"name" yaml:"name" toml:"name"`
}
```

## Common Pitfalls

1. **Assuming omitempty affects unmarshaling** - It only affects marshaling!
2. **Not handling errors** - Always check unmarshal errors
3. **Using wrong struct tags** - yaml.v3 uses `yaml:`, not `yml:`
4. **Forgetting to take addresses** - Unmarshal needs a pointer: `&config`
5. **Expecting to detect omitted vs zero** - Use pointers if you need this

## Learning Resources

- [YAML Package Documentation](https://pkg.go.dev/gopkg.in/yaml.v3)
- [JSON Package Documentation](https://pkg.go.dev/encoding/json)
- [YAML Specification](https://yaml.org/spec/)
- [Go Struct Tags](https://go.dev/wiki/Well-known-struct-tags)

## Extensions (Optional Challenges)

1. **Custom Types**: Implement UnmarshalYAML for time.Duration
2. **Nested Structs**: Parse multi-level configurations
3. **Validation**: Add struct validation with custom rules
4. **Environment Variables**: Combine YAML with env var overrides
5. **Multiple Formats**: Support YAML, JSON, and TOML
6. **Schema Validation**: Validate against a schema
7. **Hot Reload**: Watch config file for changes
8. **Secrets Management**: Handle encrypted config values

## Real-World Applications

- **Application Configuration**: Server settings, database credentials
- **CI/CD Pipelines**: GitHub Actions, GitLab CI configurations
- **Kubernetes**: Pod specs, deployments, services
- **Docker Compose**: Multi-container configurations
- **API Contracts**: Request/response schemas

---

**Note**: This module demonstrates a fundamental Go pattern that applies to many serialization formats!

