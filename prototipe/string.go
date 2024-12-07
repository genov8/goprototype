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

func (p *Prototype) Trim(cutset string) (*Prototype, error) {
	if str, ok := p.value.(string); ok {
		return &Prototype{value: strings.Trim(str, cutset)}, nil
	}
	return nil, errors.New("value is not a string")
}

func (p *Prototype) Split(separator string) (*Prototype, error) {
	if str, ok := p.value.(string); ok {
		return &Prototype{value: strings.Split(str, separator)}, nil
	}
	return nil, errors.New("value is not a string")
}

func (p *Prototype) Reverse() (*Prototype, error) {
	if str, ok := p.value.(string); ok {
		runes := []rune(str)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		return &Prototype{value: string(runes)}, nil
	}
	return nil, errors.New("value is not a string")
}

func (p *Prototype) StartsWith(prefix string) (*Prototype, error) {
	if str, ok := p.value.(string); ok {
		result := strings.HasPrefix(str, prefix)
		return &Prototype{value: result}, nil
	}
	return nil, errors.New("value is not a string")
}
