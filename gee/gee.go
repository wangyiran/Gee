package gee

import (
	"net/http"
)

type HandlerFunc func(c *Context)

type Engine struct {
	r *Router
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := newContext(w, r)
	e.r.Handle(c)
}

func New() *Engine {
	return &Engine{
		r: newRouter(),
	}
}

func (e *Engine) addRoute(method string, path string, handler HandlerFunc) {
	e.r.addRoute(method, path, handler)
}

func (e *Engine) GET(path string, handler HandlerFunc) {
	e.addRoute("GET", path, handler)
}

func (e *Engine) POST(path string, handler HandlerFunc) {
	e.addRoute("POST", path, handler)
}

func (e *Engine) RUN(port string) error {
	return http.ListenAndServe(port, e)
}
