package client

import (
	"io"
	"net/http"
	"net/url"
	"strings"
)

func requestGet(targetURL string, values url.Values) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, targetURL, nil)
	if err != nil {
		return nil, err
	}
	req.URL.RawQuery = values.Encode()

	return http.DefaultClient.Do(req)
}

func requestPost(targetURL string, values url.Values) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodPost, targetURL, strings.NewReader(values.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	return http.DefaultClient.Do(req)
}

func requestPostJson(targetURL string, values io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodPost, targetURL, values)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")

	return http.DefaultClient.Do(req)
}
