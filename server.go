package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := 5000

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello Gophers!\n")
	})

	log.Printf("Running on http://127.0.0.1:%d 🚀🚀", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
