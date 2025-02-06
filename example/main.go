package main

import (
	"fmt"
	"github.com/genov8/goprototype/v2/prototype"
	"log"
)

func main() {
	num := prototype.NewNumberPrototype(10)
	numAdd, err := num.Add(2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Result add:", numAdd.Value())

	sliceProto := prototype.NewPrototype([]interface{}{1, "hello", 3.14})

	newSliceProto, err := sliceProto.Append("new_element")
	if err != nil {
		log.Fatalf("Error adding element: %v", err)
	}
	fmt.Println("Slice after adding:", newSliceProto.Value())

	// usage with chain

	str := prototype.NewPrototype("Hello world")
	result := prototype.NewChain(str).
		Invoke(func(p *prototype.Prototype) (*prototype.Prototype, error) {
			return p.ReverseWords()
		}).
		Invoke(func(p *prototype.Prototype) (*prototype.Prototype, error) {
			return p.ToUpperCase()
		}).
		Must()

	fmt.Println("Result:", result.Value()) // "WORLD HELLO"
}
