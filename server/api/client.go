package api

import "net/http"

var client http.Client

func init() {
	c := http.Client{
		Transport:     nil,
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       0,
	}

	client = c
}
