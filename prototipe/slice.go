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

func (p *Prototype) PrintSlice() (string, error) {
	if arr, ok := p.value.([]interface{}); ok {
		return fmt.Sprintf("%v", arr), nil
	}
	return "", errors.New("value is not an array")
}
