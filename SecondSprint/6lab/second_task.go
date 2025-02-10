package main

import (
	"errors"
	"time"
)

func TimeoutFibonacci(n int, timeout time.Duration) (int, error) {
	c := make(chan int, 1)
	go func() {
		F := FibonacciGen(n)
		c <- F
	}()

	select {
	case res := <-c:
		return res, nil
	case <-time.After(timeout):
		return 0, errors.New("время вышло")
	}
}

func FibonacciGen(n int) int {
	arr := []int{1, 1}
	if n == 1 || n == 2 {
		return 1
	} else {
		for i := 2; i < n+1; i++ {
			if i+1 == n {
				return arr[i-1] + arr[i-2]
			}
			arr = append(arr, arr[i-1]+arr[i-2])
		}
	}
	return 0
}
