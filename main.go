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

	r.GET("/", common.RouteIndex)
	r.GET("/error/:error", common.RouteError)

	r.GET("/login", login.RouteGET)
	r.POST("/login", login.RoutePOST)
	r.GET("/logout", login.RouteLogout)

	//user/:name
	userGroup := r.Group("/user")
	{
		userGroup.GET("/:id", user.RouteByID)
	}

	//topic/:id
	topicGroup := r.Group("/topic")
	{
		topicGroup.GET("/:id", topic.RouteByID)
	}
	r.Run(":8000")
}
