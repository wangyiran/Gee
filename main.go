package main

import (
	"fmt"
	"net/http"

	gee "./gee"
)

func main() {
	hs := gee.NEW()
	hs.GET("/", func(c *gee.Context) {
		fmt.Fprintf(c.Writer, "path:%s", c.Req.URL)
	})
	hs.GET("/hehe", func(c *gee.Context) {
		fmt.Fprintf(c.Writer, "hehehe")
	})
	hs.POST("/login", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	hs.RUN(":9999")
}
