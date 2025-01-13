package prototipe

import (
	"errors"
	"strings"
)

func (p *Prototype) ToUpperCase() (*Prototype, error) {
	return p.processString(strings.ToUpper)
}

func (p *Prototype) ToLowerCase() (*Prototype, error) {
	return p.processString(strings.ToLower)
}

func (p *Prototype) Capitalize() (*Prototype, error) {
	return p.processString(func(str string) string {
		if len(str) == 0 {
			return str
		}
		firstSymbol := str[:1]
		lastStr := str[1:]
		return strings.ToUpper(firstSymbol) + strings.ToLower(lastStr)
	})
}

func (p *Prototype) Trim(cutset string) (*Prototype, error) {
	return p.processString(func(str string) string {
		return strings.Trim(str, cutset)
	})
}

func (p *Prototype) Split(separator string) (*Prototype, error) {
	if str, ok := p.value.(string); ok {
		return &Prototype{value: strings.Split(str, separator)}, nil
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

func (p *Prototype) EndsWith(prefix string) (*Prototype, error) {
	if str, ok := p.value.(string); ok {
		result := strings.HasSuffix(str, prefix)
		return &Prototype{value: result}, nil
	}
	return nil, errors.New("value is not a string")
}

func (p *Prototype) RemoveWhitespace() (*Prototype, error) {
	return p.processString(func(str string) string {
		return strings.ReplaceAll(str, " ", "")
	})
}

func (p *Prototype) Replace(old, new string) (*Prototype, error) {
	return p.processString(func(str string) string {
		return strings.ReplaceAll(str, old, new)
	})
}

func (p *Prototype) Repeat(count int) (*Prototype, error) {
	if str, ok := p.value.(string); ok {
		if count < 0 {
			return nil, errors.New("repeat count must be non-negative")
		}
		return &Prototype{value: strings.Repeat(str, count)}, nil
	}
	return nil, errors.New("value is not a string")
}

func (p *Prototype) ReverseWords() (*Prototype, error) {
	if str, ok := p.value.(string); ok {
		words := strings.Fields(str)
		for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
			words[i], words[j] = words[j], words[i]
		}
		return &Prototype{value: strings.Join(words, " ")}, nil
	}
	return nil, errors.New("value is not a string")
}

func (p *Prototype) processString(operation func(string) string) (*Prototype, error) {
	if str, ok := p.value.(string); ok {
		return &Prototype{value: operation(str)}, nil
	}
	return nil, errors.New("value is not a string")
}
