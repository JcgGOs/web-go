package main

import (
	"net/http"

	"bloom.io/ctrl"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	{
		r.StaticFile("/robots.txt", "./static/robots.txt")
		r.StaticFile("/favicon.ico", "./static/favicon.ico")
		r.Static("/static", "./static")
		r.LoadHTMLGlob("views/*")
		r.NoRoute(func(c *gin.Context) {
			c.HTML(http.StatusNotFound, "404.html", gin.H{"msg": "Page NOT FOUND"})
		})
	}

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "welcome",
		})
	})

	r.GET("/welcome", ctrl.Index)
	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	v2 := r.Group("/v2")
	{
		v2.GET("/user/:name", func(c *gin.Context) {
			name := c.Param("name")
			c.String(http.StatusOK, "v2 Hello %s", name)
		})
	}

	r.Run(":8000")
}
