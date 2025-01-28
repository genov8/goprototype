package prototype

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
