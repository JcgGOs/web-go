package service

import (
	"time"

	"bloom.io/model"
)

//GetTopics service topics
func GetTopics(num int) []model.Topic {
	return []model.Topic{
		model.Topic{ID: 1, Title: "title", Username: "username", CreateAt: time.Now()},
		model.Topic{ID: 2, Title: "title_2", Username: "username_2", CreateAt: time.Now()},
	}
}
