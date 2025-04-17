package main

import (
	"context"
	"sync"
)

func FanIn[T any](ctx context.Context, channels ...<-chan T) <-chan T {
	out := make(chan T)
	wg := sync.WaitGroup{}
	for _, channel := range channels {
		wg.Add(1)
		go func(ch <-chan T) {
			defer wg.Done()
			for {
				select {
				case data, ok := <-ch:
					if !ok {
						return
					}
					out <- data
				case <-ctx.Done():
					return
				}
			}
		}(channel)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
