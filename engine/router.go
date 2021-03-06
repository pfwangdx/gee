package engine

import (
	"fmt"
	"strings"
)

type router struct {
	roots    map[string]*node
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{
		roots:    make(map[string]*node),
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
		r.roots[method] = &node{}
	}

	Isok := r.roots[method].insert(path, searchPath, 0)
	if !Isok {
		fmt.Println("insert eror")
	}
	r.handlers[key] = handler
	fmt.Println(" >>>> router: ", r.roots)
}

// getRoute used to retrieve the search tree node and param values
func (r *router) getRoute(method string, path string) (*node, map[string]string) {
	searchPath := parsePath(path)
	params := make(map[string]string)
	root, ok := r.roots[method]

	if !ok {
		return nil, nil
	}

	n := root.search(searchPath, 0)
	// The route match
	if n != nil {
		// why
		parts := parsePath(n.pattern)
		// update params
		for index, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = searchPath[index]
			}
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(searchPath[index:], "/")
				break
			}
		}
		return n, params
	}
	return nil, nil
}

// handle the context which contains request info
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
	searchPath := make([]string, 0)
	for _, item := range vs {
		if item != "" {
			searchPath = append(searchPath, item)
			if item[0] == '*' {
				break
			}
		}
	}
	return searchPath
}
