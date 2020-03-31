package engine

import (
	"net/http"
)

type Context struct {
	response http.ResponseWriter
	req *http.Request
}

// NewContext Context constructor
func NewContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{w, req}
}




