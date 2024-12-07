package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

var (
	a int
	b int
	c int
)

func FibonacciHandler(w http.ResponseWriter, r *http.Request) {
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

func StartServer(timeout time.Duration) {
	http.HandleFunc("/", FibonacciHandler)
	//	http.HandleFunc("/metrics/", MetricsHandler)
	http.ListenAndServe(":8080", nil)
}
