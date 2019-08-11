package ctrl

import (
	"net/http"
	"strconv"

	"bloom.io/service"

	"github.com/gin-gonic/gin"
)

//Topic topic handle
func Topic(c *gin.Context) {
	id := c.Param("id")
	if len(id) == 0 {
		c.Redirect(http.StatusMovedPermanently, "/error/404")
		return
	}

	topicID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.Redirect(http.StatusMovedPermanently, "/error/404")
		c.Abort()
	}

	topic, _ := service.TopicByID(topicID)
	c.HTML(http.StatusOK, "topic.html", gin.H{
		"topic": topic,
	})
}
