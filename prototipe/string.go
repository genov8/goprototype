package prototipe

import "errors"

func (p *Prototype) Concat(s string) (*Prototype, error) {
	if str, ok := p.value.(string); ok {
		return &Prototype{value: str + s}, nil
	}
	return nil, errors.New("value is not a string")
}
