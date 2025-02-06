package prototype

import (
	"errors"
	"regexp"
	"strings"
	"unicode"
)

func (p *StringPrototype) ToUpperCase() (*StringPrototype, error) {
	return p.processString(strings.ToUpper)
}

func (p *StringPrototype) ToLowerCase() (*StringPrototype, error) {
	return p.processString(strings.ToLower)
}

func (p *StringPrototype) Capitalize() (*StringPrototype, error) {
	return p.processString(func(str string) string {
		if len(str) == 0 {
			return str
		}
		firstSymbol := str[:1]
		lastStr := str[1:]
		return strings.ToUpper(firstSymbol) + strings.ToLower(lastStr)
	})
}

func (p *StringPrototype) Trim(cutset string) (*StringPrototype, error) {
	return p.processString(func(str string) string {
		return strings.Trim(str, cutset)
	})
}

func (p *StringPrototype) Split(separator string) (*StringPrototype, error) {
	if str, ok := p.value.(string); ok {
		return &StringPrototype{&Prototype{value: strings.Split(str, separator)}}, nil
	}
	return nil, errors.New("value is not a string")
}

func (p *StringPrototype) StartsWith(prefix string) (*StringPrototype, error) {
	if str, ok := p.value.(string); ok {
		result := strings.HasPrefix(str, prefix)
		return &StringPrototype{&Prototype{value: result}}, nil
	}
	return nil, errors.New("value is not a string")
}

func (p *StringPrototype) EndsWith(prefix string) (*StringPrototype, error) {
	if str, ok := p.value.(string); ok {
		result := strings.HasSuffix(str, prefix)
		return &StringPrototype{&Prototype{value: result}}, nil
	}
	return nil, errors.New("value is not a string")
}

func (p *StringPrototype) RemoveWhitespace() (*StringPrototype, error) {
	return p.processString(func(str string) string {
		return strings.ReplaceAll(str, " ", "")
	})
}

func (p *StringPrototype) Replace(old, new string) (*StringPrototype, error) {
	return p.processString(func(str string) string {
		return strings.ReplaceAll(str, old, new)
	})
}

func (p *StringPrototype) Repeat(count int) (*StringPrototype, error) {
	if str, ok := p.value.(string); ok {
		if count < 0 {
			return nil, errors.New("repeat count must be non-negative")
		}
		return &StringPrototype{&Prototype{value: strings.Repeat(str, count)}}, nil
	}
	return nil, errors.New("value is not a string")
}

func (p *StringPrototype) ReverseWords() (*StringPrototype, error) {
	if str, ok := p.value.(string); ok {
		words := strings.Fields(str)
		for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
			words[i], words[j] = words[j], words[i]
		}
		return &StringPrototype{&Prototype{value: strings.Join(words, " ")}}, nil
	}
	return nil, errors.New("value is not a string")
}

func (p *StringPrototype) Slugify() (*StringPrototype, error) {
	if str, ok := p.value.(string); ok {
		re := regexp.MustCompile(`[^a-zA-Z0-9\s-]+`)
		cleaned := re.ReplaceAllString(str, "")
		slug := strings.ToLower(strings.ReplaceAll(cleaned, " ", "-"))
		slug = regexp.MustCompile(`-{2,}`).ReplaceAllString(slug, "-")
		return &StringPrototype{&Prototype{value: slug}}, nil
	}
	return nil, errors.New("value is not a string")
}

func (p *StringPrototype) Pad(length int, char string, padLeft bool) (*StringPrototype, error) {
	if str, ok := p.value.(string); ok {
		if len(char) != 1 {
			return nil, errors.New("padding character must be a single character")
		}

		strLen := len(str)
		if strLen >= length {
			return &StringPrototype{&Prototype{value: str}}, nil
		}

		padding := strings.Repeat(char, length-strLen)
		if padLeft {
			return &StringPrototype{&Prototype{value: padding + str}}, nil
		}
		return &StringPrototype{&Prototype{value: str + padding}}, nil
	}
	return nil, errors.New("value is not a string")
}

func (p *StringPrototype) WordCount() (*StringPrototype, error) {
	if str, ok := p.value.(string); ok {
		words := strings.Fields(str)
		return &StringPrototype{&Prototype{value: len(words)}}, nil
	}
	return nil, errors.New("value is not a string")
}

func (p *StringPrototype) ToCamelCase() (*StringPrototype, error) {
	if str, ok := p.value.(string); ok {
		words := strings.FieldsFunc(str, func(r rune) bool {
			return !unicode.IsLetter(r) && !unicode.IsNumber(r)
		})

		for i := range words {
			words[i] = strings.Title(strings.ToLower(words[i]))
		}

		camelCase := strings.Join(words, "")
		return &StringPrototype{&Prototype{value: camelCase}}, nil
	}
	return nil, errors.New("value is not a string")
}

func (p *StringPrototype) processString(operation func(string) string) (*StringPrototype, error) {
	if str, ok := p.value.(string); ok {
		return &StringPrototype{&Prototype{value: operation(str)}}, nil
	}
	return nil, errors.New("value is not a string")
}
