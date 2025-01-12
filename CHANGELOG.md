# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/), and this project adheres to [Semantic Versioning](https://semver.org/).

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

[1.1.0]: https://github.com/genov8/goprototipe/releases/tag/v1.1.0
[1.0.0]: https://github.com/genov8/goprototipe/releases/tag/v1.0.0