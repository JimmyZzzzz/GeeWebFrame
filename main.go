package main

import (
	"gee"
	"net/http"
)

func indexHandle(c *gee.Context) {
	c.JSON(http.StatusOK, gee.H{
		"hello": "index",
	})
}
func v1indexHandle(c *gee.Context) {
	c.JSON(http.StatusOK, gee.H{
		"hello": "index-v1",
	})
}

func helloHandle(c *gee.Context) {
	c.JSON(http.StatusOK, gee.H{
		"hello": "world",
	})
}

func v1helloHandle(c *gee.Context) {
	c.JSON(http.StatusOK, gee.H{
		"hello": "world-v1",
	})
}

func main() {
	r := gee.New()

	r.GET("/index", indexHandle)
	r.GET("/hello", helloHandle)

	v1 := r.Group("/v1")

	v1.GET("/index", v1indexHandle)
	v1.GET("/hello", v1helloHandle)

	r.Run(":8999")
}
