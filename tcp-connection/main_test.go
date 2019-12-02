package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	noBodyHandler = http.HandlerFunc(func(_ http.ResponseWriter, _ *http.Request) {
	})
	hasBodyHandler = http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintf(w, `{"name":"hoge"}`)
	})
)

func BenchmarkRequest(b *testing.B) {
	tests := []struct {
		name        string
		server      http.HandlerFunc
		requestFunc func(url string)
	}{
		{
			name:   "no_body_without_discard",
			server: noBodyHandler,
			requestFunc: func(url string) {
				res, _ := http.Get(url)
				defer res.Body.Close()
			},
		},
		{
			name:   "no_body_with_discard",
			server: noBodyHandler,
			requestFunc: func(url string) {
				res, _ := http.Get(url)
				io.Copy(ioutil.Discard, res.Body)
				defer res.Body.Close()
			},
		},
		{
			name:   "has_body_without_discard",
			server: hasBodyHandler,
			requestFunc: func(url string) {
				res, _ := http.Get(url)
				defer res.Body.Close()
			},
		},
		{
			name:   "has_body_with_discard",
			server: hasBodyHandler,
			requestFunc: func(url string) {
				res, _ := http.Get(url)
				io.Copy(ioutil.Discard, res.Body)
				defer res.Body.Close()
			},
		},
	}

	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			ts := httptest.NewServer(tt.server)
			defer ts.Close()

			for i := 0; i < b.N; i++ {
				tt.requestFunc(ts.URL)
			}
		})

	}
}
