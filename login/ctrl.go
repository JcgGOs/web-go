package login

import (
	"net/http"
	"net/url"

	"bloom.io/user"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// SessionKey sessionkey
const SessionKey string = "_session_"

// GET login
func GET(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{"loginUrl": c.Request.RequestURI})
}

//POST handle login
func POST(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	user := user.User{
		Username: username,
		Password: password,
	}
	session := sessions.Default(c)
	{
		session.Set("user", user)
		session.Save()
	}

	redirect := c.DefaultQuery("redirect", "/")
	uri, _ := url.QueryUnescape(redirect)
	c.Redirect(http.StatusMovedPermanently, uri)
}

//Logout log out
func Logout(c *gin.Context) {
	sessions.Default(c).Clear()

	redirect := c.DefaultQuery("redirect", "/")
	c.Redirect(http.StatusMovedPermanently, redirect)
}
