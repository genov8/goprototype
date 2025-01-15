package tests

import (
	"github.com/genov8/goprototipe/prototipe"
	"testing"
)

func TestChainSuccess(t *testing.T) {
	str := prototipe.NewPrototype("Hello World")
	chain := prototipe.NewChain(str).
		Invoke(func(p *prototipe.Prototype) (*prototipe.Prototype, error) {
			return p.ReverseWords()
		}).
		Invoke(func(p *prototipe.Prototype) (*prototipe.Prototype, error) {
			return p.ToUpperCase()
		})

	if chain.Err != nil {
		t.Errorf("Expected no error, got %v", chain.Err)
	}

	if chain.Prototype.Value() != "WORLD HELLO" {
		t.Errorf("Expected 'WORLD HELLO', got %v", chain.Prototype.Value())
	}
}

func TestChainError(t *testing.T) {
	str := prototipe.NewPrototype(42)
	chain := prototipe.NewChain(str).
		Invoke(func(p *prototipe.Prototype) (*prototipe.Prototype, error) {
			return p.ToUpperCase()
		}).
		Invoke(func(p *prototipe.Prototype) (*prototipe.Prototype, error) {
			return p.ReverseWords()
		})

	if chain.Err == nil {
		t.Errorf("Expected an error, got nil")
	}
}

func TestChainMustSuccess(t *testing.T) {
	num := prototipe.NewPrototype(42)
	chain := prototipe.NewChain(num).
		Invoke(func(p *prototipe.Prototype) (*prototipe.Prototype, error) {
			return p.Add(42)
		}).
		Must()

	if chain.Value() != 84 {
		t.Errorf("Expected '84' got %v", chain.Value())
	}
}

func TestChainMustError(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic, but no panic occurred")
		}
	}()

	str := prototipe.NewPrototype(42)
	prototipe.NewChain(str).
		Invoke(func(p *prototipe.Prototype) (*prototipe.Prototype, error) {
			return p.ToUpperCase()
		}).
		Must()
}
