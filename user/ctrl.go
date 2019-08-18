package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//GetByName user handler
func GetByName(c *gin.Context) {
	name := c.Param("name")
	c.String(http.StatusOK, "Hello %s", name)
}

//PostByName user handler
func PostByName(c *gin.Context) {
	name := c.Param("name")
	c.String(http.StatusOK, "Hello %s", name)
}
