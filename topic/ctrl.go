package topic

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//GetByID topic handle
func GetByID(c *gin.Context) {
	id := c.Param("id")
	if len(id) == 0 {
		c.Redirect(http.StatusMovedPermanently, "/error/404")
		return
	}

	topicID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.Redirect(http.StatusMovedPermanently, "/error/404")
		c.Abort()
		return
	}

	topic, _ := FindByID(topicID)
	c.HTML(http.StatusOK, "topic.html", gin.H{
		"topic": topic,
	})
}
