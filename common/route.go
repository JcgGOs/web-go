package common

import (
	"net/http"
	"strconv"
	"time"

	"bloom.io/topic"

	"github.com/gin-contrib/sessions"

	"github.com/gin-gonic/gin"
)

//RouteIndex ctrl.RouteIndex
func RouteIndex(c *gin.Context) {
	user := sessions.Default(c).Get("user")
	// if user == nil {
	// 	user = model.User{Username: "anonymous"}
	// }

	c.HTML(http.StatusOK, "index.html", gin.H{
		"user":   user,
		"topics": mock(50),
	})
}

func mock(num int) []topic.TO {
	topics := make([]topic.TO, num)
	for i := 0; i < num; i++ {
		topics[i] = topic.TO{
			Topic: topic.Topic{
				ID:       int64(i),
				UserID:   int64(i),
				Title:    "我的Go语言学习之路_" + strconv.FormatInt(int64(i), 10),
				Status:   1,
				CreateAt: time.Now(),
			},
			Username: "username_" + strconv.FormatInt(int64(i), 10),
		}
	}
	return topics
}

//RouteError ctrl.Index
func RouteError(c *gin.Context) {
	code := c.Param("error")
	switch code {
	case "404":
		c.HTML(http.StatusNotFound, "404.html", gin.H{})
	default:
		c.HTML(http.StatusInternalServerError, "500.html", gin.H{})
	}
}
