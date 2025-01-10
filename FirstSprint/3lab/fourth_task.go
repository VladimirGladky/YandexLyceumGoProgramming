package main

import "io"

func Copy(r io.Reader, w io.Writer, n uint) error {
	data := make([]byte, 1024)
	bytesRead, err := r.Read(data)
	if err != nil && err != io.EOF {
		return err
	}
	var err1 error
	if int(n) < bytesRead {
		_, err1 = w.Write(data[:n])
	} else {
		_, err1 = w.Write(data[:bytesRead])
	}
	if err1 != io.EOF {
		return err1
	}
	return nil
}
