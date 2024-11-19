package prototipe

import (
	"errors"
	"strings"
)

func (p *Prototype) Concat(s string) (*Prototype, error) {
	if str, ok := p.value.(string); ok {
		return &Prototype{value: str + s}, nil
	}
	return nil, errors.New("value is not a string")
}

func (p *Prototype) ToUpper() (*Prototype, error) {
	if str, ok := p.value.(string); ok {
		return &Prototype{value: strings.ToUpper(str)}, nil
	}
	return nil, errors.New("value is not a string")
}
