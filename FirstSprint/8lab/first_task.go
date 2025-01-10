package main

//
//import (
//	"fmt"
//	"log"
//	"net/http"
//	"regexp"
//	"time"
//)
//
//var (
//	IsLetter = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString
//)
//
//func HelloHandler(w http.ResponseWriter, r *http.Request) {
//	name := r.URL.Query().Get("name")
//	if IsLetter(name) == true && name != "" {
//		w.WriteHeader(http.StatusOK)
//		fmt.Fprintf(w, "Hello, %s!", name)
//		return
//	}
//}
//
//func Sanitize(next http.HandlerFunc) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		name := r.URL.Query().Get("name")
//		if IsLetter(name) != true && name != "" {
//			w.WriteHeader(http.StatusOK)
//			fmt.Fprintf(w, "Hello, dirty hacker!")
//			return
//		}
//		next(w, r)
//	}
//}
//
//func SetDefaultName(next http.HandlerFunc) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		name := r.URL.Query().Get("name")
//		if name == "" {
//			w.WriteHeader(http.StatusOK)
//			fmt.Fprintf(w, "Hello, stranger!")
//			return
//		}
//		next(w, r)
//	}
//}
//
//func StartServer(timeout time.Duration) {
//	http.HandleFunc("/hello", SetDefaultName(Sanitize(HelloHandler)))
//	if err := http.ListenAndServe(":8080", nil); err != nil {
//		log.Fatal(err)
//	}
//}
