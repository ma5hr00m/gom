package main

import (
	"fmt"
	"gom"
	"html/template"
	"net/http"
	"time"
)

type student struct {
	Name string
	Age  int8
}

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func main() {
	r := gom.New()
	r.Use(gom.Logger())
	r.SetFuncMap(template.FuncMap{
		"FormatAsDate": FormatAsDate,
	})
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./static")

	stu1 := &student{Name: "Tom", Age: 20}
	stu2 := &student{Name: "Jack", Age: 22}
	r.GET("/", func(c *gom.Context) {
		c.HTML(http.StatusOK, "css.tmpl", nil)
	})
	r.GET("/students", func(c *gom.Context) {
		c.HTML(http.StatusOK, "arr.tmpl", gom.H{
			"title":  "gom",
			"stuArr": [2]*student{stu1, stu2},
		})
	})

	r.GET("/date", func(c *gom.Context) {
		c.HTML(http.StatusOK, "custom_func.tmpl", gom.H{
			"title": "gom",
			"now":   time.Date(2024, 4, 15, 0, 0, 0, 0, time.UTC),
		})
	})

	r.Run(":9999")
}
