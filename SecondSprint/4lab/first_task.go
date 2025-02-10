package main

import (
	"bytes"
	"context"
	"errors"
	"io"
)

func Contains(ctx context.Context, r io.Reader, seq []byte) (bool, error) {
	if len(seq) == 0 {
		return false, errors.New("последовательность не может быть длины 0")
	}
	data := make([]byte, 1024)
	var buf []byte
	for {
		select {
		case <-ctx.Done():
			return false, ctx.Err()
		default:
			_, err := r.Read(buf)
			if err != nil && err != io.EOF {
				return false, err
			}
			data = append(data, buf...)
			if bytes.Contains(data, seq) {
				return true, nil
			}
			if err == io.EOF {
				return false, nil
			}
			if len(data) > len(seq)-1 {
				data = data[1:]
			}
		}
	}
}
