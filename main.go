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
		r.LoadHTMLGlob("views/*.html")
		r.NoRoute(func(c *gin.Context) {
			c.HTML(http.StatusNotFound, "404.html", gin.H{"msg": "Page NOT FOUND"})
		})
	}

	r.GET("/", ctrl.Index)
	r.GET("/login", ctrl.Login)
	r.GET("/user/:name",ctrl.SessionFilter, ctrl.GetUserByName)

	v2 := r.Group("/v2")
	{
		v2.GET("/user/:name", func(c *gin.Context) {
			name := c.Param("name")
			c.String(http.StatusOK, "v2 Hello %s", name)
		})

		v2.GET("/profile/:id", func(c *gin.Context) {
			id := c.Param("id")
			c.String(http.StatusOK, "v2 id  %s", id)
		})
	}

	r.Run(":8000")
}
