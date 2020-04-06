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

func helloHandle(c *gee.Context) {
	c.JSON(http.StatusOK, gee.H{
		"hello": "world",
	})
}

func main() {
	r := gee.New()

	r.GET("/index", indexHandle)
	r.GET("/hello", helloHandle)

	r.Run(":8999")
}
