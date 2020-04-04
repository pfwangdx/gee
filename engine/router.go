package engine

import (
	"fmt"
	"strings"
)

type router struct {
	roots map[string]*node
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{
		roots: make(map[string]*node),
		handlers: make(map[string]HandlerFunc),
	}
}

// AddRoute about the handler func
func (r *router) addRoute(method string, path string, handler HandlerFunc) {
	// search_path := parsePath(path)
	key := method + "-" + path

	// roots operaions
	
	r.handlers[key] = handler
}

func (r *router) getRoute(method string, path string) {

}

// func (r *router) list() map[string]HandlerFunc {
// 	return r.r
// }

func (r *router) handle(context *Context) {
	params := parsePath(context.req.URL.Path)

	key := context.req.Method + "-" + context.req.URL.Path

	if handler, ok := r.handlers[key]; ok {
		handler(context)
	} else {
		fmt.Fprintf(context.response, "404 NOT FOUND = %s\n", context.req.URL.Path)
	}
}

// parsePath to return search path of the real request path 
func parsePath(path string) []string {
	vs := strings.Split(path, "/")
	search_path := make([]string, 0)
	for _, item := range vs {
		if item != "" {
			search_path = append(search_path, item)
			if item == "*" {
				break
			}
		}
	}
	return search_path
}
