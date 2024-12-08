package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

var (
	a            int
	b            int
	c            int
	requestCount int
)

func FibonacciHandler(w http.ResponseWriter, r *http.Request) {
	requestCount++
	if a == 0 && b == 0 {
		fmt.Fprintf(w, strconv.Itoa(a))
		b = 777
	} else if a == 0 && b == 777 {
		fmt.Fprintf(w, strconv.Itoa(1))
		b = 1
	} else {
		fmt.Fprintf(w, strconv.Itoa(a+b))
		c = b
		b = a + b
		a = c
	}
}

func MetricsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "rpc_duration_milliseconds_count %d", requestCount)
}

func Metrics(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	}
}

func StartServer1(timeout time.Duration) {
	http.HandleFunc("/", Metrics(FibonacciHandler))
	http.HandleFunc("/metrics/", Metrics(MetricsHandler))
	http.ListenAndServe(":8080", nil)
}
