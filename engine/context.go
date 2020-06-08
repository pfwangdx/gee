package engine

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Context struct {
	Writer http.ResponseWriter
	Req      *http.Request
	Params   map[string]string
	handlers []HandlerFunc
	StatusCode int
	index int
}

// NewContext Context constructor
func NewContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w, 
		Req: req, 
		Params: make(map[string]string),
		index: -1,
	}
}

func (c *Context) Next() {
	c.index++
	s:=len(c.handlers)
	for ; c.index < s; c.index++ {
		c.handlers[c.index](c)
	}
}

func (c *Context) Param(key string) string {
	value, _ := c.Params[key]
	return value
}

func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

// Output as string
func (c *Context) String(format string, values ...interface{}) {
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

// Json output
func (c *Context) Json(obj ...interface{}) {
	c.SetHeader("Content-Type", "application/json")
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

// Html output
func (c *Context) Html(format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/html")
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}
