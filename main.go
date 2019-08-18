package main

import (
	"encoding/gob"
	"html/template"
	"net/http"

	"bloom.io/common"
	"bloom.io/login"
	"bloom.io/topic"
	"bloom.io/user"

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
			gob.Register(user.User{})
		}

		// r.Use(middleware.Authz)

		r.StaticFile("/robots.txt", "./static/robots.txt")
		r.StaticFile("/favicon.ico", "./static/favicon.ico")
		r.Static("/static", "./static")
		r.SetFuncMap(template.FuncMap{
			"fmtDate": common.FmtDate,
			"fmtTime": common.FmtTime,
		})
		r.LoadHTMLGlob("views/*.html")
		r.NoRoute(func(c *gin.Context) {
			c.HTML(http.StatusNotFound, "404.html", gin.H{})
		})
	}

	r.GET("/", common.Index)
	r.GET("/error/:error", common.Error)

	r.GET("/login", login.GET)
	r.POST("/login", login.POST)
	r.GET("/logout", login.Logout)

	//user/:name
	userGroup := r.Group("/user")
	{
		userGroup.GET("/:name", user.GetByName)
	}

	//topic/:id
	topicGroup := r.Group("/topic")
	{
		topicGroup.GET("/:id", topic.GetByID)
	}

	r.Run(":8000")
}
