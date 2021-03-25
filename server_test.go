package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func setUp() (string, func()) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", isJSON(home))

	server := httptest.NewServer(mux)
	return server.URL, func() {
		server.Close()
	}
}

func TestValidContentType(t *testing.T) {
	server, tearDown := setUp()
	defer tearDown()

	client := &http.Client{}
	req, err := http.NewRequest("GET", server, nil)
	if err != nil {
		t.Fatalf("error while initializing new request: %s", err.Error())
	}
	req.Header.Add("content-type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("error while sending http request: %s", err.Error())
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("response returned a non-successful status code: want 200; got %d", resp.StatusCode)
	}
}

func TestInValidContentType(t *testing.T) {
	server, tearDown := setUp()
	defer tearDown()

	resp, err := http.Get(server)
	if err != nil {
		t.Fatalf("error while sending http request: %s", err.Error())
	}

	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("response returned a successful status code: want %d; got %d", http.StatusBadRequest, resp.StatusCode)
	}
}
