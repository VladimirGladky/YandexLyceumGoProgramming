package main

import (
	"bufio"
	"os"
)

func LineByNum(inputFilename string, lineNum int) string {
	f, err := os.Open(inputFilename)
	defer f.Close()
	if err != nil {
		return ""
	}
	var i int
	fileScanner := bufio.NewScanner(f)
	for fileScanner.Scan() {
		if i == lineNum {
			return fileScanner.Text()
		}
		i++
	}
	return ""
}
