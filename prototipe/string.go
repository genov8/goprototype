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

func (p *Prototype) ToUpperCase() (*Prototype, error) {
	if str, ok := p.value.(string); ok {
		return &Prototype{value: strings.ToUpper(str)}, nil
	}
	return nil, errors.New("value is not a string")
}

func (p *Prototype) ToLowerCase() (*Prototype, error) {
	if str, ok := p.value.(string); ok {
		return &Prototype{value: strings.ToLower(str)}, nil
	}
	return nil, errors.New("value is not a string")
}

func (p *Prototype) Capitalize() (*Prototype, error) {
	if str, ok := p.value.(string); ok {
		firstSymbol := str[:1]
		lastStr := str[1:]
		newStr := strings.ToUpper(firstSymbol) + strings.ToLower(lastStr)
		return &Prototype{value: newStr}, nil
	}
	return nil, errors.New("value is not a string")
}
