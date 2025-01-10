package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

func AnswerHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "The answer is 42")
}

func Authorization(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth == "" {
			http.Error(w, "401 Unauthorized", http.StatusUnauthorized)
			return
		}
		if !strings.HasPrefix(auth, "Basic userid:") {
			http.Error(w, "401 Unauthorized", http.StatusUnauthorized)
			return
		}
		auth = strings.TrimLeft(auth, "Basic userid:")
		if auth == "" {
			http.Error(w, "401 Unauthorized", http.StatusUnauthorized)
			return
		}
		next(w, r)

	}
}

func StartServer2(time time.Duration) {
	http.HandleFunc("/auth", Authorization(AnswerHandler))
	http.ListenAndServe(":8080", nil)
}
