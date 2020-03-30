package engine

import (
	"fmt"
	"net/http"
)

// Engine used deal all requests
type Engine struct{
	router map[string]HandlerFunc
}
// Funcion handler
type HandlerFunc func(http.ResponseWriter, *http.Request)

// Engine Constructor
func New() *Engine {
	return &Engine{ make(map[string]HandlerFunc)}
}

// Engine run process
func (engine *Engine) Run() {
	http.ListenAndServe(":9999", engine)
}


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

func helloHandler(w http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		fmt.Fprintln(w, "header[%q] = %q\n", k, v)
	}
}
