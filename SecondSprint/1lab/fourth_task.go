package main

func Process(nums []int) chan int {
	ch := make(chan int, 10)
	for _, x := range nums {
		ch <- x
	}
	return ch
}
