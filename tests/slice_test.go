package tests

import (
	"fmt"
	"goprototipe/prototipe"
	"testing"
)

func TestAppend(t *testing.T) {
	sliceProto := prototipe.NewPrototype([]interface{}{1, 2, 3})

	newSliceProto, err := sliceProto.Append(4)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	expected := []interface{}{1, 2, 3, 4}
	if fmt.Sprintf("%v", newSliceProto.Value()) != fmt.Sprintf("%v", expected) {
		t.Errorf("Expected %v, got %v", expected, newSliceProto.Value())
	}
}

func TestRemove(t *testing.T) {
	sliceProto := prototipe.NewPrototype([]interface{}{1, 2, 3})

	newSliceProto, err := sliceProto.Remove(1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	expected := []interface{}{1, 3}
	if fmt.Sprintf("%v", newSliceProto.Value()) != fmt.Sprintf("%v", expected) {
		t.Errorf("Expected %v, got %v", expected, newSliceProto.Value())
	}

	_, err = newSliceProto.Remove(5)
	if err == nil {
		t.Error("Expected error for out of bounds index, got none")
	}
}

func TestPrintSlice(t *testing.T) {
	sliceProto := prototipe.NewPrototype([]interface{}{1, 2, 3, 4})

	result, err := sliceProto.PrintSlice()

	if err != nil {
		t.Errorf("Expected no error while printing slice, got: %v", err)
	}

	expected := "[1 2 3 4]"
	if result != expected {
		t.Errorf("Expected printed slice %s, got %s", expected, result)
	}
}
