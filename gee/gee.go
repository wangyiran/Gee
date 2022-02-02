package gee

import (
	"fmt"
	"net/http"
)

type HandleFunc func(http.ResponseWriter, *http.Request)

type Engine struct {
	RouterTable map[string]HandleFunc
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if handler, ok := e.RouterTable[r.Method+"-"+r.URL.Path]; ok {
		handler(w, r)
	} else {
		fmt.Fprintf(w, "404")
	}
}

func NEW() *Engine {
	return &Engine{
		RouterTable: make(map[string]HandleFunc),
	}
}

func (e *Engine) addRouter(method string, path string, handler HandleFunc) {
	e.RouterTable[method+"-"+path] = handler
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
