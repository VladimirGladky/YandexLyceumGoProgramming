package main

import (
	"errors"
	"fmt"
)

func UnderLimit(nums []int, limit int, n int) ([]int, error) {
	var res []int
	var c int
	if nums == nil {
		return nil, errors.New("input slice is nil")
	}
	if n < 0 {
		return nil, errors.New("n is less than 0")
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] < limit && c < n {
			res = append(res, nums[i])
		}
	}
	return res, nil
}

func Clean(nums []int, x int) []int {
	var j int
	for i := 0; i < len(nums); i++ {
		if nums[i] != x {
			nums[j] = nums[i]
			j++
		}
	}
	return nums[:j]
}

func SliceCopy(nums []int) []int {
	res := make([]int, len(nums))
	for i := range nums {
		res[i] = nums[i]
	}
	return res
}

func Join(nums1, nums2 []int) []int {
	res := make([]int, len(nums1)+len(nums2))
	for i := 0; i < len(nums1); i++ {
		res[i] = nums1[i]
	}
	for i := 0; i < len(nums2); i++ {
		res[len(nums1)+i] = nums2[i]
	}
	return res
}

func Mix(nums []int) []int {
	res := make([]int, 0, len(nums))
	for i := 0; i < len(nums)/2; i++ {
		res = append(res, nums[i])
		res = append(res, nums[len(nums)/2+i])
	}
	return res
}

func main() {
	var a []int
	a, _ = UnderLimit([]int{}, 5, -1)
	fmt.Println(a)

	fmt.Println(Clean([]int{3, 3, 3, 1, 2, 5}, 3))

	fmt.Println(SliceCopy([]int{1, 2, 3}))

	fmt.Println(Join([]int{1, 2, 3}, []int{4, 5}))

	fmt.Println(Mix([]int{3, 3, 3, 2, 2, 2}))
}
