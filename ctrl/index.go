package ctrl

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Index ctrl.Index
func Index(c *gin.Context) {
	c.HTML(http.StatusOK,"index.html",gin.H{"msg":"ok"})
}
