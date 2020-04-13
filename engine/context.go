package engine

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Context struct {
	response http.ResponseWriter
	req      *http.Request
	params   map[string]string
}

func (c *Context) Params(key string) string {
	value, _ := c.params[key]
	return value
}

func (c *Context) PostForm(key string) string {
	return c.req.FormValue(key)
}

func (c *Context) Query(key string) string {
	return c.req.URL.Query().Get(key)
}

func (c *Context) SetHeader(key string, value string) {
	c.response.Header().Set(key, value)
}

// NewContext Context constructor
func NewContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{w, req, make(map[string]string)}
}

// Output as string
func (c *Context) String(format string, values ...interface{}) {
	c.response.Write([]byte(fmt.Sprintf(format, values...)))
}

// Json output
func (c *Context) Json(obj ...interface{}) {
	c.SetHeader("Content-Type", "application/json")
	encoder := json.NewEncoder(c.response)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.response, err.Error(), 500)
	}
}

// Html output
func (c *Context) Html(format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/html")
	c.response.Write([]byte(fmt.Sprintf(format, values...)))
}
