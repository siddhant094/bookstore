package main

import (
	"bytes"
	"net/http"
	"testing"
	"time"

	"gofr.dev/pkg/gofr/request"
)

func TestIntegration(t *testing.T) {
	go main()
	time.Sleep(3 * time.Second)

	tests := []struct {
		desc       string
		method     string
		endpoint   string
		statusCode int
		body       []byte
	}{
		{"get all books", http.MethodGet, "book", http.StatusOK, nil},
		{"get books by ID", http.MethodGet, "book/1", http.StatusOK, nil},
		{"post books", http.MethodPost, "book", http.StatusCreated, []byte(`{  
			"Name":"1984",
			"Author":"George Orwell",
			"Publication":"Penguin"
			}`),
		},
		// {"Update books", http.MethodPut, "book/1", http.StatusOK, []byte(`{  
		// 	"Name":"1984",
		// 	"Author":"George Orwell",
		// 	"Publication":"Penguin"
		// }`),},
		// {"delete book", http.MethodDelete, "book/1", http.StatusNoContent, nil},
	}

	for i, tc := range tests {
		req, _ := request.NewMock(tc.method, "http://localhost:9000/"+tc.endpoint, bytes.NewBuffer(tc.body))

		c := http.Client{}

		resp, err := c.Do(req)
		if err != nil {
			t.Errorf("TEST[%v] Failed.\tHTTP request encountered Err: %v\n%s", i, err, tc.desc)
			continue
		}

		if resp.StatusCode != tc.statusCode {
			t.Errorf("TEST[%v] Failed.\tExpected %v\tGot %v\n%s", i, tc.statusCode, resp.StatusCode, tc.desc)
		}

		_ = resp.Body.Close()
	}
}