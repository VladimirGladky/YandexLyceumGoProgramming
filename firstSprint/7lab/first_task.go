package main

import (
	"fmt"
	"net/http"
	"regexp"
)

func StrangerHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	var IsLetter = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString
	if name == "" {
		fmt.Fprintf(w, "hello stranger")
	} else if IsLetter(name) != true {
		fmt.Fprintf(w, "hello dirty hacker")
	} else {
		fmt.Fprintf(w, "hello %s", name)
	}
}
