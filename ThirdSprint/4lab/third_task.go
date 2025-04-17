package main

import (
	"bufio"
	"os"
	"strconv"
)

func SumValuesPipeline(filename string) int {
	return Sum(Filter1(NumbersGen1(filename)))
}

func NumbersGen1(filename string) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		f, err := os.Open(filename)
		if err != nil {
			return
		}
		defer f.Close()
		fileScanner := bufio.NewScanner(f)
		for fileScanner.Scan() {
			num, err := strconv.Atoi(fileScanner.Text())
			if err != nil {
				continue
			}
			out <- num
		}
	}()
	return out
}

func Filter1(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for num := range in {
			if num%2 == 0 {
				out <- num
			}
		}
	}()
	return out
}

func Sum(in <-chan int) int {
	var res int
	for num := range in {
		res += num
	}
	return res
}
