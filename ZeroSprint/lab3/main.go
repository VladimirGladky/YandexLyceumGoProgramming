package main

import "fmt"

func separatingZero() {
	var number int
	fmt.Scanln(&number)
	if number == 0 {
		fmt.Println("Введен ноль")
	} else if number >= 0 {
		fmt.Println("Число положительное")
	} else {
		fmt.Println("Число отрицательное")
	}
}

func subjectivelyMore() {
	var a1, a2 int
	fmt.Scanf("%d %d", &a1, &a2)
	if a1 == a2 {
		fmt.Println("Числа равны")
	} else if a1 >= a2 {
		fmt.Println("Первое число больше второго")
	} else {
		fmt.Println("Второе число больше первого")
	}
}

func darkForest() {
	var number int
	fmt.Scanln(&number)
	if number == 0 {
		fmt.Println("Число равно нулю")
	} else if number > 0 {
		if number%2 == 0 {
			fmt.Println("Число положительное и четное")
		} else {
			fmt.Println("Число положительное и нечетное")
		}
	} else {
		if number%2 == 0 {
			fmt.Println("Число отрицательное и четное")
		} else {
			fmt.Println("Число отрицательное и нечетное")
		}
	}
}

func aboutZero() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)
	if a == 0 || b == 0 {
		fmt.Println("Одно из чисел равно нулю")
	} else if a > 0 && b > 0 {
		fmt.Println("Оба числа положительные")
	} else if a < 0 && b < 0 {
		fmt.Println("Оба числа отрицательные")
	} else {
		fmt.Println("Одно число положительное, а другое отрицательное")
	}
}

func equality() {
	var a, b, c int
	_, err := fmt.Scanf("%d %d %d", &a, &b, &c)
	if err != nil {
		fmt.Println("Некорректный ввод")
		return
	}
	if a == c && b == c {
		fmt.Println("Все числа равны")
	} else if a == c || a == b || b == c {
		fmt.Println("Два числа равны")
	} else {
		fmt.Println("Все числа разные")
	}
}

func manyStars() {
	var number int
	_, err := fmt.Scanln(&number)
	if err != nil || number < 0 {
		fmt.Println("Некорректный ввод")
		return
	}

	if number <= 10 {
		fmt.Println("Число меньше 10")
	} else if number <= 100 {
		fmt.Println("Число меньше 100")
	} else if number < 1000 {
		fmt.Println("Число меньше 1000")
	} else {
		fmt.Println("Число больше или равно 1000")
	}

}

func main() {
	//separatingZero()
	//subjectivelyMore()
	//darkForest()
	//aboutZero()
	//equality()
	manyStars()
}
