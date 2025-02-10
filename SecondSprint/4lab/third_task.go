package main

import (
	"context"
	"io"
	"net/http"
	"sync"
	"time"
)

type APIResponse struct {
	URL        string // запрошенный URL
	Data       string // тело ответа
	StatusCode int    // код ответа
	Err        error  // ошибка, если возникла
}

func FetchAPI(ctx context.Context, urls []string, timeout time.Duration) []*APIResponse {
	var mu sync.Mutex
	var res []*APIResponse
	var wg sync.WaitGroup
	for _, x := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			ctxt, cancel := context.WithTimeout(ctx, timeout)
			defer cancel()
			req, err := http.NewRequestWithContext(ctxt, "GET", url, nil)
			if err != nil {
				mu.Lock()
				res = append(res, &APIResponse{URL: url, Err: err})
				mu.Unlock()
				return
			}
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				mu.Lock()
				res = append(res, &APIResponse{URL: url, Err: err})
				mu.Unlock()
				return
			}
			defer resp.Body.Close()
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				mu.Lock()
				res = append(res, &APIResponse{URL: url, StatusCode: resp.StatusCode, Err: err})
				mu.Unlock()
				return
			}
			mu.Lock()
			res = append(res, &APIResponse{url, string(body), resp.StatusCode, nil})
			mu.Unlock()
		}(x)
	}
	wg.Wait()
	return res
}
