package main

import (
	"fmt"
	"github.com/genov8/goprototipe/prototipe"
	"log"
)

func main() {
	num := prototipe.NewPrototype(10)
	numAdd, err := num.Add(2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Result add:", numAdd.Value())

	sliceProto := prototipe.NewPrototype([]interface{}{1, "hello", 3.14})

	newSliceProto, err := sliceProto.Append("new_element")
	if err != nil {
		log.Fatalf("Error adding element: %v", err)
	}
	fmt.Println("Slice after adding:", newSliceProto.Value())

	// usage with chain

	str := prototipe.NewPrototype("Hello")
	result := prototipe.NewChain(str).
		Invoke(func(p *prototipe.Prototype) (*prototipe.Prototype, error) {
			return p.ReverseWords()
		}).
		Invoke(func(p *prototipe.Prototype) (*prototipe.Prototype, error) {
			return p.ToUpperCase()
		}).
		Must()

	fmt.Println("Result:", result.Value()) // "WORLD HELLO"
}
