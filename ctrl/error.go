package ctrl

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Error404 ctrl.Index
func Error404(c *gin.Context) {
	c.HTML(http.StatusOK, "404.html", gin.H{})
}

//Error500 ctrl.Index
func Error500(c *gin.Context) {
	c.HTML(http.StatusOK, "500.html", gin.H{})
}
