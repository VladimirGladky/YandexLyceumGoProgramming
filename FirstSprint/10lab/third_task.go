package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Hello struct {
	Name string `json:"name"`
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "stranger"
	}
	response := Hello{Name: name}
	jsonBytes, _ := json.Marshal(response)
	w.Write(jsonBytes)
	log.Printf("{\"name\":\"%s\"}", name)
}

func StartServer(time time.Duration) {
	http.HandleFunc("/hello", HelloHandler)
	http.ListenAndServe(":8080", nil)
}
