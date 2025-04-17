package main

import (
	"bufio"
	"os"
	"strconv"
)

func NumbersGen2(filename string) <-chan int {
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
