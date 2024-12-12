package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func RequestHandler(w http.ResponseWriter, r *http.Request) {
	query, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Bad request")
		return
	}
	name := query.Get("name")
	if len(name) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "You must supply a name")
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello %s", name)
}
func main() {
	http.HandleFunc("/greet", RequestHandler)
	log.Fatal(http.ListenAndServe(":3030", nil))
}
