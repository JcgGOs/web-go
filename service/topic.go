package service

import (
	"strconv"
	"time"

	"bloom.io/model"
)

//Topics service topics
func Topics(num int) []model.Topic {
	return []model.Topic{
		model.Topic{ID: 1, Title: "title", Content: "username", CreateAt: time.Now()},
		model.Topic{ID: 2, Title: "title_2", Content: "username_2", CreateAt: time.Now()},
	}
}

//TopicByID service topics
func TopicByID(num int64) (topic model.TopicTO, err error) {
	return model.TopicTO{
		Topic: model.Topic{
			ID:      int64(num),
			Content: "content" + strconv.FormatInt(int64(num), 10),
		},
		Username: "小明",
		Tags:     []string{"java", "golang", "javascript"},
		Replies: []model.ReplyTO{
			model.ReplyTO{
				Reply: model.Reply{
					Content: "Reply Content 1",
				},
				Username: "Reply1",
			},
			model.ReplyTO{
				Reply: model.Reply{
					Content: "Reply Content 2",
				},
				Username: "Reply2",
			},
			model.ReplyTO{
				Reply: model.Reply{
					Content: "Reply Content 3",
				},
				Username: "Reply3",
			},
			model.ReplyTO{
				Reply: model.Reply{
					Content: "Reply Content 4",
				},
				Username: "Reply4",
			},
			model.ReplyTO{
				Reply: model.Reply{
					Content: "Reply Content 4",
				},
				Username: "Reply4",
			},
		},
	}, nil
}
