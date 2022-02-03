package gee

import (
	"net/http"
)

type HandleFunc func(*Context)

type Engine struct {
	r router
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := newContext(w, r)
	e.r.handle(c)

}

func NEW() *Engine {
	return &Engine{
		r: *newRouter(),
	}
}

func (e *Engine) addRouter(method string, path string, handler HandleFunc) {
	e.r.addRouter(method, path, handler)
}

func (e *Engine) GET(path string, handler HandleFunc) {
	e.addRouter("GET", path, handler)
}

func (e *Engine) POST(path string, handler HandleFunc) {
	e.addRouter("POST", path, handler)
}

func (e *Engine) RUN(port string) error {
	return http.ListenAndServe(port, e)
}
