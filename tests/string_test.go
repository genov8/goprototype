package tests

import (
	"fmt"
	"goprototipe/prototipe"
	"testing"
)

func TestToUpper(t *testing.T) {
	str := prototipe.NewPrototype("hello")
	upperStr, err := str.ToUpper()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if upperStr.Value() != "HELLO" {
		t.Errorf("Expected 'HELLO', got %v", upperStr.Value())
	}
}

func TestToLower(t *testing.T) {
	str := prototipe.NewPrototype("HELLO")
	upperStr, err := str.ToLower()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if upperStr.Value() != "hello" {
		t.Errorf("Expected 'hello', got %v", upperStr.Value())
	}
}

func TestToUpperFirst(t *testing.T) {
	str := prototipe.NewPrototype("HELLO")
	upperStr, err := str.ToUpperFirst()
	fmt.Println(upperStr)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if upperStr.Value() != "Hello" {
		t.Errorf("Expected 'Hello', got %v", upperStr.Value())
	}
}
