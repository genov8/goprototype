package prototipe

import "errors"

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
