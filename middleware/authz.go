package middleware

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-contrib/sessions"

	"github.com/gin-gonic/gin"
)

var excludePath = []string{
	"/",
	"/login",
}

//Authz authz
func Authz(c *gin.Context) {
	uri := strings.Split(c.Request.RequestURI, "?")[0]

	for index := range excludePath {
		if excludePath[index] == uri {
			return
		}
	}

	user := sessions.Default(c).Get("user")
	if user == nil {
		redirect, _ := url.QueryUnescape(c.Request.RequestURI)
		redirect = fmt.Sprintf("/login?redirect=%v", redirect)
		c.Redirect(http.StatusMovedPermanently, redirect)
		c.Abort()
	}
}

//Logout log out
func Logout(c *gin.Context) {
	sessions.Default(c).Clear()

	redirect := c.DefaultQuery("redirect", "/")
	c.Redirect(http.StatusMovedPermanently, redirect)
}
