package gee

import (
	"net/http"
	"strings"
)

type Router struct {
	RouterTable map[string]HandlerFunc
	roots       map[string]*Node
}

func newRouter() *Router {
	return &Router{
		RouterTable: make(map[string]HandlerFunc),
		roots:       make(map[string]*Node),
	}
}

func parsePattern(pattern string) []string {
	res := strings.Split(pattern, "/")
	parts := make([]string, 0)
	parts = append(parts, "/")
	for _, part := range res {
		if len(part) > 0 {
			parts = append(parts, part)
			if part[0] == '*' {
				break
			}
		}
	}
	return parts
}

func (r *Router) addRoute(method string, path string, handler HandlerFunc) {
	key := method + "-" + path
	parts := parsePattern(path)
	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &Node{
			part: " ",
		}
	}
	r.roots[method].insertRoute(path, parts, 0)
	r.RouterTable[key] = handler

}

func (r *Router) getRoute(method string, path string) (*Node, map[string]string) {
	searchParts := parsePattern(path)
	root, ok := r.roots[method]
	if !ok {
		return nil, nil
	}
	params := make(map[string]string)

	node := root.searchRoute(searchParts, 0)

	if node != nil {
		parts := parsePattern(node.pattern)
		for index, value := range parts {
			if value[0] == ':' {
				params[value[1:]] = searchParts[index]
			}
			if value[0] == '*' && len(value) > 1 {
				params[value[1:]] = strings.Join(searchParts[index:], "/")
			}
		}
		return node, params
	}
	return nil, nil
}

func (r *Router) Handle(c *Context) {
	node, params := r.getRoute(c.Method, c.Path)
	if node != nil {
		key := c.Method + "-" + node.pattern
		c.Params = params
		r.RouterTable[key](c)
	} else {
		c.String(http.StatusNotFound, "404 not found :%s", c.Req.URL)
	}

}
