package engine

import (
	"fmt"
	"net/http"
)

// Engine used deal all requests
type Engine struct{
	router *router
}
// HandlerFunc definition
type HandlerFunc func(http.ResponseWriter, *http.Request)

// New Engine
func New() *Engine {
	return &Engine{ newRouter() }
}

// Run engine process
func (engine *Engine) Run() (err error) {
	return http.ListenAndServe(":9999", engine)
}

// AddRoute about the handler func
func (engine *Engine) AddRoute(method string, pattern string, handler HandlerFunc) {
	engine.router.addRoute(method, pattern, handler)
}

// GET handler
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.AddRoute("GET", pattern, handler)
}

// POST handler
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.AddRoute("POST", pattern, handler)
}


// Handle http requests
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path

	routes := engine.router.list()
	if handler, ok := routes[key]; ok {
		handler(w, req)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND = %s\n", req.URL.Path)
	}
}