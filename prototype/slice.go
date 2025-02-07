package prototype

import (
	"errors"
	"fmt"
	"sort"
)

func (p *SlicePrototype) Append(element interface{}) (*SlicePrototype, error) {
	return p.processSlice(func(slice []interface{}) ([]interface{}, error) {
		return append(slice, element), nil
	})
}

func (p *SlicePrototype) Remove(index int) (*SlicePrototype, error) {
	return p.processSlice(func(slice []interface{}) ([]interface{}, error) {
		if index < 0 || index >= len(slice) {
			return nil, errors.New("index out of bounds")
		}
		return append(slice[:index], slice[index+1:]...), nil
	})
}

func (p *Prototype) PrintSlice() (*Prototype, error) {
	if arr, ok := p.value.([]interface{}); ok {
		result := fmt.Sprintf("%v", arr)
		return &Prototype{value: result}, nil
	}
	return nil, errors.New("value is not a slice")
}

func (p *SlicePrototype) Unique() (*SlicePrototype, error) {
	return p.processSlice(func(slice []interface{}) ([]interface{}, error) {
		seen := make(map[interface{}]bool)
		var uniqueSlice []interface{}

		for _, elem := range slice {
			if !seen[elem] {
				seen[elem] = true
				uniqueSlice = append(uniqueSlice, elem)
			}
		}
		return uniqueSlice, nil
	})
}

func (p *SlicePrototype) IndexOf(element interface{}) (*SlicePrototype, error) {
	if slice, ok := p.value.([]interface{}); ok {
		for i, v := range slice {
			if v == element {
				return &SlicePrototype{&Prototype{value: i}}, nil
			}
		}
		return &SlicePrototype{&Prototype{value: -1}}, nil
	}
	return nil, errors.New("value is not a slice")
}

func (p *SlicePrototype) Flatten() (*SlicePrototype, error) {
	if slice, ok := p.value.([]interface{}); ok {
		var flatSlice []interface{}
		stack := slice
		for len(stack) > 0 {
			item := stack[0]
			stack = stack[1:]

			if nestedSlice, ok := item.([]interface{}); ok {
				stack = append(nestedSlice, stack...)
			} else {
				flatSlice = append(flatSlice, item)
			}
		}
		return &SlicePrototype{&Prototype{value: flatSlice}}, nil
	}
	return nil, errors.New("value is not a slice")
}

func (p *SlicePrototype) Map(operation func(interface{}) interface{}) (*SlicePrototype, error) {
	if slice, ok := p.value.([]interface{}); ok {
		mappedSlice := make([]interface{}, len(slice))
		for i, v := range slice {
			mappedSlice[i] = operation(v)
		}
		return &SlicePrototype{&Prototype{value: mappedSlice}}, nil
	}
	return nil, errors.New("value is not a slice")
}

func (p *SlicePrototype) Filter(predicate func(interface{}) bool) (*SlicePrototype, error) {
	if slice, ok := p.value.([]interface{}); ok {
		var filteredSlice []interface{}
		for _, v := range slice {
			if predicate(v) {
				filteredSlice = append(filteredSlice, v)
			}
		}
		return &SlicePrototype{&Prototype{value: filteredSlice}}, nil
	}
	return nil, errors.New("value is not a slice")
}

func (p *SlicePrototype) Chunk(size int) (*SlicePrototype, error) {
	if size <= 0 {
		return nil, errors.New("chunk size must be greater than 0")
	}

	if slice, ok := p.value.([]interface{}); ok {
		var chunks [][]interface{}
		for i := 0; i < len(slice); i += size {
			end := i + size
			if end > len(slice) {
				end = len(slice)
			}
			chunks = append(chunks, slice[i:end])
		}
		return &SlicePrototype{&Prototype{value: chunks}}, nil
	}
	return nil, errors.New("value is not a slice")
}

func (p *SlicePrototype) Intersect(other *SlicePrototype) (*SlicePrototype, error) {
	slice1, ok1 := p.value.([]interface{})
	slice2, ok2 := other.value.([]interface{})
	if !ok1 || !ok2 {
		return nil, errors.New("both values must be slices")
	}

	hashMap := make(map[interface{}]bool)
	for _, item := range slice1 {
		hashMap[item] = true
	}

	var intersection []interface{}
	for _, item := range slice2 {
		if hashMap[item] {
			intersection = append(intersection, item)
		}
	}

	return &SlicePrototype{&Prototype{value: intersection}}, nil
}

func (p *SlicePrototype) Zip(other *SlicePrototype) (*SlicePrototype, error) {
	slice1, ok1 := p.value.([]interface{})
	slice2, ok2 := other.value.([]interface{})
	if !ok1 || !ok2 {
		return nil, errors.New("both values must be slices")
	}

	length := len(slice1)
	if len(slice2) < length {
		length = len(slice2)
	}

	zipped := make([][2]interface{}, length)
	for i := 0; i < length; i++ {
		zipped[i] = [2]interface{}{slice1[i], slice2[i]}
	}

	return &SlicePrototype{&Prototype{value: zipped}}, nil
}

func (p *SlicePrototype) Sort(customComparator func(a, b interface{}) bool) (*SlicePrototype, error) {
	if slice, ok := p.value.([]interface{}); ok {
		if customComparator == nil {
			if isHomogeneous(slice, "int") {
				sort.Slice(slice, func(i, j int) bool {
					return slice[i].(int) < slice[j].(int)
				})
			} else if isHomogeneous(slice, "string") {
				sort.Slice(slice, func(i, j int) bool {
					return slice[i].(string) < slice[j].(string)
				})
			} else {
				return nil, errors.New("slice must contain either all numbers or all strings for default sorting")
			}
		} else {
			sort.Slice(slice, func(i, j int) bool {
				return customComparator(slice[i], slice[j])
			})
		}

		return &SlicePrototype{&Prototype{value: slice}}, nil
	}

	return nil, errors.New("value is not a slice")
}

func isHomogeneous(slice []interface{}, kind string) bool {
	for _, v := range slice {
		switch kind {
		case "int":
			if _, ok := v.(int); !ok {
				return false
			}
		case "string":
			if _, ok := v.(string); !ok {
				return false
			}
		}
	}
	return true
}

func (p *SlicePrototype) Rotate(steps int) (*SlicePrototype, error) {
	slice, ok := p.value.([]interface{})
	if !ok {
		return nil, errors.New("value is not a slice")
	}

	n := len(slice)
	if n == 0 || steps == 0 {
		return &SlicePrototype{&Prototype{value: slice}}, nil
	}

	steps = ((steps % n) + n) % n

	rotated := append(slice[n-steps:], slice[:n-steps]...)

	return &SlicePrototype{&Prototype{value: rotated}}, nil
}

func (p *SlicePrototype) processSlice(operation func([]interface{}) ([]interface{}, error)) (*SlicePrototype, error) {
	if slice, ok := p.value.([]interface{}); ok {
		result, err := operation(slice)
		if err != nil {
			return nil, err
		}
		return &SlicePrototype{&Prototype{value: result}}, nil
	}
	return nil, errors.New("value is not a slice")
}
