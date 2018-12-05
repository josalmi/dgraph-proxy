package main

import (
	"os"
	"net/http"
	"net/url"
	"net/http/httputil"

	"golang.org/x/text/encoding/charmap"
)

func main() {
	decoder := charmap.ISO8859_1.NewDecoder()

	targetHost, err := url.Parse(os.Getenv("DGRAPH_HOST"))
	if err != nil {
		panic(err)
	}

	varsHeaderKey := "x-dgraph-vars"
        director := func(req *http.Request) {
                req.URL.Scheme = targetHost.Scheme
                req.URL.Host = targetHost.Host
		vars, err := decoder.String(req.Header.Get(varsHeaderKey))
		if err == nil {
			req.Header.Set(varsHeaderKey, vars)
		}
        }
        proxy := &httputil.ReverseProxy{Director: director}
        http.ListenAndServe(":8080", proxy)
}

