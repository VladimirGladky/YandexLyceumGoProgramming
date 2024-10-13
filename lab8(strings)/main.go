package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

func StringLength(input string) int {
	return len(input)
}

func ConcatenateStrings(str1, str2 string) string {
	return str1 + " " + str2
}

func isLatin(input string) bool {
	for _, letter := range input {
		if !unicode.Is(unicode.Latin, letter) {
			return false
		}
	}
	return true
}

func IsPalindrome(input string) bool {
	input = strings.ReplaceAll(input, " ", "")
	var reverseString string
	for i := len(input) - 1; i >= 0; i-- {
		reverseString = reverseString + string(input[i])
	}
	return input == reverseString
}

func Permutations(input string) []string {
	// Базовый случай: если длина строки 1, возвращаем эту строку
	if len(input) == 1 {
		return []string{input}
	}

	var result []string
	// Проходим по каждому символу строки
	for i, ch := range input {
		// Создаем новую строку без символа на позиции i
		rest := input[:i] + input[i+1:]

		// Рекурсивно вызываем функцию для оставшейся строки
		for _, subPerm := range Permutations(rest) {
			// Добавляем текущий символ в начало каждой перестановки
			result = append(result, string(ch)+subPerm)
		}
	}

	// Сортируем результат в алфавитном порядке
	sort.Strings(result)
	return result
}

func main() {
	fmt.Println("12345")
	fmt.Println(ConcatenateStrings("Ira", "Hom"))
	fmt.Println(IsPalindrome(" a b c       cba  "))
	fmt.Println(Permutations("cab"))
}
