package client

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequestGet(t *testing.T) {
	requestValues := url.Values{}
	requestValues.Add("hoge", "fuga")
	requestValues.Add("foo", "bar")

	ts := httptest.NewServer(http.HandlerFunc(
		func(_ http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodGet {
				t.Fatal("request method is not GET")
			}

			assert.Equal(t, requestValues, r.URL.Query())
		},
	))
	defer ts.Close()

	res, err := requestGet(ts.URL, requestValues)
	if err != nil {
		t.Fatalf("error occuered, %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		t.Fatalf("unexpected http response status, %d", res.StatusCode)
	}
}

func TestRequestPost(t *testing.T) {
	requestValues := url.Values{}
	requestValues.Add("hoge", "fuga")
	requestValues.Add("foo", "bar")

	ts := httptest.NewServer(http.HandlerFunc(
		func(_ http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost {
				t.Fatal("request method is not POST")
			}

			if ct := r.Header.Get("Content-Type"); ct != "application/x-www-form-urlencoded" {
				t.Fatalf(`Content-Type, expected:"application/x-www-form-urlencoded", actual:%s`, ct)
			}

			if err := r.ParseForm(); err != nil {
				t.Fatalf("ParseForm() error, %v", err)
			}
			assert.Equal(t, requestValues, r.Form)
		},
	))
	defer ts.Close()

	res, err := requestPost(ts.URL, requestValues)
	if err != nil {
		t.Fatalf("error occuered, %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		t.Fatalf("unexpected http response status, %d", res.StatusCode)
	}
}

func TestRequestPostJson(t *testing.T) {
	requestBody := `{"hoge":"fuga","foo":"bar"}`

	ts := httptest.NewServer(http.HandlerFunc(
		func(_ http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost {
				t.Fatal("request method is not POST")
			}

			if ct := r.Header.Get("Content-Type"); ct != "application/json" {
				t.Fatalf("Content-Type, expected:application/json, actual:%s", ct)
			}

			rb, err := ioutil.ReadAll(r.Body)
			if err != nil {
				t.Fatal("read body error")
			}
			assert.Equal(t, []byte(requestBody), rb)
		},
	))
	defer ts.Close()

	res, err := requestPostJson(ts.URL, strings.NewReader(requestBody))
	if err != nil {
		t.Fatalf("error occuered, %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		t.Fatalf("unexpected http response status, %d", res.StatusCode)
	}
}
