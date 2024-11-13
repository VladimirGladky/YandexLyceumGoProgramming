package main

import (
	"fmt"
	"math"
)

func IsPowerOfTwoRecursive(N int) {
	i := 1
	for i < N {
		i *= 2
	}
	if i == N {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func CalculateSeriesSum(n int) float64 {
	var sum float64
	for i := 1; i < n+1; i++ {
		sum += 1.0 / float64(i)
	}
	return sum
}

func CalculateDigitalRoot(n int) int {
	if n < 10 {
		return n
	}
	sum := SumDigitsRecursive(n)
	for sum >= 10 {
		sum = SumDigitsRecursive(sum)
	}

	return sum
}

func FindIntersection(k1, b1, k2, b2 float64) (float64, float64) {
	if k1 == k2 {
		return math.NaN(), math.NaN()
	} else {
		x := (b2 - b1) / (k1 - k2)
		y := k1*x + b1
		return x, y
	}
}

func SumDigitsRecursive(n int) int {
	if n < 10 {
		return n
	}
	return SumDigitsRecursive(n/10) + n%10
}

func sqRoots() {
	var a, b, c float64
	_, err := fmt.Scanln(&a, &b, &c)
	if err != nil && a != 0 {
		fmt.Println("Некорректный ввод")
		return
	}

	discr := math.Pow(b, 2) - 4*a*c
	if discr < 0 {
		fmt.Println("0 0")
	} else if discr == 0 {
		fmt.Printf("%v", -b/(2*a))
	} else {
		first, second := (-b+math.Sqrt(discr))/(2*a), (-b-math.Sqrt(discr))/(2*a)
		if first > second {
			fmt.Println(second, first)
		} else {
			fmt.Println(first, second)
		}

	}

}

func main() {
	//sqRoots()
	//fmt.Println(SumDigitsRecursive(13333))
	fmt.Println(CalculateSeriesSum(3))
}
