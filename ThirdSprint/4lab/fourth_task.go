package main

import (
	"sync"
)

func MultiplyPipeline(inputNums ...[]int) int {
	res := 1
	var wg sync.WaitGroup
	var mu sync.Mutex
	for _, nums := range inputNums {
		wg.Add(1)
		go func(nums []int) {
			defer wg.Done()
			prod := Multiply(Filter(NumbersGen(nums...)))
			mu.Lock()
			res *= prod
			mu.Unlock()
		}(nums)
	}
	wg.Wait()
	return res
}

func NumbersGen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, num := range nums {
			out <- num
		}
	}()
	return out
}

func Filter(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for num := range in {
			if num >= 0 {
				out <- num
			}
		}
	}()
	return out
}

func Multiply(in <-chan int) int {
	res := 1
	for num := range in {
		res *= num
	}
	return res
}
