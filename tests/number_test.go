package tests

import (
	"github.com/genov8/goprototipe/prototipe"
	"testing"
)

func TestAdd(t *testing.T) {
	num := prototipe.NewPrototype(10)
	newNum, err := num.Add(5)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if newNum.Value() != 15 {
		t.Errorf("Expected 15, got %v", newNum.Value())
	}
}

func TestMultiply(t *testing.T) {
	num := prototipe.NewPrototype(10)
	newNum, err := num.Multiply(3)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if newNum.Value() != 30 {
		t.Errorf("Expected 30, got %v", newNum.Value())
	}
}

func TestNumber(t *testing.T) {
	numProto := prototipe.NewPrototype(10)
	resultProto, err := numProto.Divide(2)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if resultProto.Value() != 5.0 {
		t.Errorf("Expected 5, got %v", resultProto.Value())
	}

	floatProto := prototipe.NewPrototype(15.5)
	resultProto, err = floatProto.Divide(3.1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if resultProto.Value() != 15.5/3.1 {
		t.Errorf("Expected %v, got %v", 15.5/3.1, resultProto.Value())
	}

	_, err = numProto.Divide(0)
	if err == nil {
		t.Error("Expected division by zero error, got nil")
	}

	strProto := prototipe.NewPrototype("hello")
	_, err = strProto.Divide(2)
	if err == nil {
		t.Error("Expected an error for non-number value, got nil")
	}
}

func TestIsEven(t *testing.T) {
	intProto := prototipe.NewPrototype(10)
	resultProto, err := intProto.IsEven()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if resultProto.Value() != true {
		t.Errorf("Expected true, got %v", resultProto.Value())
	}

	oddProto := prototipe.NewPrototype(7)
	resultProto, err = oddProto.IsEven()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if resultProto.Value() != false {
		t.Errorf("Expected false, got %v", resultProto.Value())
	}

	floatProto := prototipe.NewPrototype(10.5)
	_, err = floatProto.IsEven()
	if err == nil {
		t.Error("Expected an error for float value, got nil")
	}

	strProto := prototipe.NewPrototype("hello")
	_, err = strProto.IsEven()
	if err == nil {
		t.Error("Expected an error for non-number value, got nil")
	}
}

func TestIsOdd(t *testing.T) {
	intProto := prototipe.NewPrototype(7)
	resultProto, err := intProto.IsOdd()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if resultProto.Value() != true {
		t.Errorf("Expected true, got %v", resultProto.Value())
	}

	evenProto := prototipe.NewPrototype(10)
	resultProto, err = evenProto.IsOdd()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if resultProto.Value() != false {
		t.Errorf("Expected false, got %v", resultProto.Value())
	}

	floatProto := prototipe.NewPrototype(10.5)
	_, err = floatProto.IsOdd()
	if err == nil {
		t.Error("Expected an error for float value, got nil")
	}

	strProto := prototipe.NewPrototype("hello")
	_, err = strProto.IsOdd()
	if err == nil {
		t.Error("Expected an error for non-number value, got nil")
	}
}

func TestModulo(t *testing.T) {
	intProto := prototipe.NewPrototype(10)
	resultProto, err := intProto.Modulo(3)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if resultProto.Value() != 1 {
		t.Errorf("Expected 1, got %v", resultProto.Value())
	}

	_, err = intProto.Modulo(0)
	if err == nil {
		t.Error("Expected division by zero error, got nil")
	}

	floatProto := prototipe.NewPrototype(10.5)
	_, err = floatProto.Modulo(3)
	if err == nil {
		t.Error("Expected error for float number, got nil")
	}

	strProto := prototipe.NewPrototype("hello")
	_, err = strProto.Modulo(3)
	if err == nil {
		t.Error("Expected error for non-number value, got nil")
	}
}

func TestPower(t *testing.T) {
	intProto := prototipe.NewPrototype(2)
	resultProto, err := intProto.Power(3)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if resultProto.Value() != 8.0 {
		t.Errorf("Expected 8, got %v", resultProto.Value())
	}

	floatProto := prototipe.NewPrototype(2.5)
	resultProto, err = floatProto.Power(2)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if resultProto.Value() != 6.25 {
		t.Errorf("Expected 6.25, got %v", resultProto.Value())
	}

	strProto := prototipe.NewPrototype("hello")
	_, err = strProto.Power(2)
	if err == nil {
		t.Error("Expected an error for non-number value, got nil")
	}
}

func TestSubtract(t *testing.T) {
	intProto := prototipe.NewPrototype(10)
	resultProto, err := intProto.Subtract(3)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if resultProto.Value() != 7 {
		t.Errorf("Expected 7, got %v", resultProto.Value())
	}

	floatProto := prototipe.NewPrototype(15.5)
	resultProto, err = floatProto.Subtract(2.5)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if resultProto.Value() != 13.0 {
		t.Errorf("Expected 13.0, got %v", resultProto.Value())
	}

	strProto := prototipe.NewPrototype("hello")
	_, err = strProto.Subtract(3)
	if err == nil {
		t.Error("Expected an error for non-number value, got nil")
	}
}

func TestAbs(t *testing.T) {
	intPrototype := prototipe.NewPrototype(-10)
	intResult, err := intPrototype.Abs()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if intResult.Value() != 10 {
		t.Errorf("Expected 10, got %v", intResult.Value())
	}

	floatPrototype := prototipe.NewPrototype(-10.5)
	floatResult, err := floatPrototype.Abs()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if floatResult.Value() != 10.5 {
		t.Errorf("Expected 10.5, got %v", floatResult.Value())
	}

	posPrototype := prototipe.NewPrototype(10)
	posResult, err := posPrototype.Abs()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if posResult.Value() != 10 {
		t.Errorf("Expected 10, got %v", posResult.Value())
	}

	stringPrototype := prototipe.NewPrototype("string")
	_, err = stringPrototype.Abs()
	if err == nil {
		t.Errorf("Expected an error for non-numeric value, got none")
	}
}

func TestRound(t *testing.T) {
	numPrototype := prototipe.NewPrototype(10.4)
	result, err := numPrototype.Round()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result.Value() != 10.0 {
		t.Errorf("Expected 10, got %v", result.Value())
	}

	numPrototype = prototipe.NewPrototype(10.6)
	result, err = numPrototype.Round()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result.Value() != 11.0 {
		t.Errorf("Expected 11, got %v", result.Value())
	}

	numPrototype = prototipe.NewPrototype(10)
	result, err = numPrototype.Round()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result.Value() != 10 {
		t.Errorf("Expected 10, got %v", result.Value())
	}

	strPrototype := prototipe.NewPrototype("string")
	_, err = strPrototype.Round()
	if err == nil {
		t.Errorf("Expected an error for non-numeric value, got none")
	}
}

func TestSqrt(t *testing.T) {
	intPrototype := prototipe.NewPrototype(16)
	result, err := intPrototype.Sqrt()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result.Value() != 4.0 {
		t.Errorf("Expected 4, got %v", result.Value())
	}

	floatPrototype := prototipe.NewPrototype(20.25)
	result, err = floatPrototype.Sqrt()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result.Value() != 4.5 {
		t.Errorf("Expected 4.5, got %v", result.Value())
	}

	negativePrototype := prototipe.NewPrototype(-16)
	_, err = negativePrototype.Sqrt()
	if err == nil {
		t.Errorf("Expected an error for negative number, got none")
	}

	strPrototype := prototipe.NewPrototype("string")
	_, err = strPrototype.Sqrt()
	if err == nil {
		t.Errorf("Expected an error for non-numeric value, got none")
	}
}

func TestFactorial(t *testing.T) {
	num := prototipe.NewPrototype(5)
	result, err := num.Factorial()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	expected := 120
	if result.Value() != expected {
		t.Errorf("Expected %v, got %v", expected, result.Value())
	}

	num = prototipe.NewPrototype(0)
	result, err = num.Factorial()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	expected = 1
	if result.Value() != expected {
		t.Errorf("Expected %v, got %v", expected, result.Value())
	}

	num = prototipe.NewPrototype(-3)
	_, err = num.Factorial()
	if err == nil {
		t.Errorf("Expected an error for negative numbers, got none")
	}
}
