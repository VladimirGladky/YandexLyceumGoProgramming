package rpn

import (
	"errors"
	"strconv"
	"unicode"
)

func Calc(expression string) (float64, error) {
	var numbersStack []float64
	var operationsStack []rune
	var i int
	var leftBrackets int
	var rightBrackets int

	for i < len(expression) {
		if unicode.IsDigit(rune(expression[i])) {
			num, _ := strconv.ParseFloat(string(expression[i]), 64)
			numbersStack = append(numbersStack, num)
			i++
			continue
		}
		switch expression[i] {
		case '+', '-', '*', '/':
			for len(operationsStack) > 0 && priorityOperations(operationsStack[len(operationsStack)-1]) >= priorityOperations(rune(expression[i])) {
				numbersStack, operationsStack = makeOperation(numbersStack, operationsStack)
			}
			operationsStack = append(operationsStack, rune(expression[i]))
			i++
			continue
		}

		if expression[i] == '(' {
			operationsStack = append(operationsStack, rune(expression[i]))
			leftBrackets++
			i++
			continue
		}

		if expression[i] == ')' {
			rightBrackets++
			if leftBrackets >= rightBrackets {
				for operationsStack[len(operationsStack)-1] != '(' {
					numbersStack, operationsStack = makeOperation(numbersStack, operationsStack)
				}
				operationsStack = operationsStack[:len(numbersStack)-1]
				leftBrackets--
				rightBrackets--
			} else {
				return 0, errors.New("неправильный ввод")
			}
			i++
			continue
		}
	}

	if leftBrackets != 0 || rightBrackets != 0 {
		return 0, errors.New("неправильный ввод")
	}

	if len(numbersStack)-1 != len(operationsStack) {
		return 0, errors.New("неправильный ввод")
	}

	for len(operationsStack) > 0 {
		numbersStack, operationsStack = makeOperation(numbersStack, operationsStack)
	}

	return numbersStack[0], nil
}

func makeOperation(numbersStack []float64, operationsStack []rune) ([]float64, []rune) {
	if len(numbersStack) < 2 || len(operationsStack) == 0 {
		return numbersStack, operationsStack
	}
	a := numbersStack[len(numbersStack)-2]
	b := numbersStack[len(numbersStack)-1]
	operation := operationsStack[len(operationsStack)-1]

	var result float64
	switch operation {
	case '+':
		result = a + b
	case '-':
		result = a - b
	case '*':
		result = a * b
	case '/':
		if b == 0 {
			panic("деление на ноль")
		}
		result = a / b
	default:
		panic("неизвестная операция")
	}

	numbersStack = numbersStack[:len(numbersStack)-2]
	operationsStack = operationsStack[:len(operationsStack)-1]
	return append(numbersStack, result), operationsStack
}

func priorityOperations(operation rune) int {
	switch operation {
	case '*', '/':
		return 2
	case '+', '-':
		return 1
	}
	return 0
}
