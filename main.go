package main

import (
	"gom"
	"log"
	"net/http"
	"time"
)

func onlyForV2() gom.HandlerFunc {
	return func(c *gom.Context) {
		t := time.Now()
		// if a server error occurred
		c.Fail(500, "Internal Server Error")
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func main() {
	r := gom.New()
	r.Use(gom.Logger())
	r.GET("/", func(c *gom.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	v2 := r.Group("/v2")
	v2.Use(onlyForV2())
	{
		v2.GET("/hello/:name", func(c *gom.Context) {
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
	}

	r.Run(":9090")
}
