package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func ExtractLog(inputFileName string, start, end time.Time) ([]string, error) {
	f, _ := os.Open(inputFileName)
	var res []string
	fileScanner := bufio.NewScanner(f)
	for fileScanner.Scan() {
		str := fileScanner.Text()
		s := strings.Split(str, " ")
		s = strings.Split(s[0], ".")
		year, _ := strconv.Atoi(s[2])
		m, _ := strconv.Atoi(s[1])
		day, _ := strconv.Atoi(s[0])
		str1 := time.Date(year, time.Month(m), day, 0, 0, 0, 0, time.UTC)
		if (str1.Compare(start) == 1 || str1.Compare(start) == 0) && (str1.Compare(end) == -1 || str1.Compare(end) == 0) {
			res = append(res, str)
		}
	}
	if len(res) == 0 {
		return nil, fmt.Errorf("ни одна строка не попала в указанный диапазон")
	}
	return res, nil
}
