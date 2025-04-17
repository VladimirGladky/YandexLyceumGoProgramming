package main

import "fmt"

func ToString[T any](done <-chan struct{}, valueStream <-chan T) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for v := range valueStream {
			select {
			case <-done:
				return
			case out <- fmt.Sprint(v):
			}
		}
	}()
	return out
}
