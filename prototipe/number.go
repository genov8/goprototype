package prototipe

import (
	"errors"
)

func (p *Prototype) Add(n float64) (*Prototype, error) {
	switch v := p.value.(type) {
	case int:
		return &Prototype{value: v + int(n)}, nil
	case float64:
		return &Prototype{value: v + n}, nil
	default:
		return nil, errors.New("value is not a number")
	}
}

func (p *Prototype) Multiply(n float64) (*Prototype, error) {
	switch v := p.value.(type) {
	case int:
		return &Prototype{value: v * int(n)}, nil
	case float64:
		return &Prototype{value: v * n}, nil
	default:
		return nil, errors.New("value is not a number")
	}
}
