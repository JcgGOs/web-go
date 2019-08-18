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

//RouteByID user handler
func RouteByID(c *gin.Context) {
	id := c.Param("id")
	c.String(http.StatusOK, "Hello %s", id)
}

//PostByName user handler
func PostByName(c *gin.Context) {
	name := c.Param("name")
	c.String(http.StatusOK, "Hello %s", name)
}
