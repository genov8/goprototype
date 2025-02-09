package prototype

import (
	"errors"
	"math"
)

func (p *NumberPrototype) Add(n float64) (*NumberPrototype, error) {
	switch v := p.value.(type) {
	case int:
		return &NumberPrototype{&Prototype{value: v + int(n)}}, nil
	case float64:
		return &NumberPrototype{&Prototype{value: v + n}}, nil
	default:
		return nil, errors.New("value is not a number")
	}
}

func (p *NumberPrototype) Multiply(n float64) (*NumberPrototype, error) {
	switch v := p.value.(type) {
	case int:
		return &NumberPrototype{&Prototype{value: v * int(n)}}, nil
	case float64:
		return &NumberPrototype{&Prototype{value: v * n}}, nil
	default:
		return nil, errors.New("value is not a number")
	}
}

func (p *NumberPrototype) Divide(divisor float64) (*NumberPrototype, error) {
	switch value := p.value.(type) {
	case int:
		if divisor == 0 {
			return nil, errors.New("division by zero is not allowed")
		}
		return &NumberPrototype{&Prototype{value: float64(value) / divisor}}, nil
	case float64:
		if divisor == 0 {
			return nil, errors.New("division by zero is not allowed")
		}
		return &NumberPrototype{&Prototype{value: value / divisor}}, nil
	default:
		return nil, errors.New("value is not a number")
	}
}

func (p *NumberPrototype) IsEven() (*NumberPrototype, error) {
	if err := p.checkError(); err != nil {
		return nil, err
	}

	switch value := p.value.(type) {
	case int:
		return &NumberPrototype{&Prototype{value: value%2 == 0}}, nil
	case float64:
		if value == float64(int(value)) {
			return &NumberPrototype{&Prototype{value: int(value)%2 == 0}}, nil
		}
		return nil, errors.New("value is not an integer")
	default:
		return nil, errors.New("value is not a number")
	}
}

func (p *NumberPrototype) IsOdd() (*NumberPrototype, error) {
	if err := p.checkError(); err != nil {
		return nil, err
	}

	switch value := p.value.(type) {
	case int:
		return &NumberPrototype{&Prototype{value: value%2 != 0}}, nil
	case float64:
		if value == float64(int(value)) {
			return &NumberPrototype{&Prototype{value: int(value)%2 != 0}}, nil
		}
		return nil, errors.New("value is not an integer")
	default:
		return nil, errors.New("value is not a number")
	}
}

func (p *NumberPrototype) Modulo(divisor int) (*NumberPrototype, error) {
	if err := p.checkError(); err != nil {
		return nil, err
	}
	switch value := p.value.(type) {
	case int:
		if divisor == 0 {
			return nil, errors.New("division by zero is not allowed")
		}
		return &NumberPrototype{&Prototype{value: value % divisor}}, nil
	case float64:
		if divisor == 0 {
			return nil, errors.New("division by zero is not allowed")
		}
		return nil, errors.New("modulo is not supported for float numbers")
	default:
		return nil, errors.New("value is not an number")
	}
}

func (p *NumberPrototype) Power(exp float64) (*NumberPrototype, error) {
	if err := p.checkError(); err != nil {
		return nil, err
	}
	switch value := p.value.(type) {
	case int:
		result := math.Pow(float64(value), exp)
		return &NumberPrototype{&Prototype{value: result}}, nil
	case float64:
		result := math.Pow(value, exp)
		return &NumberPrototype{&Prototype{value: result}}, nil
	default:
		return nil, errors.New("value is not a number")
	}
}

func (p *NumberPrototype) Subtract(value float64) (*NumberPrototype, error) {
	if err := p.checkError(); err != nil {
		return nil, err
	}
	switch current := p.value.(type) {
	case int:
		return &NumberPrototype{&Prototype{value: current - int(value)}}, nil
	case float64:
		return &NumberPrototype{&Prototype{value: current - value}}, nil
	default:
		return nil, errors.New("value is not a number")
	}
}

