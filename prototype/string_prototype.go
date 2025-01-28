package prototype

type StringPrototype struct {
	*Prototype
}

func NewStringPrototype(v string) *StringPrototype {
	return &StringPrototype{
		Prototype: &Prototype{value: v},
	}
}
