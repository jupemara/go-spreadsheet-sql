package util

import (
	"net/http"

	"golang.org/x/net/http2"
)

func NewClient() (*http.Client, error) {
	t := &http.Transport{}
	if err := http2.ConfigureTransport(t); err != nil {
		return nil, err
	}
	return &http.Client{Transport: t}, nil
}
