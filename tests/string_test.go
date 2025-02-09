package tests

import (
	"github.com/genov8/goprototype/v2/prototype"
	"reflect"
	"testing"
)

func TestToUpperCase(t *testing.T) {
	str := prototype.NewStringPrototype("hello")
	upperStr, err := str.ToUpperCase()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if upperStr.Value() != "HELLO" {
		t.Errorf("Expected 'HELLO', got %v", upperStr.Value())
	}
}

func TestToLowerCase(t *testing.T) {
	str := prototype.NewStringPrototype("HELLO")
	upperStr, err := str.ToLowerCase()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if upperStr.Value() != "hello" {
		t.Errorf("Expected 'hello', got %v", upperStr.Value())
	}
}

func TestCapitalize(t *testing.T) {
	str := prototype.NewStringPrototype("HELLO")
	upperStr, err := str.Capitalize()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if upperStr.Value() != "Hello" {
		t.Errorf("Expected 'Hello', got %v", upperStr.Value())
	}
}

func TestTrim(t *testing.T) {
	str := prototype.NewStringPrototype("Hello World!")
	upperStr, err := str.Trim("!")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if upperStr.Value() != "Hello World" {
		t.Errorf("Expected 'Hello World', got %v", upperStr.Value())
	}
}

func TestSplit(t *testing.T) {
	str := prototype.NewStringPrototype("Hello World")
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
	str := prototype.NewStringPrototype("Hello world")
	strStartsWith, err := str.StartsWith("Hello")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if strStartsWith.Value() != true {
		t.Errorf("Expected true, got %v", strStartsWith.Value())
	}
}

func TestEndsWith(t *testing.T) {
	str := prototype.NewStringPrototype("Hello world")
	strEndsWith, err := str.EndsWith("world")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if strEndsWith.Value() != true {
		t.Errorf("Expected true, got %v", strEndsWith.Value())
	}
}

func TestRemoveWhitespace(t *testing.T) {
	str := prototype.NewStringPrototype("Hello World!")

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
	strPrototype := prototype.NewStringPrototype("hello world")
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

	intPrototype := prototype.NewStringPrototype(123)
	_, err = intPrototype.Replace("1", "2")
	if err == nil {
		t.Errorf("Expected an error for non-string value, got none")
	}
}

func TestRepeat(t *testing.T) {
	strPrototype := prototype.NewStringPrototype("Go")
	result, err := strPrototype.Repeat(3)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result.Value() != "GoGoGo" {
		t.Errorf("Expected 'GoGoGo', got %v", result.Value())
	}

	result, err = strPrototype.Repeat(0)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result.Value() != "" {
		t.Errorf("Expected an empty string, got %v", result.Value())
	}

	_, err = strPrototype.Repeat(-1)
	if err == nil {
		t.Errorf("Expected an error for negative repeat count, got none")
	}

	intPrototype := prototype.NewStringPrototype(123)
	_, err = intPrototype.Repeat(3)
	if err == nil {
		t.Errorf("Expected an error for non-string value, got none")
	}
}

func TestReverseWords(t *testing.T) {
	str := prototype.NewStringPrototype("Hello World Go")
	result, err := str.ReverseWords()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	expected := "Go World Hello"
	if result.Value() != expected {
		t.Errorf("Expected %v, got %v", expected, result.Value())
	}

	str = prototype.NewStringPrototype("Hello")
	result, err = str.ReverseWords()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	expected = "Hello"
	if result.Value() != expected {
		t.Errorf("Expected %v, got %v", expected, result.Value())
	}

	str = prototype.NewStringPrototype("")
	result, err = str.ReverseWords()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	expected = ""
	if result.Value() != expected {
		t.Errorf("Expected %v, got %v", expected, result.Value())
	}
}

func TestSlugify(t *testing.T) {
	str := prototype.NewStringPrototype("Hello, World!")
	result, err := str.Slugify()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result.Value() != "hello-world" {
		t.Errorf("Expected 'hello-world', got %v", result.Value())
	}

	str = prototype.NewStringPrototype("simple-slug")
	result, err = str.Slugify()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result.Value() != "simple-slug" {
		t.Errorf("Expected 'simple-slug', got %v", result.Value())
	}

	str = prototype.NewStringPrototype(42)
	_, err = str.Slugify()
	if err == nil {
		t.Errorf("Expected an error, got nil")
	}
}

func TestPad(t *testing.T) {
	str := prototype.NewStringPrototype("Go")
	result, err := str.Pad(5, "*", false)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	expected := "Go***"
	if result.Value() != expected {
		t.Errorf("Expected %v, got %v", expected, result.Value())
	}

	str = prototype.NewStringPrototype("Go")
	result, err = str.Pad(5, "*", true)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	expected = "***Go"
	if result.Value() != expected {
		t.Errorf("Expected %v, got %v", expected, result.Value())
	}

	str = prototype.NewStringPrototype("Golang")
	result, err = str.Pad(5, "*", false)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	expected = "Golang"
	if result.Value() != expected {
		t.Errorf("Expected %v, got %v", expected, result.Value())
	}

	str = prototype.NewStringPrototype("Go")
	_, err = str.Pad(5, "**", false)
	if err == nil {
		t.Errorf("Expected an error, got nil")
	}

	notStr := prototype.NewStringPrototype(42)
	_, err = notStr.Pad(5, "*", false)
	if err == nil {
		t.Errorf("Expected an error, got nil")
	}
}

func TestWordCount(t *testing.T) {
	str := prototype.NewStringPrototype("Hello world from Go")
	result, err := str.WordCount()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	expected := 4
	if result.Value() != expected {
		t.Errorf("Expected %v, got %v", expected, result.Value())
	}

	str = prototype.NewStringPrototype("  Hello    world  ")
	result, err = str.WordCount()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	expected = 2
	if result.Value() != expected {
		t.Errorf("Expected %v, got %v", expected, result.Value())
	}

	str = prototype.NewStringPrototype("")
	result, err = str.WordCount()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	expected = 0
	if result.Value() != expected {
		t.Errorf("Expected %v, got %v", expected, result.Value())
	}

	notStr := prototype.NewStringPrototype(42)
	_, err = notStr.WordCount()
	if err == nil {
		t.Errorf("Expected an error, got nil")
	}
}

func TestToCamelCase(t *testing.T) {
	str := prototype.NewStringPrototype("hello world")
	result, err := str.ToCamelCase()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	expected := "HelloWorld"
	if result.Value() != expected {
		t.Errorf("Expected %v, got %v", expected, result.Value())
	}

	str = prototype.NewStringPrototype("snake_case_example")
	result, err = str.ToCamelCase()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	expected = "SnakeCaseExample"
	if result.Value() != expected {
		t.Errorf("Expected %v, got %v", expected, result.Value())
	}

	str = prototype.NewStringPrototype("kebab-case-example")
	result, err = str.ToCamelCase()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	expected = "KebabCaseExample"
	if result.Value() != expected {
		t.Errorf("Expected %v, got %v", expected, result.Value())
	}

	str = prototype.NewStringPrototype("123 test example")
	result, err = str.ToCamelCase()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	expected = "123TestExample"
	if result.Value() != expected {
		t.Errorf("Expected %v, got %v", expected, result.Value())
	}

	notStr := prototype.NewStringPrototype(42)
	_, err = notStr.ToCamelCase()
	if err == nil {
		t.Errorf("Expected an error, got nil")
	}
}
