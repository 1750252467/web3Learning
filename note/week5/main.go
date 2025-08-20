package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})

	r.Any("/login", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})

	grp1 := r.Group("/grp1")
	{
		grp1.GET("/a", func(c *gin.Context) {
			c.String(200, "Hello, Group 1 - A")
		})
	}
	grp2 := r.Group("/grp2")
	{
		grp2.GET("/b", func(c *gin.Context) {
			c.String(200, "Hello, Group 2 - B")
		})
	}

	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
