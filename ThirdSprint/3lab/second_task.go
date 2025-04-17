package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
)

type CounterWorker interface {
	ProcessFiles(files ...string) error // для запуска обработки файлов
	ProcessReader(r io.Reader) error    // для подсчёта слов в одном файле
}

type WordCounter struct {
	wordsCount map[string]int // здесь должен быть список слов с указанием количества повторений во всех файлах.
	// можно добавлять другие поля
	mu sync.Mutex
}

func (w *WordCounter) ProcessFiles(files ...string) error {
	var wg sync.WaitGroup
	for _, file := range files {
		wg.Add(1)
		go func(file string) {
			defer wg.Done()
			f, err := os.Open(file)
			if err != nil {
				return
			}
			err = w.ProcessReader(f)
			if err != nil {
				return
			}
		}(file)
	}
	wg.Wait()
	return nil
}

func (w *WordCounter) ProcessReader(r io.Reader) error {
	for {
		var word string
		_, err := fmt.Fscan(r, &word)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		word = strings.ToLower(word)
		if word == "" {
			continue
		}
		w.mu.Lock()
		w.wordsCount[word]++
		w.mu.Unlock()
	}
	return nil
}
