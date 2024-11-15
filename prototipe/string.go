package prototipe

import "errors"

func (p *Prototype) Length() (int, error) {
	if str, ok := p.value.(string); ok {
		return len(str), nil
	}
	return 0, errors.New("value is not a string")
}

func (p *Prototype) Concat(s string) (*Prototype, error) {
	if str, ok := p.value.(string); ok {
		return &Prototype{value: str + s}, nil
	}
	return nil, errors.New("value is not a string")
}
