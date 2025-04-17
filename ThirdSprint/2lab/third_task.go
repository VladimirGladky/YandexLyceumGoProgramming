package main

import (
	"encoding/csv"
	"os"
)

func ReadCSV(file string) (<-chan []string, error) {
	out := make(chan []string)
	File, err := os.Open(file)
	f := csv.NewReader(File)
	if err != nil {
		return nil, err
	}
	go func() {
		defer File.Close()
		defer close(out)
		for {
			record, err := f.Read()
			if err != nil {
				return
			}
			out <- record
		}
	}()

	return out, nil
}
