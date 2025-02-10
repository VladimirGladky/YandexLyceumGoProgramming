package main

import (
	"context"
	"io"
	"net/http"
	"time"
)

//type APIResponse struct {
//	Data       string // тело ответа
//	StatusCode int    // код ответа
//}

func fetchAPI(ctx context.Context, url string, timeout time.Duration) (*APIResponse, error) {
	ctxt, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	for {
		select {
		case <-ctxt.Done():
			return nil, context.DeadlineExceeded
		default:
			req, err := http.NewRequestWithContext(ctxt, "GET", url, nil)
			if err != nil {
				return nil, err
			}
			res, err := http.DefaultClient.Do(req)
			if err != nil {
				return nil, err
			}
			body, err := io.ReadAll(res.Body)
			if err != nil {
				return nil, err
			}
			return &APIResponse{Data: string(body), StatusCode: res.StatusCode}, nil
		}
	}
}
