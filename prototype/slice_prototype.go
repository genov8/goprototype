package prototype

import (
	"fmt"
	"reflect"
)

type SlicePrototype struct {
	*Prototype
}

func NewSlicePrototype(v interface{}) *SlicePrototype {
	if reflect.TypeOf(v).Kind() != reflect.Slice {
		return &SlicePrototype{
			Prototype: &Prototype{
				value: nil,
				err:   fmt.Errorf("invalid type for SlicePrototype: must be a slice"),
			},
		}
	}
	return &SlicePrototype{Prototype: &Prototype{value: v}}
}
