package main

import (
	"io"
	"net/http"
	"time"
)

func readSourceHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://localhost:8081/provideData")
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

func StartServer(maxTimeout time.Duration) {
	http.Handle("/readSource", http.TimeoutHandler(http.HandlerFunc(readSourceHandler), maxTimeout, "Timeout"))
	http.ListenAndServe(":8080", nil)
}
