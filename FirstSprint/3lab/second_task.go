package main

import (
	"io"
)

func ReadString(r io.Reader) (string, error) {
	data := make([]byte, 1024)
	bytesRead, err := r.Read(data)
	if err != nil && err != io.EOF {
		return "", err
	}
	return string(data[:bytesRead]), nil
}
