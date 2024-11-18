package prototipe

type Prototype struct {
	value interface{}
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
