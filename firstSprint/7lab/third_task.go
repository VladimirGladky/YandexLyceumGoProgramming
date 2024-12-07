package main

import (
	//"fmt"
	"net/http"
	//"strconv"
)
//
//var (
//	a            int
//	b            int
//	c            int
//	requestCount int
//)
//
//func FibonacciHandler(w http.ResponseWriter, r *http.Request) {
//	requestCount++
//	if a == 0 && b == 0 {
//		fmt.Fprintf(w, strconv.Itoa(a))
//		b = 777
//	} else if a == 0 && b == 777 {
//		fmt.Fprintf(w, strconv.Itoa(1))
//		b = 1
//	} else {
//		fmt.Fprintf(w, strconv.Itoa(a+b))
//		c = b
//		b = a + b
//		a = c
//	}
//}

func MetricsHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "rpc_duration_milliseconds_count %d", requestCount)
}

//func StartServer(timeout time.Duration) {
//	http.HandleFunc("/", FibonacciHandler)
//	http.HandleFunc("/metrics/", MetricsHandler)
//	http.ListenAndServe(":8080", nil)
//}
