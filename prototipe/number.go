package prototipe

import (
	"errors"
	"math"
)

func (p *Prototype) Add(n float64) (*Prototype, error) {
	switch v := p.value.(type) {
	case int:
		return &Prototype{value: v + int(n)}, nil
	case float64:
		return &Prototype{value: v + n}, nil
	default:
		return nil, errors.New("value is not a number")
	}
}

func (p *Prototype) Multiply(n float64) (*Prototype, error) {
	switch v := p.value.(type) {
	case int:
		return &Prototype{value: v * int(n)}, nil
	case float64:
		return &Prototype{value: v * n}, nil
	default:
		return nil, errors.New("value is not a number")
	}
}

func (p *Prototype) Divide(divisor float64) (*Prototype, error) {
	switch value := p.value.(type) {
	case int:
		if divisor == 0 {
			return nil, errors.New("division by zero is not allowed")
		}
		return &Prototype{value: float64(value) / divisor}, nil
	case float64:
		if divisor == 0 {
			return nil, errors.New("division by zero is not allowed")
		}
		return &Prototype{value: value / divisor}, nil
	default:
		return nil, errors.New("value is not a number")
	}
}

func (p *Prototype) IsEven() (*Prototype, error) {
	switch value := p.value.(type) {
	case int:
		return &Prototype{value: value%2 == 0}, nil
	case float64:
		if value == float64(int(value)) {
			return &Prototype{value: int(value)%2 == 0}, nil
		}
		return nil, errors.New("value is not an integer")
	default:
		return nil, errors.New("value is not a number")
	}
}

func (p *Prototype) IsOdd() (*Prototype, error) {
	switch value := p.value.(type) {
	case int:
		return &Prototype{value: value%2 != 0}, nil
	case float64:
		if value == float64(int(value)) {
			return &Prototype{value: int(value)%2 != 0}, nil
		}
		return nil, errors.New("value is not an integer")
	default:
		return nil, errors.New("value is not a number")
	}
}

func (p *Prototype) Modulo(divisor int) (*Prototype, error) {
	switch value := p.value.(type) {
	case int:
		if divisor == 0 {
			return nil, errors.New("division by zero is not allowed")
		}
		return &Prototype{value: value % divisor}, nil
	case float64:
		if divisor == 0 {
			return nil, errors.New("division by zero is not allowed")
		}
		return nil, errors.New("modulo is not supported for float numbers")
	default:
		return nil, errors.New("value is not an number")
	}
}

func (p *Prototype) Power(exp float64) (*Prototype, error) {
	switch value := p.value.(type) {
	case int:
		result := math.Pow(float64(value), exp)
		return &Prototype{value: result}, nil
	case float64:
		result := math.Pow(value, exp)
		return &Prototype{value: result}, nil
	default:
		return nil, errors.New("value is not a number")
	}
}

func (p *Prototype) Subtract(value float64) (*Prototype, error) {
	switch current := p.value.(type) {
	case int:
		return &Prototype{value: current - int(value)}, nil
	case float64:
		return &Prototype{value: current - value}, nil
	default:
		return nil, errors.New("value is not a number")
	}
}
