# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/), and this project adheres to [Semantic Versioning](https://semver.org/).

## [2.0.0] - 2025-01-24
### Repository Renamed
This repository has been renamed from **`goprototipe`** to **`goprototype`**.

## [1.3.0] - 2025-01-22
### Added
- New methods implemented for numbers:
  - `IsPrime` – Checks if a number is a prime number.
  - `Clamp` – Restricts a number to a specified range.
  - `Fibonacci` – Generates a Fibonacci sequence of a given length.
  - `IsPowerOfTwo` – Checks if a number is a power of two.

- New methods implemented for slices:
  - `Intersect` – Finds the common elements between two slices.
  - `Zip` – Combines two slices into pairs (tuples).
  - `Sort` – Sorts a slice in ascending order (supports custom comparators).
  - `Rotate` – Rotates elements in a slice left or right by a given number of steps.

- New methods implemented for strings:
  - `Slugify` – Converts a string into a URL-friendly format by replacing spaces and special characters.
  - `Pad` – Pads a string to a specified length with a given character.
  - `WordCount` – Counts the number of words in a string.
  - `ToCamelCase` – Converts a string to CamelCase format.

- New methods implemented for common:
  - `Type` – Returns the type of the stored value (e.g., string, slice, int).
  - `IsEmpty` – Checks if a value is empty (e.g., string is "", slice is empty, or number is zero).

## [1.2.0] - 2025-01-15
### Added
- New method implemented for numbers:
  - `Factorial` – Calculates factorial numbers.

- New method implemented for slices:
  - `Chunk` – Splits a slice into pieces of a given size.

- New method implemented for strings:
  - `ReverseWords` – Reverses the order of words in a string.

- New method implemented for common:
  - `Compare` – Compares two values (numbers, strings, or slices) for equality.

- Introduced the Chain structure to enable chainable method calls with automatic error handling:
  - Added method Invoke: allows executing methods on Prototype in a chain while automatically stopping further execution if an error occurs.
  - Added method Must: provides strict error handling by invoking a panic if any error is encountered during the chain.

## [1.1.0] - 2025-01-12
### Added
- New methods implemented for numbers:
    - `Abs`: Returns the absolute value of a number.
    - `Round`: Rounds a number to the nearest integer.
    - `Sqrt`: Computes the square root of a number.

- New methods implemented for slices:
    - `Flatten`: Flattens nested slices into a single slice.
    - `Filter`: Filters elements based on a condition.
    - `Map`: Applies a function to each element in a slice and returns the transformed slice.

- New methods implemented for strings:
    - `Replace`: Replaces all occurrences of a substring with another substring.
    - `Repeat`: Repeats a string a specified number of times.

### Changed
- `Contains` method is now universal and works for both strings and slices.

## [1.0.0] - 2025-01-09
### Added
- Initial release with foundational methods for numbers, strings, and slices.

[2.0.0]: https://github.com/genov8/goprototype/releases/tag/v2.0.0
[1.3.0]: https://github.com/genov8/goprototipe/releases/tag/v1.3.0
[1.2.0]: https://github.com/genov8/goprototipe/releases/tag/v1.2.0
[1.1.0]: https://github.com/genov8/goprototipe/releases/tag/v1.1.0
[1.0.0]: https://github.com/genov8/goprototipe/releases/tag/v1.0.0
