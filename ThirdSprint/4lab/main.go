package main

import "fmt"

func main() {
	inputNums := [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(MultiplyPipeline(inputNums...))
}
