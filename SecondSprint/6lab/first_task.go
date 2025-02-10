package main

import (
	"net/http"
)

func FetchURL(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		return "Failed to fetch"
	}
	if resp.StatusCode == 200 {
		return "Successfully fetched"
	} else {
		return "Failed to fetch"
	}
}
