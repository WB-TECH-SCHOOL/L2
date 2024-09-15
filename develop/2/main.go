package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func unpackString(input string) (string, error) {
	if input == "" {
		return "", nil
	}

	var result strings.Builder
	runes := []rune(input)
	length := len(runes)

	for i := 0; i < length; i++ {
		char := runes[i]

		if unicode.IsDigit(char) {
			return "", fmt.Errorf("invalid string: starts with a digit or contains consecutive digits")
		}

		if i+1 < length && unicode.IsDigit(runes[i+1]) {
			repeatCount, err := strconv.Atoi(string(runes[i+1]))
			if err != nil {
				return "", fmt.Errorf("invalid repeat count")
			}
			result.WriteString(strings.Repeat(string(char), repeatCount))
			i++
		} else {
			result.WriteRune(char)
		}
	}

	return result.String(), nil
}

func main() {
	tests := []string{
		"a4bc2d5e",
		"abcd",
		"45",
		"",
	}

	for _, test := range tests {
		unpacked, err := unpackString(test)
		if err != nil {
			fmt.Printf("Error unpacking string %q: %v\n", test, err)
		} else {
			fmt.Printf("Unpacked string %q: %q\n", test, unpacked)
		}
	}
}
