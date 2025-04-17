package main

func Tee[T any](done <-chan struct{}, in <-chan T) (<-chan T, <-chan T) {
	out1 := make(chan T)
	out2 := make(chan T)
	go func() {
		defer close(out1)
		defer close(out2)
		for v := range in {
			select {
			case <-done:
				return
			case out1 <- v:
				out2 <- v
			}
		}
	}()
	return out1, out2
}
