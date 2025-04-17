package gzipper

import (
	"compress/gzip"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"sync"
)

type Work struct {
	FilePath string
}

func FileNameGen(dir string, pattern *regexp.Regexp) <-chan Work {
	jobs := make(chan Work)
	go func() {
		defer close(jobs)
		filepath.Walk(dir, func(path string, d fs.FileInfo, err error) error {
			if err != nil {
				return nil
			}
			if d.IsDir() {
				return nil
			}
			if pattern != nil && !pattern.MatchString(d.Name()) {
				return nil
			}
			yandexIssue := regexp.MustCompile("(output.txt|input.txt|trim.txt)")
			if yandexIssue.MatchString(d.Name()) {
				return nil
			}
			if filepath.Dir(path) != filepath.Clean(dir) {
				return nil
			}
			jobs <- Work{
				FilePath: path,
			}
			return nil
		})
	}()
	return jobs
}

func compress(jobs <-chan Work) {
	var wg sync.WaitGroup

	for job := range jobs {
		wg.Add(1)

		go func(job Work) {
			defer wg.Done()
			file, err := os.Open(job.FilePath)
			if err != nil {
				return
			}
			defer file.Close()

			outputPath := job.FilePath + ".gz"
			outPutFile, err := os.Create(outputPath)
			if err != nil {
				return
			}

			gzipW := gzip.NewWriter(outPutFile)

			_, err = io.Copy(gzipW, file)

			closeErr := gzipW.Close()
			fileCloseErr := outPutFile.Close()

			if err != nil {
				return
			}
			if closeErr != nil {
				return
			}
			if fileCloseErr != nil {
				return
			}
		}(job)
	}

	wg.Wait()
}
