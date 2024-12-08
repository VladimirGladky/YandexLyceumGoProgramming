package main

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"
	"time"
)

type RPCHello struct {
	Status string                 `json:"status"`
	Result map[string]interface{} `json:"result"`
}

var (
	IsLetter = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "stranger"
	}
	response := RPCHello{Status: "ok", Result: map[string]interface{}{
		"greetings": "hello",
		"name":      name,
	}}
	jsonBytes, _ := json.Marshal(response)
	w.Write(jsonBytes)
}

func RPC(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				response := RPCHello{
					Status: "error",
					Result: map[string]interface{}{},
				}
				jsonBytes, _ := json.Marshal(response)
				w.Write(jsonBytes)
			}
		}()
		next(w, r)
	}
}

func Sanitize(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if !IsLetter(name) && name != "" {
			panic("invalid name")
		}
		next(w, r)
	}
}

func SetDefaultName(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		next(w, r)
	}
}

func StartServer3(timeout time.Duration) {
	http.HandleFunc("/hello", RPC(SetDefaultName(Sanitize(HelloHandler))))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
