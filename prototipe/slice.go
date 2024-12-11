package prototipe

import (
	"errors"
	"fmt"
)

func (p *Prototype) Append(element interface{}) (*Prototype, error) {
	if arr, ok := p.value.([]interface{}); ok {
		newArr := append(arr, element)
		return &Prototype{value: newArr}, nil
	}
	return nil, errors.New("value is not an array")
}

func (p *Prototype) Remove(index int) (*Prototype, error) {
	if arr, ok := p.value.([]interface{}); ok {
		if index < 0 || index >= len(arr) {
			return nil, errors.New("index out of bounds")
		}
		newArr := append(arr[:index], arr[index+1:]...)
		return &Prototype{value: newArr}, nil
	}
	return nil, errors.New("value is not an array")
}

func (p *Prototype) PrintSlice() (*Prototype, error) {
	if arr, ok := p.value.([]interface{}); ok {
		result := fmt.Sprintf("%v", arr)
		return &Prototype{value: result}, nil
	}
	return nil, errors.New("value is not a slice")
}

func (p *Prototype) Unique() (*Prototype, error) {
	if slice, ok := p.value.([]interface{}); ok {
		seen := make(map[interface{}]bool)
		var uniqueSlice []interface{}

		for _, elem := range slice {
			if !seen[elem] {
				seen[elem] = true
				uniqueSlice = append(uniqueSlice, elem)
			}
		}

		return &Prototype{value: uniqueSlice}, nil
	}
	return nil, errors.New("value is not a slice")
}
