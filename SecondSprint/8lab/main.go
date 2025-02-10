package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var sharedData int
	iterations := 100
	var mu sync.Mutex

	for i := 0; i < iterations; i++ {
		wg.Add(1)
		go func() {
			mu.Lock()
			sharedData++ // несинхронизированный доступ к общим данным
			mu.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Final value of sharedData:", sharedData)
}
