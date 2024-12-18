package tests

import (
	"fmt"
	"goprototipe/prototipe"
	"reflect"
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
	sliceProto := prototipe.NewPrototype([]interface{}{1, 2, 3, 4, 5})
	printProto, err := sliceProto.PrintSlice()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	expected := "[1 2 3 4 5]"
	if printProto.Value() != expected {
		t.Errorf("Expected %v, got %v", expected, printProto.Value())
	}
}

func TestUnique(t *testing.T) {
	sliceProto := prototipe.NewPrototype([]interface{}{1, 2, 2, 3, 4, 4, 5})
	uniqueProto, err := sliceProto.Unique()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	expected := []interface{}{1, 2, 3, 4, 5}
	if !reflect.DeepEqual(uniqueProto.Value(), expected) {
		t.Errorf("Expected %v, got %v", expected, uniqueProto.Value())
	}
}

func TestContains(t *testing.T) {
	sliceProto := prototipe.NewPrototype([]interface{}{1, 2, 3, 4, 5})

	containsProto, err := sliceProto.Contains(3)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if containsProto.Value() != true {
		t.Errorf("Expected true, got %v", containsProto.Value())
	}
}

func TestIndexOf(t *testing.T) {
	sliceProto := prototipe.NewPrototype([]interface{}{1, 2, 3, 4, 5})

	indexProto, err := sliceProto.IndexOf(3)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if indexProto.Value() != 2 {
		t.Errorf("Expected index 2, got %v", indexProto.Value())
	}
}
