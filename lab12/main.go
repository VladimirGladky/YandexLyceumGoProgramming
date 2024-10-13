package main

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"unicode/utf8"
)

func ConcatStringsAndInt(str1, str2 string, num int) string {
	str3 := strconv.Itoa(num)
	return str1 + str2 + str3
}

func Factorial(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("factorial is not defined for negative numbers")
	}
	if n == 0 {
		return 0, nil
	}
	res := 1
	for i := n; i > 0; i-- {
		res *= i
	}
	return res, nil
}

func DivideIntegers(a, b int) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero is not allowed")
	}
	return float64(a / b), nil
}

func GetCharacterAtPosition(str string, position int) (rune, error) {
	if position < 0 || position >= utf8.RuneCountInString(str) {
		return rune(0), errors.New("position out of range")
	}
	var ans rune
	for i, letter := range str {
		if i == position {
			ans = letter
		}
	}
	return ans, nil
}

func IntToBinary(num int) (string, error) {
	var res string
	if num < 0 {
		return "", errors.New("negative numbers are not allowed")
	}
	for num > 0 {
		res = strconv.Itoa(num%2) + res
		num = num / 2
	}
	return res, nil
}

func SumTwoIntegers(a, b string) (int, error) {
	int1, ok1 := strconv.Atoi(a)
	int2, ok2 := strconv.Atoi(b)
	if ok1 != nil || ok2 != nil {
		return 0, errors.New("invalid input, please provide two integers")
	}
	return int1 + int2, nil
}

func AreAnagrams(str1, str2 string) bool {
	str1 = strings.ToLower(strings.ReplaceAll(str1, " ", ""))
	str2 = strings.ToLower(strings.ReplaceAll(str2, " ", ""))

	runes1 := []rune(str1)
	runes2 := []rune(str2)

	sort.Slice(runes1, func(i, j int) bool {
		return runes1[i] < runes1[j]
	})
	sort.Slice(runes2, func(i, j int) bool {
		return runes2[i] < runes2[j]
	})

	return string(runes1) == string(runes2)
}

func main() {
	fmt.Println(Factorial(6))

	fmt.Println(IntToBinary(11))

	fmt.Println(SumTwoIntegers("1	.0", "2"))
}
