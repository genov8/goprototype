package tests

import (
	"fmt"
	"github.com/genov8/goprototipe/prototipe"
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

func TestFlatten(t *testing.T) {
	nestedSlice := prototipe.NewPrototype([]interface{}{
		[]interface{}{1, 2},
		[]interface{}{3, []interface{}{4, 5}},
		6,
	})
	result, err := nestedSlice.Flatten()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	expected := []interface{}{1, 2, 3, 4, 5, 6}
	if !reflect.DeepEqual(result.Value(), expected) {
		t.Errorf("Expected %v, got %v", expected, result.Value())
	}

	flatSlice := prototipe.NewPrototype([]interface{}{1, 2, 3})
	result, err = flatSlice.Flatten()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if !reflect.DeepEqual(result.Value(), flatSlice.Value()) {
		t.Errorf("Expected %v, got %v", flatSlice.Value(), result.Value())
	}

	nonSlice := prototipe.NewPrototype(123)
	_, err = nonSlice.Flatten()
	if err == nil {
		t.Errorf("Expected an error for non-slice value, got none")
	}
}

func TestMap(t *testing.T) {
	slicePrototype := prototipe.NewPrototype([]interface{}{1, 2, 3, 4})
	result, err := slicePrototype.Map(func(val interface{}) interface{} {
		if num, ok := val.(int); ok {
			return num * 2
		}
		return val
	})
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	expected := []interface{}{2, 4, 6, 8}
	if !reflect.DeepEqual(result.Value(), expected) {
		t.Errorf("Expected %v, got %v", expected, result.Value())
	}

	slicePrototype = prototipe.NewPrototype([]interface{}{"a", "b", "c"})
	result, err = slicePrototype.Map(func(val interface{}) interface{} {
		if str, ok := val.(string); ok {
			return str + "!"
		}
		return val
	})
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	expected = []interface{}{"a!", "b!", "c!"}
	if !reflect.DeepEqual(result.Value(), expected) {
		t.Errorf("Expected %v, got %v", expected, result.Value())
	}

	intPrototype := prototipe.NewPrototype(123)
	_, err = intPrototype.Map(func(val interface{}) interface{} {
		return val
	})
	if err == nil {
		t.Errorf("Expected an error for non-slice value, got none")
	}
}

func TestFilter(t *testing.T) {
	slicePrototype := prototipe.NewPrototype([]interface{}{1, 2, 3, 4, 5})
	result, err := slicePrototype.Filter(func(val interface{}) bool {
		if num, ok := val.(int); ok {
			return num%2 == 0
		}
		return false
	})
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	expected := []interface{}{2, 4}
	if !reflect.DeepEqual(result.Value(), expected) {
		t.Errorf("Expected %v, got %v", expected, result.Value())
	}

	slicePrototype = prototipe.NewPrototype([]interface{}{"apple", "banana", "cherry"})
	result, err = slicePrototype.Filter(func(val interface{}) bool {
		if str, ok := val.(string); ok {
			return len(str) > 5
		}
		return false
	})
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	expected = []interface{}{"banana", "cherry"}
	if !reflect.DeepEqual(result.Value(), expected) {
		t.Errorf("Expected %v, got %v", expected, result.Value())
	}

	intPrototype := prototipe.NewPrototype(123)
	_, err = intPrototype.Filter(func(val interface{}) bool {
		return true
	})
	if err == nil {
		t.Errorf("Expected an error for non-slice value, got none")
	}
}

func TestChunk(t *testing.T) {
	slice := prototipe.NewPrototype([]interface{}{1, 2, 3, 4, 5, 6})
	result, err := slice.Chunk(2)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	expected := [][]interface{}{
		{1, 2},
		{3, 4},
		{5, 6},
	}
	if !reflect.DeepEqual(result.Value(), expected) {
		t.Errorf("Expected %v, got %v", expected, result.Value())
	}

	slice = prototipe.NewPrototype([]interface{}{1, 2, 3, 4, 5})
	result, err = slice.Chunk(2)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	expected = [][]interface{}{
		{1, 2},
		{3, 4},
		{5},
	}
	if !reflect.DeepEqual(result.Value(), expected) {
		t.Errorf("Expected %v, got %v", expected, result.Value())
	}

	_, err = slice.Chunk(0)
	if err == nil {
		t.Errorf("Expected an error for chunk size <= 0, got none")
	}

	notSlice := prototipe.NewPrototype(123)
	_, err = notSlice.Chunk(2)
	if err == nil {
		t.Errorf("Expected an error for non-slice value, got none")
	}
}

func TestIntersect(t *testing.T) {
	slice1 := prototipe.NewPrototype([]interface{}{1, 2, 3, 4})
	slice2 := prototipe.NewPrototype([]interface{}{3, 4, 5, 6})
	result, err := slice1.Intersect(slice2)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	expected := []interface{}{3, 4}
	if !equalSlices(result.Value().([]interface{}), expected) {
		t.Errorf("Expected %v, got %v", expected, result.Value())
	}

	slice1 = prototipe.NewPrototype([]interface{}{1, 2})
	slice2 = prototipe.NewPrototype([]interface{}{3, 4})
	result, err = slice1.Intersect(slice2)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	expected = []interface{}{}
	if !equalSlices(result.Value().([]interface{}), expected) {
		t.Errorf("Expected %v, got %v", expected, result.Value())
	}

	notSlice := prototipe.NewPrototype(42)
	_, err = slice1.Intersect(notSlice)
	if err == nil {
		t.Errorf("Expected an error, got nil")
	}
}

func TestZip(t *testing.T) {
	slice1 := prototipe.NewPrototype([]interface{}{1, 2, 3})
	slice2 := prototipe.NewPrototype([]interface{}{"a", "b", "c"})
	result, err := slice1.Zip(slice2)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	expected := [][2]interface{}{
		{1, "a"},
		{2, "b"},
		{3, "c"},
	}
	if !reflect.DeepEqual(result.Value(), expected) {
		t.Errorf("Expected %v, got %v", expected, result.Value())
	}

	slice1 = prototipe.NewPrototype([]interface{}{1, 2})
	slice2 = prototipe.NewPrototype([]interface{}{"a", "b", "c"})
	result, err = slice1.Zip(slice2)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	expected = [][2]interface{}{
		{1, "a"},
		{2, "b"},
	}
	if !reflect.DeepEqual(result.Value(), expected) {
		t.Errorf("Expected %v, got %v", expected, result.Value())
	}

	notSlice := prototipe.NewPrototype(42)
	_, err = slice1.Zip(notSlice)
	if err == nil {
		t.Errorf("Expected an error, got nil")
	}
}

func equalSlices(a, b []interface{}) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
