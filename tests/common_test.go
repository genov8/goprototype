package tests

import (
	"github.com/genov8/goprototype/prototype"
	"reflect"
	"testing"
)

func TestLength(t *testing.T) {
	str := prototype.NewPrototype("Hello")
	length, err := str.Length()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if length != 5 {
		t.Errorf("Expected 5, got %v", length)
	}

	slice := prototype.NewPrototype([]interface{}{1, 2, 3, 4})
	length, err = slice.Length()
	if err != nil {
		t.Errorf("Expected no error for slice, got %v", err)
	}
	if length != 4 {
		t.Errorf("Expected length 4 for slice, got %v", length)
	}
}

func TestReverse(t *testing.T) {
	strProto := prototype.NewPrototype("Hello")
	reversedProto, err := strProto.Reverse()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	expectedStr := "olleH"
	if reversedProto.Value() != expectedStr {
		t.Errorf("Expected %v, got %v", expectedStr, reversedProto.Value())
	}

	sliceProto := prototype.NewPrototype([]interface{}{1, 2, 3, 4})
	reversedSliceProto, err := sliceProto.Reverse()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	expectedSlice := []interface{}{4, 3, 2, 1}
	if !reflect.DeepEqual(reversedSliceProto.Value(), expectedSlice) {
		t.Errorf("Expected %v, got %v", expectedSlice, reversedSliceProto.Value())
	}
}

func TestConcat(t *testing.T) {
	str1 := prototype.NewPrototype("Hello")
	str2 := prototype.NewPrototype(" World!")
	newStr, err := str1.Concat(str2)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if newStr.Value() != "Hello World!" {
		t.Errorf("Expected 'Hello World', got %v", newStr.Value())
	}

	slice1 := prototype.NewPrototype([]interface{}{1, 2, 3})
	slice2 := prototype.NewPrototype([]interface{}{4, 5, 6})
	concatSliceProto, err := slice1.Concat(slice2)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	expectedSlice := []interface{}{1, 2, 3, 4, 5, 6}
	if !reflect.DeepEqual(concatSliceProto.Value(), expectedSlice) {
		t.Errorf("Expected %v, got %v", expectedSlice, concatSliceProto.Value())
	}
}

func TestCompare(t *testing.T) {
	num := prototype.NewPrototype(10)
	result, err := num.Compare(10)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result.Value() != true {
		t.Errorf("Expected true, got %v", result.Value())
	}

	str := prototype.NewPrototype("Hello")
	result, err = str.Compare("World")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result.Value() != false {
		t.Errorf("Expected false, got %v", result.Value())
	}

	slice := prototype.NewPrototype([]interface{}{1, 2, 3})
	result, err = slice.Compare([]interface{}{1, 2, 3})
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result.Value() != true {
		t.Errorf("Expected true, got %v", result.Value())
	}

	result, err = slice.Compare([]interface{}{1, 2})
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result.Value() != false {
		t.Errorf("Expected false, got %v", result.Value())
	}
}

func TestType(t *testing.T) {
	str := prototype.NewPrototype("Hello")
	result := str.Type()
	if result.Value() != "string" {
		t.Errorf("Expected 'string', got %v", result.Value())
	}

	num := prototype.NewPrototype(42)
	result = num.Type()
	if result.Value() != "int" {
		t.Errorf("Expected 'int', got %v", result.Value())
	}

	slice := prototype.NewPrototype([]interface{}{1, 2, 3})
	result = slice.Type()
	if result.Value() != "[]interface {}" {
		t.Errorf("Expected '[]interface {}', got %v", result.Value())
	}
}

func TestIsEmpty(t *testing.T) {
	str := prototype.NewPrototype("")
	result, err := str.IsEmpty()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result.Value() != true {
		t.Errorf("Expected true, got %v", result.Value())
	}

	str = prototype.NewPrototype("Go")
	result, err = str.IsEmpty()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result.Value() != false {
		t.Errorf("Expected false, got %v", result.Value())
	}

	slice := prototype.NewPrototype([]interface{}{})
	result, err = slice.IsEmpty()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result.Value() != true {
		t.Errorf("Expected true, got %v", result.Value())
	}

	slice = prototype.NewPrototype([]interface{}{1, 2, 3})
	result, err = slice.IsEmpty()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result.Value() != false {
		t.Errorf("Expected false, got %v", result.Value())
	}

	num := prototype.NewPrototype(0)
	result, err = num.IsEmpty()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result.Value() != true {
		t.Errorf("Expected true, got %v", result.Value())
	}

	unknown := prototype.NewPrototype(struct{}{})
	_, err = unknown.IsEmpty()
	if err == nil {
		t.Errorf("Expected an error, got nil")
	}
}
