package main

import (
	"fmt"
	"log"
	"net/http"
)

func logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Request received")
		/*
			any code written here will be executed before
			passing the request object to the next middleware
		*/
		next.ServeHTTP(w, r)
		log.Println("Response sent")
	}
}

func main() {
	port := 5000

	http.HandleFunc("/", logger(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello Gophers!\n")
	}))

	log.Printf("Running on http://127.0.0.1:%d 🚀🚀", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
