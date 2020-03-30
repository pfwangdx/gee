package main

import (
	"fmt"
	"net/http"
)

// Engine used deal all requests
type Engine struct{}

// Handle http requests
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(w, "url.path = %q\n", req.URL.Path)
	case "/hello":
		helloHandler(w, req)
	default:
		fmt.Fprintf(w, "404 NOT FOUND = %s\n", req.URL.Path)
	}
}

func main() {
	engine := new(Engine)
	http.ListenAndServe(":9999", engine)
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		fmt.Fprintln(w, "header[%q] = %q\n", k, v)
	}
}
