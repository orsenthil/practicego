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
	// TODO: Unmarshal the YAML string below into a Config struct
	// The override field is omitted in the YAML
	yamlStr := `
name: test-service
# override field is omitted
`
	var config Config
	// TODO: Use yaml.Unmarshal to parse yamlStr into config
	// Return the config and any error
	return config, nil
}

// UnmarshalYAMLZero unmarshals YAML with an explicit zero value
func UnmarshalYAMLZero() (Config, error) {
	// TODO: Unmarshal YAML with override explicitly set to 0.0
	yamlStr := `
name: test-service
override: 0.0
`
	var config Config
	// TODO: Use yaml.Unmarshal to parse yamlStr into config
	return config, nil
}

// UnmarshalYAMLNonZero unmarshals YAML with a non-zero value
func UnmarshalYAMLNonZero() (Config, error) {
	// TODO: Unmarshal YAML with override set to 0.75
	yamlStr := `
name: test-service
override: 0.75
`
	var config Config
	// TODO: Use yaml.Unmarshal to parse yamlStr into config
	return config, nil
}

// UnmarshalJSONOmitted unmarshals JSON with an omitted field
func UnmarshalJSONOmitted() (Config, error) {
	// TODO: Unmarshal JSON string with override field omitted
	jsonStr := `{
	"name": "test-service"
}`
	var config Config
	// TODO: Use json.Unmarshal to parse jsonStr into config
	return config, nil
}

// UnmarshalJSONZero unmarshals JSON with an explicit zero value
func UnmarshalJSONZero() (Config, error) {
	// TODO: Unmarshal JSON with override explicitly set to 0.0
	jsonStr := `{
	"name": "test-service",
	"override": 0.0
}`
	var config Config
	// TODO: Use json.Unmarshal to parse jsonStr into config
	return config, nil
}

// UnmarshalJSONNonZero unmarshals JSON with a non-zero value
func UnmarshalJSONNonZero() (Config, error) {
	// TODO: Unmarshal JSON with override set to 0.75
	jsonStr := `{
	"name": "test-service",
	"override": 0.75
}`
	var config Config
	// TODO: Use json.Unmarshal to parse jsonStr into config
	return config, nil
}

// MarshalYAMLZero marshals a Config struct with zero override value to YAML
func MarshalYAMLZero() (string, error) {
	// TODO: Create a Config with Name="test-service" and Override=0.0
	// Marshal it to YAML and return the string
	var config Config
	// TODO: Use yaml.Marshal to convert config to YAML
	return "", nil
}

// MarshalYAMLNonZero marshals a Config struct with non-zero override to YAML
func MarshalYAMLNonZero() (string, error) {
	// TODO: Create a Config with Name="test-service" and Override=0.75
	// Marshal it to YAML and return the string
	var config Config
	// TODO: Use yaml.Marshal to convert config to YAML
	return "", nil
}

// MarshalJSONZero marshals a Config struct with zero override value to JSON
func MarshalJSONZero() (string, error) {
	// TODO: Create a Config with Name="test-service" and Override=0.0
	// Marshal it to JSON and return the string
	var config Config
	// TODO: Use json.Marshal to convert config to JSON
	return "", nil
}

// MarshalJSONNonZero marshals a Config struct with non-zero override to JSON
func MarshalJSONNonZero() (string, error) {
	// TODO: Create a Config with Name="test-service" and Override=0.75
	// Marshal it to JSON and return the string
	var config Config
	// TODO: Use json.Marshal to convert config to JSON
	return "", nil
}

func main() {
	// TODO: Uncomment and complete when ready to test
	/*
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
	*/
}

// Notes:
// - YAML and JSON both use struct tags to map fields
// - omitempty tag excludes zero values when marshaling
// - Both formats cannot distinguish between omitted and explicit zero values
// - Use pointers (*float64) if you need to detect missing vs zero values

