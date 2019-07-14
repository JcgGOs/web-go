package ctrl

import (
	"net/http"

	"github.com/gin-contrib/sessions"

	"bloom.io/service"

	"github.com/gin-gonic/gin"
)

// SessionKey sessionkey
const SessionKey string = "_session_"

// Login login
func Login(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(SessionKey)
	if user == nil {
		c.HTML(http.StatusOK, "login.html", gin.H{})
	}

}

//HandleLogin handle login
func HandleLogin(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	service.Login(username, password)
}
