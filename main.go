package main

import (
	"net/http"

	"aee"
)

func main() {
	r := aee.New()
	r.GET("/index", func(c *aee.Context) {
		c.HTML(http.StatusOK, "<h1>Index Page</h1>")
	})
	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *aee.Context) {
			c.HTML(http.StatusOK, "<h1>Hello Aee</h1>")
		})

		v1.GET("/hello", func(c *aee.Context) {
			// expect /hello?name=aliaxy
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}
	v2 := r.Group("/v2")
	{
		v2.GET("/hello/:name", func(c *aee.Context) {
			// expect /hello/aliaxy
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
		v2.POST("/login", func(c *aee.Context) {
			c.JSON(http.StatusOK, aee.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})

	}

	r.Run(":9999")
}
