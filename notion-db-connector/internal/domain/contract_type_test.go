package domain

import "testing"

func TestNewContractTypeContract(t *testing.T) {
	expected := "contract"
	result, err := NewContractType(expected)
	if err != nil {
		t.Fatalf("expected no error, but got %v", err)
	}
	if result.Value != expected {
		t.Fatalf("expected %v, but got %v", expected, result.Value)
	}
}

func TestNewContractTypeTraial(t *testing.T) {
	expected := "trial"
	result, err := NewContractType(expected)

	if err != nil {
		t.Fatalf("expected no error, but got %v", err)
	}

	if result.Value != expected {
		t.Fatalf("expected %v, but got %v", expected, result.Value)
	}
}

func TestNewContractTypeEmpty(t *testing.T) {
	expected := ""
	result, err := NewContractType(expected)

	if err == nil {
		t.Fatalf("expected error, but no error.")
	}

	if result != nil {
		t.Fatalf("expected nil, but got %v", result.Value)
	}
}

func TestNewContractTypeInvalid(t *testing.T) {
	expected := "invalid"
	result, err := NewContractType(expected)

	if err == nil {
		t.Fatalf("expected error, but no error.")
	}

	if result != nil {
		t.Fatalf("expected nil, but got %v", result.Value)
	}
}
