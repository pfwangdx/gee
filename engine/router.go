package engine

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

