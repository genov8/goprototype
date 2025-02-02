package prototype

import (
	"errors"
	"reflect"
)



func (p *Prototype) IsEmpty() (*Prototype, error) {
	switch v := p.value.(type) {
	case string:
		return &Prototype{value: v == ""}, nil
	case []interface{}:
		return &Prototype{value: len(v) == 0}, nil
	case int:
		return &Prototype{value: v == 0}, nil
	case float64:
		return &Prototype{value: v == 0}, nil
	default:
		return nil, errors.New("unsupported type for IsEmpty")
	}
}

