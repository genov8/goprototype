package tests

import (
	"fmt"
	"goprototipe/prototipe"
	"testing"
)

func TestToUpperCase(t *testing.T) {
	str := prototipe.NewPrototype("hello")
	upperStr, err := str.ToUpperCase()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if upperStr.Value() != "HELLO" {
		t.Errorf("Expected 'HELLO', got %v", upperStr.Value())
	}
}

func TestToLowerCase(t *testing.T) {
	str := prototipe.NewPrototype("HELLO")
	upperStr, err := str.ToLowerCase()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if upperStr.Value() != "hello" {
		t.Errorf("Expected 'hello', got %v", upperStr.Value())
	}
}

func TestCapitalize(t *testing.T) {
	str := prototipe.NewPrototype("HELLO")
	upperStr, err := str.Capitalize()
	fmt.Println(upperStr)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if upperStr.Value() != "Hello" {
		t.Errorf("Expected 'Hello', got %v", upperStr.Value())
	}
}
