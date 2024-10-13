package main

import "fmt"

func throwUpTo10() {
	for i := 1; i < 10; i++ {
		fmt.Println(i)
	}
}

func multiplicationTable() {
	var number int
	fmt.Scanln(&number)
	for i := 1; i < 10; i++ {
		fmt.Println(number, "*", i, "=", number*i)
	}
}

func waterInACieve() {
	var sum, number int
	fmt.Scanln(&number)
	for i := 1; i < number+1; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func permutationsLetter() {
	var number int
	res := 1
	fmt.Scanln(&number)
	for i := 1; i < number+1; i++ {
		res *= i
	}

	fmt.Println(res)
}

func oddSide() {
	var sum, number int
	_, err := fmt.Scanln(&number)
	if err != nil || number < 0 {
		fmt.Println("Некорректный ввод")
		return
	}
	for i := 1; i < number+1; i++ {
		if i%2 == 0 {
			continue
		}
		sum += i
	}
	fmt.Println(sum)
}

func forGrandma() {
	var number int
	fmt.Scanln(&number)
	for i := 1; i < number+1; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}
}

func donutTreat() {
	var number, sum int
	fmt.Scanln(&number)
	for i := 1; i < number+1; i++ {
		if i%3 == 0 || i%5 == 0 {
			continue
		}
		sum += i
	}
	fmt.Println(sum)
}

func fibonacciRabits() {
	var number, c int
	fmt.Scanln(&number)
	b, a := 0, 1
	for a <= number {
		c = a + b
		b = a
		a = c
	}

	if b == number {
		for i := 0; i < 10; i++ {
			fmt.Println(b)
			c = a + b
			b = a
			a = c
		}
	} else {
		for i := 0; i < 10; i++ {
			fmt.Println(a)
			c = a + b
			b = a
			a = c
		}
	}
}

func main() {
	//throwUpTo10()
	//multiplicationTable()
	//waterInACieve()
	//permutationsLetter()
	//oddSide()
	//forGrandma()
	//donutTreat()
	fibonacciRabits()
}
