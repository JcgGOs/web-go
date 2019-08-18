package login

import (
	"fmt"
	"net/http"
	"net/url"

	"bloom.io/common"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// SessionKey sessionkey
const SessionKey string = "_session_"

// RouteGET login
func RouteGET(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{"redirect": c.Request.RequestURI})
}

//RoutePOST handle login
func RoutePOST(c *gin.Context) {

	cmd := Command{
		Username: c.PostForm("username"),
		Password: c.PostForm("password"),
	}

	errors, ok := cmd.Validate()
	if ok {
		if user, succ := Handle(cmd); succ {
			session := sessions.Default(c)
			session.Set(SessionKey, user)
			session.Save()
		}
	}

	jump(c, errors)
}

//RouteLogout log out
func RouteLogout(c *gin.Context) {
	sessions.Default(c).Clear()
	jump(c, []common.Err{})
}

func jump(c *gin.Context, errors []common.Err) {
	redirect := c.DefaultQuery("redirect", "/")
	if len(errors) > 0 {
		redirect = fmt.Sprintf("/login?redirect=%v", redirect)
		for _, v := range errors {
			redirect = redirect + fmt.Sprintf("&%v=%v", v.Field, v.Message)
		}
	}
	uri, _ := url.QueryUnescape(redirect)
	c.Redirect(http.StatusMovedPermanently, uri)
}
