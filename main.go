package main

import (
	"gom"
	"log"
	"net/http"
)

func main() {
	r := gom.Default()
	r.GET("/", func(c *gom.Context) {
		c.String(http.StatusOK, "Hello ma5hr00m\n")
	})
	// index out of range for testing Recovery()
	r.GET("/panic", func(c *gom.Context) {
		names := []string{"ma5hr00m"}
		c.String(http.StatusOK, names[100])
	})

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Server start failed: %s", err.Error())
	}
}
