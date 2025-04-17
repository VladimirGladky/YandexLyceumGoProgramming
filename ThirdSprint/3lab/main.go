package main

import "fmt"

func main() {
	wc := WordCounter{wordsCount: make(map[string]int)}
	err := wc.ProcessFiles("name.txt")
	if err != nil {
		return
	}
	fmt.Println(wc.wordsCount)

	nums := []int{3, 5, 3, 6, 6}
	sum, err := ProcessSum(SumChunk, nums, 2)
	if err != nil {
		return
	}
	fmt.Println(sum)
}
