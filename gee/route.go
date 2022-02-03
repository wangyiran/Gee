package gee

import "net/http"

type router struct {
	routerTable map[string]HandleFunc
}

func newRouter() *router {
	return &router{
		routerTable: make(map[string]HandleFunc),
	}
}

func (r *router) addRouter(method string, path string, handler HandleFunc) {
	r.routerTable[method+"-"+path] = handler
}

func (r *router) handle(c *Context) {
	if handler, ok := r.routerTable[c.Method+"-"+c.Path]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 :%s", c.Path)
	}
}
