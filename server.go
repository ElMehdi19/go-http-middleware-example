package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		/*
			log format: HTTP_METHOD /PATH REMOTE_ADDRESS
			example: GET /home 127.0.0.1
		*/
		method := r.Method
		path := r.URL.Path
		remoteAddr := strings.Split(r.RemoteAddr, ":")[0]
		log.Printf("%s %s %s\n", method, path, remoteAddr)
		next.ServeHTTP(w, r)
	}
}

func isJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		contentType := r.Header.Get("Content-Type")
		if contentType != "application/json" {
			http.Error(w, "content-type header must be application/json", http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Howdy Gophers 👋\n")
}

func main() {
	port := 5000

	http.HandleFunc("/", logger(isJSON(home)))

	log.Printf("Running on http://127.0.0.1:%d 🚀🚀", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
