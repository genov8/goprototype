package prototype

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

func (p *Prototype) Abs() (*Prototype, error) {
	switch v := p.value.(type) {
	case int:
		return &Prototype{value: int(math.Abs(float64(v)))}, nil
	case float64:
		return &Prototype{value: math.Abs(v)}, nil
	default:
		return nil, errors.New("value is not a number")
	}
}

func (p *Prototype) Round() (*Prototype, error) {
	switch v := p.value.(type) {
	case int:
		return &Prototype{value: v}, nil
	case float64:
		return &Prototype{value: math.Round(v)}, nil
	default:
		return nil, errors.New("value is not a number")
	}
}

func (p *Prototype) Sqrt() (*Prototype, error) {
	switch v := p.value.(type) {
	case int:
		if v < 0 {
			return nil, errors.New("cannot calculate square root of a negative number")
		}
		return &Prototype{value: math.Sqrt(float64(v))}, nil
	case float64:
		if v < 0 {
			return nil, errors.New("cannot calculate square root of a negative number")
		}
		return &Prototype{value: math.Sqrt(v)}, nil
	default:
		return nil, errors.New("value is not a number")
	}
}

func (p *Prototype) Factorial() (*Prototype, error) {
	if num, ok := p.value.(int); ok {
		if num < 0 {
			return nil, errors.New("factorial is not defined for negative numbers")
		}
		result := 1
		for i := 1; i <= num; i++ {
			result *= i
		}
		return &Prototype{value: result}, nil
	}
	return nil, errors.New("value is not an integer")
}

func (p *Prototype) IsPrime() (*Prototype, error) {
	num, ok := p.value.(int)
	if !ok {
		return nil, errors.New("value is not an integer")
	}

	if num <= 1 {
		return &Prototype{value: false}, nil
	}
	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return &Prototype{value: false}, nil
		}
	}
	return &Prototype{value: true}, nil
}

func (p *Prototype) Clamp(min, max float64) (*Prototype, error) {
	if num, ok := p.value.(float64); ok {
		if min > max {
			return nil, errors.New("min cannot be greater than max")
		}
		clampedValue := num
		if num < min {
			clampedValue = min
		} else if num > max {
			clampedValue = max
		}
		return &Prototype{value: clampedValue}, nil
	}
	return nil, errors.New("value is not a number")
}

func (p *Prototype) Fibonacci() (*Prototype, error) {
	num, ok := p.value.(int)
	if !ok || num < 0 {
		return nil, errors.New("value must be a non-negative integer")
	}

	if num == 0 {
		return &Prototype{value: []int{0}}, nil
	}
	if num == 1 {
		return &Prototype{value: []int{0, 1}}, nil
	}

	fib := []int{0, 1}
	for len(fib) < num {
		next := fib[len(fib)-1] + fib[len(fib)-2]
		fib = append(fib, next)
	}

	return &Prototype{value: fib}, nil
}

func (p *Prototype) IsPowerOfTwo() (*Prototype, error) {
	num, ok := p.value.(int)
	if !ok || num <= 0 {
		return nil, errors.New("value must be a positive integer")
	}

	isPower := (num & (num - 1)) == 0
	return &Prototype{value: isPower}, nil
}
