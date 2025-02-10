package main

import (
	"context"
	"io"
	"os"
	"time"
)

func readJSON(ctx context.Context, path string, result chan<- []byte) {
	ctxto, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer func(file *os.File) {
		err1 := file.Close()
		if err1 != nil {

		}
	}(file)
	byteValue, _ := io.ReadAll(file)
	select {
	case <-ctxto.Done():
		return
	default:
		result <- byteValue
	}
}
