package prototype

import (
	"fmt"
	"reflect"
)

type StringPrototype struct {
	*Prototype
}

func NewStringPrototype(v interface{}) *StringPrototype {
	if reflect.TypeOf(v).Kind() != reflect.String {
		return &StringPrototype{
			Prototype: &Prototype{
				value: nil,
				err:   fmt.Errorf("invalid type for StringPrototype: must be a slice"),
			},
		}
	}
	return &StringPrototype{Prototype: &Prototype{value: v}}
}
