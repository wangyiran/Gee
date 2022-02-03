package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

type Context struct {
	Req    *http.Request
	Writer http.ResponseWriter

	Method string
	Path   string

	StatusCode int
}

func newContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Req:    r,
		Writer: w,
		Method: r.Method,
		Path:   r.URL.Path,
	}
}

func (c *Context) PostForm(key string) (value string) {
	value = c.Req.FormValue(key)
	return
}

func (c *Context) Query(key string) (value string) {
	value = c.Req.URL.Query().Get(key)
	return
}

func (c *Context) SetStatus(status int) {
	c.StatusCode = status
	c.Writer.WriteHeader(status)
}

func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

func (c *Context) String(status int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.SetStatus(status)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

func (c *Context) JSON(status int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.SetStatus(status)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

func (c *Context) Data(status int, data []byte) {
	c.SetStatus(status)
	c.Writer.Write(data)
}

func (c *Context) HTML(status int, html string) {
	c.SetStatus(status)
	c.SetHeader("Content-Type", "text/html")
	c.Writer.Write([]byte(html))
}
