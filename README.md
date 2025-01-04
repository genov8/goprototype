# goprototipe

A lightweight Go library for working with strings, slices, and numbers. This library provides chainable methods to enhance productivity and code readability.

---

## Installation

To install the package, use:

```bash
go get -u github.com/genov8/goprototipe
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

### Slices
Handle slices with ease:
- `Append` – Appends an element to a slice.
- `Remove` – Removes an element by index.
- `PrintSlice` – Converts a slice to a string.
- `Unique` – Removes duplicate elements.
- `Contains` – Checks if a slice contains an element.
- `IndexOf` – Finds the index of an element.

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

### Common
Utility methods applicable to multiple types:
- `Length` – Returns the length of a string or slice.
- `Reverse` – Reverses a string or slice.
- `Concat` – Concatenates strings or slices.

---

## Usage

### Basic Example
```go
package main

import (
	"fmt"
	"log"

	"github.com/genov8/goprototipe/prototipe"
)

func main() {
	// Number operations
	num := prototipe.NewPrototype(10)
	result, err := num.Add(5).Multiply(2).Divide(3)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	fmt.Println("Result:", result.Value()) // Result: 10

	// String operations
	str := prototipe.NewPrototype(" hello world ")
	trimmed, _ := str.Trim(" ").Capitalize()
	fmt.Println("Trimmed and Capitalized:", trimmed.Value()) // "Hello world"

	// Slice operations
	slice := prototipe.NewPrototype([]interface{}{1, 2, 2, 3})
	unique, _ := slice.Unique()
	fmt.Println("Unique Slice:", unique.Value()) // [1 2 3]
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
