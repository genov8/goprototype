package main

import (
	"fmt"
	"goprototipe/prototipe"
	"log"
)

func main() {
	num := prototipe.NewPrototype(10)
	numAdd, err := num.Add(2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Result add:", numAdd.Value())

	strProto := prototipe.NewPrototype("Hello")
	newStrProto, err := strProto.Concat(" World!")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Result concat:", newStrProto.Value())
}
