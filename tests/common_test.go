package tests

import (
	"github.com/genov8/goprototipe/prototipe"
	"reflect"
	"testing"
)

func TestLength(t *testing.T) {
	str := prototipe.NewPrototype("Hello")
	length, err := str.Length()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if length != 5 {
		t.Errorf("Expected 5, got %v", length)
	}

	slice := prototipe.NewPrototype([]interface{}{1, 2, 3, 4})
	length, err = slice.Length()
	if err != nil {
		t.Errorf("Expected no error for slice, got %v", err)
	}
	if length != 4 {
		t.Errorf("Expected length 4 for slice, got %v", length)
	}
}

func TestReverse(t *testing.T) {
	strProto := prototipe.NewPrototype("Hello")
	reversedProto, err := strProto.Reverse()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	expectedStr := "olleH"
	if reversedProto.Value() != expectedStr {
		t.Errorf("Expected %v, got %v", expectedStr, reversedProto.Value())
	}

	sliceProto := prototipe.NewPrototype([]interface{}{1, 2, 3, 4})
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
	str1 := prototipe.NewPrototype("Hello")
	str2 := prototipe.NewPrototype(" World!")
	newStr, err := str1.Concat(str2)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if newStr.Value() != "Hello World!" {
		t.Errorf("Expected 'Hello World', got %v", newStr.Value())
	}

	slice1 := prototipe.NewPrototype([]interface{}{1, 2, 3})
	slice2 := prototipe.NewPrototype([]interface{}{4, 5, 6})
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
	num := prototipe.NewPrototype(10)
	result, err := num.Compare(10)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result.Value() != true {
		t.Errorf("Expected true, got %v", result.Value())
	}

	str := prototipe.NewPrototype("Hello")
	result, err = str.Compare("World")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result.Value() != false {
		t.Errorf("Expected false, got %v", result.Value())
	}

	slice := prototipe.NewPrototype([]interface{}{1, 2, 3})
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