func (p *NumberPrototype) Abs() (*NumberPrototype, error) {
	if err := p.checkError(); err != nil {
		return nil, err
	}
	switch v := p.value.(type) {
	case int:
		return &NumberPrototype{&Prototype{value: int(math.Abs(float64(v)))}}, nil
	case float64:
		return &NumberPrototype{&Prototype{value: math.Abs(v)}}, nil
	default:
		return nil, errors.New("value is not a number")
	}
}

func (p *NumberPrototype) Round() (*NumberPrototype, error) {
	if err := p.checkError(); err != nil {
		return nil, err
	}
	switch v := p.value.(type) {
	case int:
		return &NumberPrototype{&Prototype{value: v}}, nil
	case float64:
		return &NumberPrototype{&Prototype{value: math.Round(v)}}, nil
	default:
		return nil, errors.New("value is not a number")
	}
}

func (p *NumberPrototype) Sqrt() (*NumberPrototype, error) {
	if err := p.checkError(); err != nil {
		return nil, err
	}
	switch v := p.value.(type) {
	case int:
		if v < 0 {
			return nil, errors.New("cannot calculate square root of a negative number")
		}
		return &NumberPrototype{&Prototype{value: math.Sqrt(float64(v))}}, nil
	case float64:
		if v < 0 {
			return nil, errors.New("cannot calculate square root of a negative number")
		}
		return &NumberPrototype{&Prototype{value: math.Sqrt(v)}}, nil
	default:
		return nil, errors.New("value is not a number")
	}
}

func (p *NumberPrototype) Factorial() (*NumberPrototype, error) {
	if err := p.checkError(); err != nil {
		return nil, err
	}
	if num, ok := p.value.(int); ok {
		if num < 0 {
			return nil, errors.New("factorial is not defined for negative numbers")
		}
		result := 1
		for i := 1; i <= num; i++ {
			result *= i
		}
		return &NumberPrototype{&Prototype{value: result}}, nil
	}
	return nil, errors.New("value is not an integer")
}

func (p *NumberPrototype) IsPrime() (*NumberPrototype, error) {
	if err := p.checkError(); err != nil {
		return nil, err
	}
	num, ok := p.value.(int)
	if !ok {
		return nil, errors.New("value is not an integer")
	}

	if num <= 1 {
		return &NumberPrototype{&Prototype{value: false}}, nil
	}
	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return &NumberPrototype{&Prototype{value: false}}, nil
		}
	}
	return &NumberPrototype{&Prototype{value: true}}, nil
}

func (p *NumberPrototype) Clamp(min, max float64) (*NumberPrototype, error) {
	if err := p.checkError(); err != nil {
		return nil, err
	}
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
		return &NumberPrototype{&Prototype{value: clampedValue}}, nil
	}
	return nil, errors.New("value is not a number")
}

func (p *NumberPrototype) Fibonacci() (*NumberPrototype, error) {
	if err := p.checkError(); err != nil {
		return nil, err
	}
	num, ok := p.value.(int)
	if !ok || num < 0 {
		return nil, errors.New("value must be a non-negative integer")
	}

	if num == 0 {
		return &NumberPrototype{&Prototype{value: []int{0}}}, nil
	}
	if num == 1 {
		return &NumberPrototype{&Prototype{value: []int{0, 1}}}, nil
	}

	fib := []int{0, 1}
	for len(fib) < num {
		next := fib[len(fib)-1] + fib[len(fib)-2]
		fib = append(fib, next)
	}

	return &NumberPrototype{&Prototype{value: fib}}, nil
}

func (p *NumberPrototype) IsPowerOfTwo() (*NumberPrototype, error) {
	if err := p.checkError(); err != nil {
		return nil, err
	}
	num, ok := p.value.(int)
	if !ok || num <= 0 {
		return nil, errors.New("value must be a positive integer")
	}

	isPower := (num & (num - 1)) == 0
	return &NumberPrototype{&Prototype{value: isPower}}, nil
}

func (p *NumberPrototype) checkError() error {
	if p.err != nil {
		return p.err
	}
	return nil
}
