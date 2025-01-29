package prototype

import "errors"

type NumberPrototype struct {
	*Prototype
}

func NewNumberPrototype(v interface{}) *NumberPrototype {
	switch v.(type) {
	case int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64,
		float32, float64:
		return &NumberPrototype{Prototype: &Prototype{value: v}}
	default:
		return &NumberPrototype{Prototype: &Prototype{value: 0, err: errors.New("invalid type for NumberPrototype")}}
	}
}
