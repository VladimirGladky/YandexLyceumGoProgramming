package main

import "context"

type sequenced interface {
	getSequence() int
}

func EvenNumbersGen[T sequenced](ctx context.Context, numbers ...T) <-chan T {
	out := make(chan T)
	go func() {
		defer close(out)
		for _, number := range numbers {
			if number.getSequence()%2 == 0 {
				select {
				case <-ctx.Done():
					return
				case out <- number:
				}
			}
		}
	}()
	return out
}
