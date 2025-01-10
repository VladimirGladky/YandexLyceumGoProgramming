package main

import "io"

func Contains(r io.Reader, seq []byte) (bool, error) {
	data := make([]byte, 1024)
	bytesRead, err := r.Read(data)
	if err != nil {
		return false, err
	}
	str := data[:bytesRead]
	for i, x := range str {
		if byte(x) == seq[0] && i+len(seq) <= bytesRead {
			if Compare(str[i:i+len(seq)], seq) {
				return true, nil
			}
			continue
		}
	}
	return false, nil
}

func Compare(arr, seq []byte) bool {
	for i, x := range arr {
		if x == seq[i] {
			continue
		} else {
			return false
		}
	}
	return true
}
