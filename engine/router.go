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
	searchPath := parsePath(path)
	
	key := method + "-" + path
	// roots operaions
	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &node {}
	}

	r.roots[method].insert(path, searchPath, 0)
	r.handlers[key] = handler
}

func (r *router) getRoute(method string, path string) (*node, map[string]string) {
	searchPath := parsePath(path)
	params := make(map[string]string)
	node, ok := r.roots[method]

	if !ok {
		return nil, nil
	}

	n := node.search(path, searchPath, 0)
	// The route match
	if n != nil {
		// why
		parts := parsePath(n.pattern)
		// update params
		for index, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = searchPath[index]
			}
			if part[0] == '*' {
				params[part[1:]] = strings.Join(searchPath[index:], "/")
				break
			}
		}
		return n, params
	}
	return nil, nil
}

func (r *router) handle(context *Context) {
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
