package main

import (
	"strings"
	"testing"
)

func TestUnmarshalYAMLOmitted(t *testing.T) {
	config, err := UnmarshalYAMLOmitted()
	if err != nil {
		t.Fatalf("UnmarshalYAMLOmitted failed: %v", err)
	}

	if config.Name != "test-service" {
		t.Errorf("Expected name 'test-service', got '%s'", config.Name)
	}

	if config.Override != 0.0 {
		t.Errorf("Expected Override to be 0.0 (default), got %f", config.Override)
	}
}

func TestUnmarshalYAMLZero(t *testing.T) {
	config, err := UnmarshalYAMLZero()
	if err != nil {
		t.Fatalf("UnmarshalYAMLZero failed: %v", err)
	}

	if config.Name != "test-service" {
		t.Errorf("Expected name 'test-service', got '%s'", config.Name)
	}

	if config.Override != 0.0 {
		t.Errorf("Expected Override to be 0.0 (explicit), got %f", config.Override)
	}
}

func TestUnmarshalYAMLNonZero(t *testing.T) {
	config, err := UnmarshalYAMLNonZero()
	if err != nil {
		t.Fatalf("UnmarshalYAMLNonZero failed: %v", err)
	}

	if config.Name != "test-service" {
		t.Errorf("Expected name 'test-service', got '%s'", config.Name)
	}

	if config.Override != 0.75 {
		t.Errorf("Expected Override to be 0.75, got %f", config.Override)
	}
}

func TestUnmarshalJSONOmitted(t *testing.T) {
	config, err := UnmarshalJSONOmitted()
	if err != nil {
		t.Fatalf("UnmarshalJSONOmitted failed: %v", err)
	}

	if config.Name != "test-service" {
		t.Errorf("Expected name 'test-service', got '%s'", config.Name)
	}

	if config.Override != 0.0 {
		t.Errorf("Expected Override to be 0.0 (default), got %f", config.Override)
	}
}

func TestUnmarshalJSONZero(t *testing.T) {
	config, err := UnmarshalJSONZero()
	if err != nil {
		t.Fatalf("UnmarshalJSONZero failed: %v", err)
	}

	if config.Name != "test-service" {
		t.Errorf("Expected name 'test-service', got '%s'", config.Name)
	}

	if config.Override != 0.0 {
		t.Errorf("Expected Override to be 0.0 (explicit), got %f", config.Override)
	}
}

func TestUnmarshalJSONNonZero(t *testing.T) {
	config, err := UnmarshalJSONNonZero()
	if err != nil {
		t.Fatalf("UnmarshalJSONNonZero failed: %v", err)
	}

	if config.Name != "test-service" {
		t.Errorf("Expected name 'test-service', got '%s'", config.Name)
	}

	if config.Override != 0.75 {
		t.Errorf("Expected Override to be 0.75, got %f", config.Override)
	}
}

func TestMarshalYAMLZero(t *testing.T) {
	yamlStr, err := MarshalYAMLZero()
	if err != nil {
		t.Fatalf("MarshalYAMLZero failed: %v", err)
	}

	// With omitempty, the override field should be omitted when zero
	if strings.Contains(yamlStr, "override") {
		t.Error("Expected YAML to omit 'override' field with zero value due to omitempty tag")
	}

	if !strings.Contains(yamlStr, "test-service") {
		t.Error("Expected YAML to contain 'test-service'")
	}
}

func TestMarshalYAMLNonZero(t *testing.T) {
	yamlStr, err := MarshalYAMLNonZero()
	if err != nil {
		t.Fatalf("MarshalYAMLNonZero failed: %v", err)
	}

	// With non-zero value, override should be included
	if !strings.Contains(yamlStr, "override") {
		t.Error("Expected YAML to include 'override' field with non-zero value")
	}

	if !strings.Contains(yamlStr, "0.75") {
		t.Error("Expected YAML to contain '0.75'")
	}

	if !strings.Contains(yamlStr, "test-service") {
		t.Error("Expected YAML to contain 'test-service'")
	}
}

func TestMarshalJSONZero(t *testing.T) {
	jsonStr, err := MarshalJSONZero()
	if err != nil {
		t.Fatalf("MarshalJSONZero failed: %v", err)
	}

	// With omitempty, the override field should be omitted when zero
	if strings.Contains(jsonStr, "override") {
		t.Error("Expected JSON to omit 'override' field with zero value due to omitempty tag")
	}

	if !strings.Contains(jsonStr, "test-service") {
		t.Error("Expected JSON to contain 'test-service'")
	}
}

func TestMarshalJSONNonZero(t *testing.T) {
	jsonStr, err := MarshalJSONNonZero()
	if err != nil {
		t.Fatalf("MarshalJSONNonZero failed: %v", err)
	}

	// With non-zero value, override should be included
	if !strings.Contains(jsonStr, "override") {
		t.Error("Expected JSON to include 'override' field with non-zero value")
	}

	if !strings.Contains(jsonStr, "0.75") {
		t.Error("Expected JSON to contain '0.75'")
	}

	if !strings.Contains(jsonStr, "test-service") {
		t.Error("Expected JSON to contain 'test-service'")
	}
}

func TestOmittedVsZeroValueBehavior(t *testing.T) {
	// This test documents the key learning: 
	// You cannot distinguish between omitted and explicit zero values
	
	omitted, _ := UnmarshalYAMLOmitted()
	explicit, _ := UnmarshalYAMLZero()

	if omitted.Override != explicit.Override {
		t.Error("Expected omitted and explicit zero to be equal")
	}

	if omitted.Override != 0.0 || explicit.Override != 0.0 {
		t.Error("Both should be 0.0")
	}

	// Both result in the same value - this is the key insight!
	t.Logf("Omitted: %f, Explicit Zero: %f - Cannot distinguish!", omitted.Override, explicit.Override)
}

func TestOmitemptyMarshaling(t *testing.T) {
	// Test that omitempty actually omits zero values when marshaling
	
	yamlZero, _ := MarshalYAMLZero()
	jsonZero, _ := MarshalJSONZero()

	if strings.Contains(yamlZero, "override") {
		t.Error("YAML should omit zero value with omitempty")
	}

	if strings.Contains(jsonZero, "override") {
		t.Error("JSON should omit zero value with omitempty")
	}

	yamlNonZero, _ := MarshalYAMLNonZero()
	jsonNonZero, _ := MarshalJSONNonZero()

	if !strings.Contains(yamlNonZero, "override") {
		t.Error("YAML should include non-zero value")
	}

	if !strings.Contains(jsonNonZero, "override") {
		t.Error("JSON should include non-zero value")
	}
}

