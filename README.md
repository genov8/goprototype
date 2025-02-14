# goprototype

A lightweight Go library for working with strings, slices, and numbers. This library provides chainable methods to enhance productivity and code readability.

---

## Installation

To install the package, use:

```bash
go get -u github.com/genov8/goprototype/v2
```

---

## Features

### Numbers
Perform mathematical operations on numbers:
- `Add` – Adds a number.
- `Multiply` – Multiplies by a number.
- `Divide` – Divides by a number (handles division by zero).
- `IsEven` – Checks if a number is even.
- `IsOdd` – Checks if a number is odd.
- `Modulo` – Calculates the remainder of division.
- `Power` – Raises a number to the power.
- `Subtract` – Subtracts a number.
- `Abs` – Returns the absolute value of a number.
- `Round` – Rounds a number to the nearest integer.
- `Sqrt` – Returns the square root of a number.
- `Factorial` – Calculates factorial numbers.
- `IsPrime` – Checks if a number is a prime number.
- `Clamp` – Restricts a number to a specified range.
- `Fibonacci` – Generates a Fibonacci sequence of a given length.
- `IsPowerOfTwo` – Checks if a number is a power of two.

### Slices
Handle slices with ease:
- `Append` – Appends an element to a slice.
- `Remove` – Removes an element by index.
- `PrintSlice` – Converts a slice to a string.
- `Unique` – Removes duplicate elements.
- `Contains` – Checks if a slice contains an element.
- `IndexOf` – Finds the index of an element.
- `Flatten` – Flattens nested slices into a single slice.
- `Map` – Applies a function to all elements of a slice.
- `Filter` – Returns a slice with elements that satisfy a condition.
- `Chunk` – Splits a slice into pieces of a given size.
- `Intersect` – Finds the common elements between two slices.
- `Zip` – Combines two slices into pairs (tuples).
- `Sort` – Sorts a slice in ascending order (supports custom comparators).
- `Rotate` – Rotates elements in a slice left or right by a given number of steps.

### Strings
Chainable methods for strings:
- `ToUpperCase` – Converts a string to uppercase.
- `ToLowerCase` – Converts a string to lowercase.
- `Capitalize` – Capitalizes the first letter.
- `Trim` – Trims characters from a string.
- `Split` – Splits a string by a separator.
- `StartsWith` – Checks if a string starts with a prefix.
- `EndsWith` – Checks if a string ends with a suffix.
- `RemoveWhitespace` – Removes all whitespace.
- `Replace` – Replaces all occurrences of a substring with another
- `Repeat` – Repeats a string a specified number of times.
- `ReverseWords` – Reverses the order of words in a string.
- `Slugify` – Converts a string into a URL-friendly format by replacing spaces and special characters.
- `Pad` – Pads a string to a specified length with a given character.
- `WordCount` – Counts the number of words in a string.
- `ToCamelCase` – Converts a string to CamelCase format.


### Common
Utility methods applicable to multiple types:
- `Length` – Returns the length of a string or slice.
- `Reverse` – Reverses a string or slice.
- `Concat` – Concatenates strings or slices.
- `Contains` – Checks if a string contains a substring.
- `Compare` – Compares two values (numbers, strings, or slices) for equality.
- `Type` – Returns the type of the stored value (e.g., string, slice, int).
- `IsEmpty` – Checks if a value is empty (e.g., string is "", slice is empty, or number is zero).

---

## Usage

### Example

```go
package main

import (
	"fmt"
	"github.com/genov8/goprototype/v2/prototype"
	"log"
)

func main() {
	// Number operations
	num := prototype.NewPrototype(10)
	result, err := num.Add(5)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	fmt.Println("Result:", result.Value()) // Result: 15

	// String operations
	str := prototype.NewPrototype(" hello world ")
	trimmed, _ := str.Trim(" ")
	fmt.Println("Trimmed and Capitalized:", trimmed.Value()) // "hello world"

	// Slice operations
	slice := prototype.NewPrototype([]interface{}{1, 2, 2, 3})
	unique, _ := slice.Unique()
	fmt.Println("Unique Slice:", unique.Value()) // [1 2 3]
}
```

## Usage with chain

### Why Chain?

The main idea of this package is to allow chainable method calls to improve code readability and simplify operations on data. However, Go encourages explicit error handling, which can make chaining verbose and less elegant.

To address this, we introduced the `Chain` mechanism:

- Automatically tracks errors during the chain execution.
- Stops the chain when an error occurs.
- Provides strict error handling via the `Must()` method.

With `Chain`, you can enjoy the benefits of chainable calls while adhering to Go's best practices for error handling.

### Example

```go
package main

import (
	"fmt"
	"github.com/genov8/goprototype/v2/prototype"
)

func main() {
	// Chain operations with error handling
	str := prototype.NewPrototype("Hello World")
	chain := prototype.NewChain(str).
		Invoke(func(p *prototype.Prototype) (*prototype.Prototype, error) {
			return p.ReverseWords()
		}).
		Invoke(func(p *prototype.Prototype) (*prototype.Prototype, error) {
			return p.ToUpperCase()
		})

	if chain.Err != nil {
		fmt.Println("Error:", chain.Err)
	} else {
		fmt.Println("Result:", chain.Prototype.Value()) // "WORLD HELLO"
	}

	// Using Must() for strict error handling
	result := prototype.NewChain(prototype.NewPrototype("Go Programming")).
		Invoke(func(p *prototype.Prototype) (*prototype.Prototype, error) {
			return p.Split(" ")
		}).
		Invoke(func(p *prototype.Prototype) (*prototype.Prototype, error) {
			return p.Reverse()
		}).
		Must()

	fmt.Println("Strict Chain Result:", result.Value()) // Output: ["Programming" "Go"]
}
```

---

## Testing

Run the test suite to ensure everything is working:

```bash
go test ./tests -v
```

---

## Contributing

Feel free to open issues or submit pull requests. Contributions are always welcome!

---

## License

This project is licensed under the MIT License.
