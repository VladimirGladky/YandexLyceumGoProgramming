package main

import "os"

func ModifyFile(filename string, pos int, val string) {
	f, _ := os.Open(filename)
	defer f.Close()
	buffer := make([]byte, 1024)
	n, _ := f.Read(buffer)
	buffer = buffer[:n]
	for i, _ := range buffer {
		if i == pos {
			for _, x := range val {
				buffer[i] = byte(x)
				i++
			}
			break
		}
	}
	os.WriteFile(filename, buffer, 0666)
}
