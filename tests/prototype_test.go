package tests

import (
	"goprototipe/prototipe"
	"testing"
)

func TestAdd(t *testing.T) {
	num := prototipe.NewPrototype(10)
	newNum, err := num.Add(5)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if newNum.Value() != 15 {
		t.Errorf("Expected 15, got %v", newNum.Value())
	}
}

func TestMultiply(t *testing.T) {
	num := prototipe.NewPrototype(10)
	newNum, err := num.Multiply(3)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if newNum.Value() != 30 {
		t.Errorf("Expected 30, got %v", newNum.Value())
	}
}

func TestConcat(t *testing.T) {
	str := prototipe.NewPrototype("Hello")
	newStr, err := str.Concat(" World")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if newStr.Value() != "Hello World" {
		t.Errorf("Expected 'Hello World', got %v", newStr.Value())
	}
}
