package main

import (
	"os"
)

func ReadContent(filename string) string {
	data, err := os.ReadFile(filename)
	if err != nil {
		return ""
	}
	return string(data)
}
