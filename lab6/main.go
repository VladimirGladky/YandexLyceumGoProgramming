package main

import "fmt"

func FiveSteps(array [5]int) [5]int {
	var numbers [5]int
	for i := len(array) - 1; i >= 0; i-- {
		numbers[len(array)-i-1] = array[i]
	}

	return numbers
}

func ThirdElementInArray(array [6]int) int {
	return array[2]
}

func FindMinMaxInArray(array [10]int) (int, int) {
	max := -1000000000
	min := 100000000
	for i := 0; i < len(array); i++ {
		if array[i] > max {
			max = array[i]
		}
		if array[i] < min {
			min = array[i]
		}
	}
	return max, min
}

func SumOfArray(array [6]int) int {
	var sum int
	for i := 0; i < len(array); i++ {
		sum += array[i]
	}
	return sum
}

func PrettyArrayOutput(array [9]string) {
	for i := 0; i < len(array); i++ {
		if i < 7 {
			fmt.Printf("%v я уже сделал: %v", i+1, array[i])
			fmt.Println()
		} else {
			fmt.Printf("%v не успел сделать: %v", i+1, array[i])
			fmt.Println()
		}
	}
}

func UnderLimit(nums []int, limit int, n int) ([]int, error) {
	res := make([]int, n)
	var c int
	for i := 0; i < len(nums); i++ {
		if nums[i] < limit && c < n {
			res = append(res, nums[i])
		}
	}
	return nums, nil
}

func main() {
	//var numbers = [...]int{1, 2, 3, 4, 5, 6}
	//fmt.Println(FiveSteps(numbers))
	//fmt.Println(ThirdElementInArray(numbers))
	//fmt.Println(FindMinMaxInArray(numbers))
	//fmt.Println(SumOfArray(numbers))
	//var deals = [...]string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	//PrettyArrayOutput(deals)
	var a []int
	a, _ = UnderLimit([]int{4, 7, 89, 3, 21, 2, 5, 7, 32, 4, 6, 8, 0, 3, 4, 6, 2, 115, 12}, 3, 5)
	fmt.Println(a)
}
