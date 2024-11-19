package tests

import (
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
