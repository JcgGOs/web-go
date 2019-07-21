package ctrl

import (
	"net/http"
	"net/url"

	"bloom.io/model"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// SessionKey sessionkey
const SessionKey string = "_session_"

// Login login
func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{"loginUrl": c.Request.RequestURI})
}

//LoginHandle handle login
func LoginHandle(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	user := model.User{
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
