package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"aee"
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
	r := aee.New()
	r.Use(aee.Logger())
	r.SetFuncMap(template.FuncMap{
		"FormatAsDate": FormatAsDate,
	})
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./static")

	stu1 := &student{Name: "aliaxy", Age: 20}
	stu2 := &student{Name: "Jack", Age: 22}
	r.GET("/", func(c *aee.Context) {
		c.HTML(http.StatusOK, "css.tmpl", nil)
	})
	r.GET("/students", func(c *aee.Context) {
		c.HTML(http.StatusOK, "arr.tmpl", aee.H{
			"title":  "aee",
			"stuArr": [2]*student{stu1, stu2},
		})
	})

	r.GET("/date", func(c *aee.Context) {
		c.HTML(http.StatusOK, "custom_func.tmpl", aee.H{
			"title": "aee",
			"now":   time.Now(),
		})
	})

	r.Run(":9999")
}
