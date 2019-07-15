package ctrl

import (
	"fmt"
	"net/http"

	"bloom.io/middleware"

	"bloom.io/service"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// SessionKey sessionkey
const SessionKey string = "_session_"

// Login login
func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}

//LoginHandle handle login
func LoginHandle(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	var redirect string
	if user, ok := service.Login(username, password); !ok {
		redirect = fmt.Sprintf("/login?redirect=%v", c.Request.RequestURI)
	} else {
		session := sessions.Default(c)
		{
			session.Set(middleware.SessionKey, user)
			session.Save()
		}

		redirect = c.DefaultQuery("redirect", "/")
	}

	c.Redirect(http.StatusMovedPermanently, redirect)
}
