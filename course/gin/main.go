package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Context struct {
	middware []func(c *Context)
	index    int8
}

func (c *Context) Use(f func(c *Context)) {
	c.middware = append(c.middware, f)
}

func (c *Context) Next() {
	c.index++
	c.middware[c.index](c)
}

func (c *Context) Run() {
	c.middware[0](c)
}

func (c *Context) Get(path string, f func(c *Context)) {
	c.middware = append(c.middware, f)
}

func middware1() func(c *Context) {
	return func(c *Context) {
		fmt.Println("middle---1--start")
		c.Next()
		fmt.Println("middle---1--end")
	}
}
func middware2() func(c *Context) {
	return func(c *Context) {
		fmt.Println("middle---2---start")
		c.Next()
		fmt.Println("middle--2---end")
	}
}

func main() {
	c := &Context{}
	c.Use(middware1())
	c.Use(middware2())
	c.Get("/", func(c *Context) {
		fmt.Println("ok")
	})
	c.Run()

	r := gin.Default()
	r.Use()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
