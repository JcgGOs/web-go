package ctrl

import (
	"net/http"

	"github.com/gin-contrib/sessions"

	"github.com/gin-gonic/gin"
)

//Index ctrl.Index
func Index(c *gin.Context) {
	user := sessions.Default(c).Get("user")
	c.HTML(http.StatusOK, "index.html", gin.H{"user": user})
}
