package main

import (
	"net/http"

	"aee"
)

func main() {
	r := aee.New()
	r.GET("/", func(c *aee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Aee</h1>")
	})

	r.GET("/hello", func(c *aee.Context) {
		// expect /hello?name=aliaxy
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.GET("/hello/:name", func(c *aee.Context) {
		// expect /hello/aliaxy
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.GET("/assets/*filepath", func(c *aee.Context) {
		c.JSON(http.StatusOK, aee.H{"filepath": c.Param("filepath")})
	})

	r.Run(":9999")
}
