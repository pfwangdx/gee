package engine

import "fmt"

type router struct {
	r map[string]HandlerFunc
}


func newRouter() *router {
	return &router{ make(map[string]HandlerFunc)}
}

// AddRoute about the handler func
func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	r.r[key] = handler
}

func (r *router) list() map[string]HandlerFunc {
	return r.r
}

func (r *router) handle(context *Context) {
	key := context.req.Method + "-" + context.req.URL.Path

	if handler, ok := r.r[key]; ok {
		handler(context.response, context.req)
	} else {
		fmt.Fprintf(context.response, "404 NOT FOUND = %s\n", context.req.URL.Path)
	}
}

