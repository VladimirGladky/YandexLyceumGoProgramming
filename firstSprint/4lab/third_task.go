package main

import "os"

func CopyFilePart(inputFilename, outFileName string, startpos int) error {
	f, err := os.Open(inputFilename)
	defer f.Close()
	if err != nil {
		return err
	}
	f.Seek(int64(startpos), 0)
	buffer := make([]byte, 1024)
	n, err1 := f.Read(buffer)
	if err1 != nil {
		return err1
	}
	err = os.WriteFile(outFileName, buffer[:n], 0666)
	if err != nil {
		return err
	}
	return nil
}
