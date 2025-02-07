package main

func MaxExpressionValue(nums []int) int {
	first := make([]int, len(nums)+1)
	for i := len(nums) - 1; i >= 0; i-- {
		first[i] = max(first[i+1], nums[i]) // Функция max — возвращает максимальное
	}
	second := make([]int, len(nums))
	for i := len(nums) - 2; i >= 0; i-- {
		second[i] = max(second[i+1], first[i+1]-nums[i])
	}

	third := make([]int, len(nums)-1)
	for i := len(nums) - 3; i >= 0; i-- {
		third[i] = max(third[i+1], second[i+1]+nums[i])
	}

	fourth := make([]int, len(nums)-2)
	for i := len(nums) - 4; i >= 0; i-- {
		fourth[i] = max(fourth[i+1], third[i+1]-nums[i])
	}

	return fourth[0]
}
