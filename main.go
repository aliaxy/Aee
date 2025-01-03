package main

import (
	"net/http"

	"aee"
)

func main() {
	r := aee.Default()
	r.GET("/", func(c *aee.Context) {
		c.String(http.StatusOK, "Hello aliaxy\n")
	})
	// index out of range for testing Recovery()
	r.GET("/panic", func(c *aee.Context) {
		names := []string{"aliaxy"}
		c.String(http.StatusOK, names[100])
	})

	r.Run(":9999")
}
