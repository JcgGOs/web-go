package main

import (
	"encoding/gob"
	"net/http"

	"bloom.io/middleware"

	"bloom.io/model"

	"bloom.io/ctrl"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	{
		store := memstore.NewStore([]byte("secret"))
		r.Use(sessions.Sessions("mysession", store))
		{
			gob.Register(model.User{})
		}

		r.Use(middleware.Authz)

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
	r.POST("/login", ctrl.LoginHandle)
	r.POST("/logout", middleware.Logout)

	r.GET("/user/:name", ctrl.UserByName)

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
