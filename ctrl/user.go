package ctrl

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//User user handler
func User(c *gin.Context) {
	name := c.Param("name")
	c.String(http.StatusOK, "Hello %s", name)
}

//GetUserByName user handler
func GetUserByName(c *gin.Context) {
	name := c.Param("name")
	c.String(http.StatusOK, "Hello %s", name)
}
