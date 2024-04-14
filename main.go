package main

import (
	"gom"
	"net/http"
)

func main() {
	r := gom.New()
	r.GET("/", func(c *gom.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	r.GET("/hello", func(c *gom.Context) {
		// expect /hello?name=gomktutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.Run(":9999")
}
