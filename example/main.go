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

	sliceProto := prototype.NewSlicePrototype([]interface{}{1, "hello", 3.14})

	newSliceProto, err := sliceProto.Append("new_element")
	if err != nil {
		log.Fatalf("Error adding element: %v", err)
	}
	fmt.Println("Slice after adding:", newSliceProto.Value())

	stringProto := prototype.NewStringPrototype("Hello")
	newStringProto, err := stringProto.Reverse()

	if err != nil {
		log.Fatalf("Error string reverse: %v", err)
	}
	fmt.Println("String reverse:", newStringProto.Value())
}
