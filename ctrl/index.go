package ctrl

import (
	"net/http"
	"strconv"
	"time"

	"bloom.io/model"

	"github.com/gin-contrib/sessions"

	"github.com/gin-gonic/gin"
)

//Index ctrl.Index
func Index(c *gin.Context) {
	user := sessions.Default(c).Get("user")
	// if user == nil {
	// 	user = model.User{Username: "anonymous"}
	// }

	c.HTML(http.StatusOK, "index.html", gin.H{
		"user":   user,
		"topics": mock(50),
	})
}

func mock(num int) []model.TopicTO {
	topics := make([]model.TopicTO, num)
	for i := 0; i < num; i++ {
		topics[i] = model.TopicTO{Topic: model.Topic{ID: int64(i), UserID: int64(i), Title: "我的Go语言学习之路_" + strconv.FormatInt(int64(i), 10), Status: 1, CreateAt: time.Now()}, Username: "username_" + strconv.FormatInt(int64(i), 10)}
	}
	return topics
}
