package main

import (
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v3"
)

// Config represents a configuration structure with both JSON and YAML tags
type Config struct {
	Name     string  `json:"name" yaml:"name"`
	Override float64 `json:"override,omitempty" yaml:"override,omitempty"`
}

// UnmarshalYAMLOmitted unmarshals YAML with an omitted field
func UnmarshalYAMLOmitted() (Config, error) {
	yamlStr := `
name: test-service
# override field is omitted
`
	var config Config
	err := yaml.Unmarshal([]byte(yamlStr), &config)
	return config, err
}

// UnmarshalYAMLZero unmarshals YAML with an explicit zero value
func UnmarshalYAMLZero() (Config, error) {
	yamlStr := `
name: test-service
override: 0.0
`
	var config Config
	err := yaml.Unmarshal([]byte(yamlStr), &config)
	return config, err
}

// UnmarshalYAMLNonZero unmarshals YAML with a non-zero value
func UnmarshalYAMLNonZero() (Config, error) {
	yamlStr := `
name: test-service
override: 0.75
`
	var config Config
	err := yaml.Unmarshal([]byte(yamlStr), &config)
	return config, err
}

// UnmarshalJSONOmitted unmarshals JSON with an omitted field
func UnmarshalJSONOmitted() (Config, error) {
	jsonStr := `{
	"name": "test-service"
}`
	var config Config
	err := json.Unmarshal([]byte(jsonStr), &config)
	return config, err
}

// UnmarshalJSONZero unmarshals JSON with an explicit zero value
func UnmarshalJSONZero() (Config, error) {
	jsonStr := `{
	"name": "test-service",
	"override": 0.0
}`
	var config Config
	err := json.Unmarshal([]byte(jsonStr), &config)
	return config, err
}

// UnmarshalJSONNonZero unmarshals JSON with a non-zero value
func UnmarshalJSONNonZero() (Config, error) {
	jsonStr := `{
	"name": "test-service",
	"override": 0.75
}`
	var config Config
	err := json.Unmarshal([]byte(jsonStr), &config)
	return config, err
}

// MarshalYAMLZero marshals a Config struct with zero override value to YAML
func MarshalYAMLZero() (string, error) {
	config := Config{Name: "test-service", Override: 0.0}
	data, err := yaml.Marshal(config)
	return string(data), err
}

// MarshalYAMLNonZero marshals a Config struct with non-zero override to YAML
func MarshalYAMLNonZero() (string, error) {
	config := Config{Name: "test-service", Override: 0.75}
	data, err := yaml.Marshal(config)
	return string(data), err
}

// MarshalJSONZero marshals a Config struct with zero override value to JSON
func MarshalJSONZero() (string, error) {
	config := Config{Name: "test-service", Override: 0.0}
	data, err := json.Marshal(config)
	return string(data), err
}

// MarshalJSONNonZero marshals a Config struct with non-zero override to JSON
func MarshalJSONNonZero() (string, error) {
	config := Config{Name: "test-service", Override: 0.75}
	data, err := json.Marshal(config)
	return string(data), err
}

func main() {
	fmt.Println("=== YAML Unmarshaling Tests ===")

	config1, err := UnmarshalYAMLOmitted()
	if err != nil {
		panic(err)
	}
	fmt.Printf("YAML Omitted - Override: %f (is zero: %t)\n", config1.Override, config1.Override == 0.0)

	config2, err := UnmarshalYAMLZero()
	if err != nil {
		panic(err)
	}
	fmt.Printf("YAML Zero - Override: %f (is zero: %t)\n", config2.Override, config2.Override == 0.0)

	config3, err := UnmarshalYAMLNonZero()
	if err != nil {
		panic(err)
	}
	fmt.Printf("YAML Non-Zero - Override: %f (is zero: %t)\n", config3.Override, config3.Override == 0.0)

	fmt.Println("\n=== JSON Unmarshaling Tests ===")

	config4, err := UnmarshalJSONOmitted()
	if err != nil {
		panic(err)
	}
	fmt.Printf("JSON Omitted - Override: %f (is zero: %t)\n", config4.Override, config4.Override == 0.0)

	config5, err := UnmarshalJSONZero()
	if err != nil {
		panic(err)
	}
	fmt.Printf("JSON Zero - Override: %f (is zero: %t)\n", config5.Override, config5.Override == 0.0)

	config6, err := UnmarshalJSONNonZero()
	if err != nil {
		panic(err)
	}
	fmt.Printf("JSON Non-Zero - Override: %f (is zero: %t)\n", config6.Override, config6.Override == 0.0)

	fmt.Println("\n=== Marshaling Tests (with omitempty) ===")

	yamlZero, _ := MarshalYAMLZero()
	jsonZero, _ := MarshalJSONZero()
	fmt.Printf("Marshal Zero Value:\nYAML:\n%s\nJSON: %s\n", yamlZero, jsonZero)

	yamlNonZero, _ := MarshalYAMLNonZero()
	jsonNonZero, _ := MarshalJSONNonZero()
	fmt.Printf("Marshal Non-Zero Value:\nYAML:\n%s\nJSON: %s\n", yamlNonZero, jsonNonZero)

	fmt.Println("\n=== Conclusion ===")
	fmt.Println("✅ Omitted fields default to 0.0")
	fmt.Println("✅ Explicitly set 0.0 also results in 0.0")
	fmt.Println("✅ Cannot distinguish between omitted and explicit 0.0")
	fmt.Println("✅ omitempty tag excludes 0.0 values when marshaling")
}

// Notes:
// - YAML and JSON both use struct tags to map fields
// - omitempty tag excludes zero values when marshaling
// - Both formats cannot distinguish between omitted and explicit zero values
// - Use pointers (*float64) if you need to detect missing vs zero values

