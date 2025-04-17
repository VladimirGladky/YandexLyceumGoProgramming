package main

import (
	"errors"
	"sync"
)

type Summer interface {
	// функция для нахождения суммы чисел
	ProcessSum(
		// функция, которая будет вызываться для нахождения суммы части слайса.
		// результат суммы записать в result
		summer func(arr []int, result chan<- int),
		// слайс с числами, сумму которых нужно найти
		nums []int,
		// сколько элементов в одной части (последняя часть может быть меньше)
		сhunkSize int,
	) (int, error) // вернуть сумму чисел
}

func ProcessSum(summer func(arr []int, result chan<- int), nums []int, chunkSize int) (int, error) {
	if chunkSize <= 0 {
		return 0, errors.New("chunkSize must be greater than 0")
	}
	if len(nums) == 0 {
		return 0, nil
	}
	var wg sync.WaitGroup
	out := make(chan int, len(nums)/chunkSize+1)
	for i := 0; i < len(nums); i += chunkSize {
		wg.Add(1)
		go func(start int) {
			defer wg.Done()
			end := start + chunkSize
			if end > len(nums) {
				end = len(nums)
			}
			summer(nums[start:end], out)
		}(i)
	}
	wg.Wait()
	close(out)
	var res int
	for value := range out {
		res += value
	}
	return res, nil
}

func SumChunk(arr []int, result chan<- int) {
	var sum int
	for _, v := range arr {
		sum += v
	}
	result <- sum
}
