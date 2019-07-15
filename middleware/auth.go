package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

const SessionKey string = "_session_"

//Authz authz
func Authz(c *gin.Context) {

	fmt.Println("JUST TEST")
	if "GET" == c.Request.Method && "/login" == c.Request.RequestURI {
		c.Next()
	} else {
		if user := sessions.Default(c).Get(SessionKey); user != nil {
			c.Next()
		} else {
			c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("/login?redirect=%v", c.Request.RequestURI))
		}
	}
}
