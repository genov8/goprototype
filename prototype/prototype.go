package prototype

import "errors"

type Prototype struct {
	value interface{}
	err   error
}

func NewPrototype(v interface{}) *Prototype {
	return &Prototype{value: v}
}

func (p *Prototype) Value() interface{} {
	return p.value
}

func (p *Prototype) SetValue(v interface{}) {
	p.value = v
}

type Chain struct {
	Prototype *Prototype
	Err       error
}

func NewChain(p *Prototype) *Chain {
	return &Chain{Prototype: p}
}

func (c *Chain) Must() *Prototype {
	if c.Err != nil {
		panic(c.Err)
	}
	return c.Prototype
}

func (c *Chain) Invoke(method func(*Prototype) (*Prototype, error)) *Chain {
	if c.Err != nil {
		return c
	}
	c.Prototype, c.Err = method(c.Prototype)
	return c
}

func (p *Prototype) Error() error {
	return p.err
}

func (p *Prototype) SetError(err error) {
	p.err = err
}

func (p *Prototype) Length() (int, error) {
	switch v := p.value.(type) {
	case string:
		return len(v), nil
	case []interface{}:
		return len(v), nil
	default:
		return 0, errors.New("value is neither a string nor an array")
	}
}

func (p *Prototype) Reverse() (*Prototype, error) {
	switch value := p.value.(type) {
	case string:
		return &Prototype{value: reverseString(value)}, nil
	case []interface{}:
		return &Prototype{value: reverseSlice(value)}, nil
	default:
		return nil, errors.New("value is not a string or slice")
	}
}

func (p *Prototype) Concat(other *Prototype) (*Prototype, error) {
	switch value1 := p.value.(type) {
	case string:
		if value2, ok := other.value.(string); ok {
			return &Prototype{value: value1 + value2}, nil
		}
		return nil, errors.New("second value is not a string")
	case []interface{}:
		if value2, ok := other.value.([]interface{}); ok {
			return &Prototype{value: append(value1, value2...)}, nil
		}
		return nil, errors.New("second value is not a slice")
	default:
		return nil, errors.New("value is not a string or slice")
	}
}

func (p *Prototype) Contains(element interface{}) (*Prototype, error) {
	switch value := p.value.(type) {
	case string:
		if substr, ok := element.(string); ok {
			return &Prototype{value: strings.Contains(value, substr)}, nil
		}
		return nil, errors.New("element must be a string for string value")
	case []interface{}:
		for _, v := range value {
			if v == element {
				return &Prototype{value: true}, nil
			}
		}
		return &Prototype{value: false}, nil
	default:
		return nil, errors.New("value is not a string or slice")
	}
}

func reverseString(input string) string {
	runes := []rune(input)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func reverseSlice(input []interface{}) []interface{} {
	length := len(input)
	reversed := make([]interface{}, length)
	for i, v := range input {
		reversed[length-1-i] = v
	}
	return reversed
}

