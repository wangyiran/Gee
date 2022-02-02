package main

import (
	"fmt"
	"net/http"

	gee "./gee"
)

func main() {
	hs := gee.NEW()
	hs.GET("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "path:%s", r.URL)
	})
	hs.GET("/hehe", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "hehehe")
	})
	hs.RUN(":9999")
}
