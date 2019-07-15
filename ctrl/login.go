package ctrl

import (
	"fmt"
	"net/http"

	"bloom.io/service"

	"github.com/gin-gonic/gin"
)

// SessionKey sessionkey
const SessionKey string = "_session_"

// Login login
func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}

//HandleLogin handle login
func HandleLogin(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	redirect := c.DefaultQuery("redirect", "/")

	if _, ok := service.Login(username, password); !ok {
		redirect = fmt.Sprintf("/login?redirect=%v", c.Request.RequestURI)
	}

	c.Redirect(http.StatusMovedPermanently, redirect)
}

//SessionFilter check session
func SessionFilter(c *gin.Context) {
	user, ok := c.Get(SessionKey)
	if ok && user != nil {
		c.Next()
		return
	}

	if user == nil {
		c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("/login?redirect=%v", c.Request.RequestURI))
	}
}
