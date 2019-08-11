package ctrl

import (
	"net/http"

	"bloom.io/model"

	"github.com/gin-gonic/gin"
)

//Topic topic handle
func Topic(c *gin.Context) {
	id := c.Param("id")
	if len(id) == 0 {
		c.Redirect(http.StatusMovedPermanently, "/error/404")
		return
	}

	topic := model.Topic{Title: "我的成神之路", Content: "just  a test"}
	c.HTML(http.StatusOK, "topic.html", gin.H{
		"topic": topic,
	})
}
