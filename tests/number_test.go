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

func TestNumber(t *testing.T) {
	numProto := prototipe.NewPrototype(10)
	resultProto, err := numProto.Divide(2)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if resultProto.Value() != 5.0 {
		t.Errorf("Expected 5, got %v", resultProto.Value())
	}

	floatProto := prototipe.NewPrototype(15.5)
	resultProto, err = floatProto.Divide(3.1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if resultProto.Value() != 15.5/3.1 {
		t.Errorf("Expected %v, got %v", 15.5/3.1, resultProto.Value())
	}

	_, err = numProto.Divide(0)
	if err == nil {
		t.Error("Expected division by zero error, got nil")
	}

	strProto := prototipe.NewPrototype("hello")
	_, err = strProto.Divide(2)
	if err == nil {
		t.Error("Expected an error for non-number value, got nil")
	}
}

func TestIsEven(t *testing.T) {
	intProto := prototipe.NewPrototype(10)
	resultProto, err := intProto.IsEven()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if resultProto.Value() != true {
		t.Errorf("Expected true, got %v", resultProto.Value())
	}

	oddProto := prototipe.NewPrototype(7)
	resultProto, err = oddProto.IsEven()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if resultProto.Value() != false {
		t.Errorf("Expected false, got %v", resultProto.Value())
	}

	floatProto := prototipe.NewPrototype(10.5)
	_, err = floatProto.IsEven()
	if err == nil {
		t.Error("Expected an error for float value, got nil")
	}

	strProto := prototipe.NewPrototype("hello")
	_, err = strProto.IsEven()
	if err == nil {
		t.Error("Expected an error for non-number value, got nil")
	}
}
