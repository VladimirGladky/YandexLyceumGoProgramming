package main

import (
	"errors"
	"fmt"
	"unicode"
	"unicode/utf8"
)

func Sum(a, b int) int {
	return a + b
}

func Multiply(a, b int) int {
	return a * b
}

func DeleteVowels(s string) string {
	var result string
	for i := 0; i < len(s); i++ {
		switch unicode.ToLower(rune(s[i])) {
		case 'a':
			continue
		case 'e':
			continue
		case 'i':
			continue
		case 'o':
			continue
		case 'u':
			continue
		}
		result += string(s[i])
	}
	return result
}

var ErrInvalidUTF8 = errors.New("invalid utf8")

func GetUTFLength(input []byte) (int, error) {
	if !utf8.Valid(input) {
		return 0, ErrInvalidUTF8
	}

	return utf8.RuneCount(input), nil
}

func main() {
	fmt.Println("5")
}
