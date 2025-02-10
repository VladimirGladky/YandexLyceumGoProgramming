package main

import "sync"

var (
	Buf []int
	mu  sync.Mutex
)

func Write(num int) {
	mu.Lock()
	Buf = append(Buf, num)
	mu.Unlock()
}

func Consume() int {
	mu.Lock()
	defer mu.Unlock()
	if len(Buf) == 0 {
		return 0
	}
	elem := Buf[0]
	Buf = Buf[1:]
	return elem
}
