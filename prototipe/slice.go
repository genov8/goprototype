package prototipe

import (
	"errors"
	"fmt"
)

func (p *Prototype) Append(element interface{}) (*Prototype, error) {
	return p.processSlice(func(slice []interface{}) ([]interface{}, error) {
		return append(slice, element), nil
	})
}

func (p *Prototype) Remove(index int) (*Prototype, error) {
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

func (p *Prototype) Unique() (*Prototype, error) {
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

func (p *Prototype) IndexOf(element interface{}) (*Prototype, error) {
	if slice, ok := p.value.([]interface{}); ok {
		for i, v := range slice {
			if v == element {
				return &Prototype{value: i}, nil
			}
		}
		return &Prototype{value: -1}, nil
	}
	return nil, errors.New("value is not a slice")
}

func (p *Prototype) Flatten() (*Prototype, error) {
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
		return &Prototype{value: flatSlice}, nil
	}
	return nil, errors.New("value is not a slice")
}

func (p *Prototype) Map(operation func(interface{}) interface{}) (*Prototype, error) {
	if slice, ok := p.value.([]interface{}); ok {
		mappedSlice := make([]interface{}, len(slice))
		for i, v := range slice {
			mappedSlice[i] = operation(v)
		}
		return &Prototype{value: mappedSlice}, nil
	}
	return nil, errors.New("value is not a slice")
}

func (p *Prototype) Filter(predicate func(interface{}) bool) (*Prototype, error) {
	if slice, ok := p.value.([]interface{}); ok {
		var filteredSlice []interface{}
		for _, v := range slice {
			if predicate(v) {
				filteredSlice = append(filteredSlice, v)
			}
		}
		return &Prototype{value: filteredSlice}, nil
	}
	return nil, errors.New("value is not a slice")
}

func (p *Prototype) Chunk(size int) (*Prototype, error) {
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
		return &Prototype{value: chunks}, nil
	}
	return nil, errors.New("value is not a slice")
}

func (p *Prototype) Intersect(other *Prototype) (*Prototype, error) {
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

	return &Prototype{value: intersection}, nil
}

func (p *Prototype) Zip(other *Prototype) (*Prototype, error) {
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

	return &Prototype{value: zipped}, nil
}

func (p *Prototype) processSlice(operation func([]interface{}) ([]interface{}, error)) (*Prototype, error) {
	if slice, ok := p.value.([]interface{}); ok {
		result, err := operation(slice)
		if err != nil {
			return nil, err
		}
		return &Prototype{value: result}, nil
	}
	return nil, errors.New("value is not a slice")
}
