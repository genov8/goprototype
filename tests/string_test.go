package tests

import (
	"github.com/genov8/goprototipe/prototipe"
	"reflect"
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
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if upperStr.Value() != "Hello" {
		t.Errorf("Expected 'Hello', got %v", upperStr.Value())
	}
}

func TestTrim(t *testing.T) {
	str := prototipe.NewPrototype("Hello World!")
	upperStr, err := str.Trim("!")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if upperStr.Value() != "Hello World" {
		t.Errorf("Expected 'Hello World', got %v", upperStr.Value())
	}
}

func TestSplit(t *testing.T) {
	str := prototipe.NewPrototype("Hello World")
	splittedStr, err := str.Split(" ")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	expected := []string{"Hello", "World"}
	if !reflect.DeepEqual(splittedStr.Value(), expected) {
		t.Errorf("Expected %v, got %v", expected, splittedStr.Value())
	}
}

func TestStartsWith(t *testing.T) {
	str := prototipe.NewPrototype("Hello world")
	strStartsWith, err := str.StartsWith("Hello")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if strStartsWith.Value() != true {
		t.Errorf("Expected true, got %v", strStartsWith.Value())
	}
}

func TestEndsWith(t *testing.T) {
	str := prototipe.NewPrototype("Hello world")
	strEndsWith, err := str.EndsWith("world")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if strEndsWith.Value() != true {
		t.Errorf("Expected true, got %v", strEndsWith.Value())
	}
}

func TestRemoveWhitespace(t *testing.T) {
	str := prototipe.NewPrototype("Hello World!")

	strRemoveWhitespace, err := str.RemoveWhitespace()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	expected := "HelloWorld!"
	if strRemoveWhitespace.Value() != expected {
		t.Errorf("Expected %v, got %v", expected, strRemoveWhitespace.Value())
	}
}

func TestReplace(t *testing.T) {
	strPrototype := prototipe.NewPrototype("hello world")
	result, err := strPrototype.Replace("world", "Go")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result.Value() != "hello Go" {
		t.Errorf("Expected 'hello Go', got %v", result.Value())
	}

	result, err = strPrototype.Replace("planet", "Earth")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result.Value() != "hello world" {
		t.Errorf("Expected 'hello world', got %v", result.Value())
	}

	intPrototype := prototipe.NewPrototype(123)
	_, err = intPrototype.Replace("1", "2")
	if err == nil {
		t.Errorf("Expected an error for non-string value, got none")
	}
}
