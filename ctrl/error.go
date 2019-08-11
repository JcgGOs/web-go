package ctrl

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Error ctrl.Index
func Error(c *gin.Context) {
	code := c.Param("error")
	switch code {
	case "404":
		c.HTML(http.StatusNotFound, "404.html", gin.H{})
	default:
		c.HTML(http.StatusInternalServerError, "500.html", gin.H{})
	}
}
