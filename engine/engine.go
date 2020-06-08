package engine

import (
	"net/http"
)

// HandlerFunc definition
type HandlerFunc func(*Context)

type(
	// Engine used deal all requests
	Engine struct {
		router *router
		groups []*RouterGroup
	}

	// RouterGroup respresent a router group 
	RouterGroup struct {
		prefix string
		engine *Engine
	}
) 


// New Engine
func New() *Engine {
	engine := &Engine{
		router: newRouter(),
		groups: []*RouterGroup{},
	}
	return engine
}

// Run engine process
func (engine *Engine) Run(port string) (err error) {
	return http.ListenAndServe(port, engine)
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
	context := NewContext(w, req)
	engine.router.handle(context)
}
