package topic

import (
	"strconv"
	"time"

	"bloom.io/reply"
)

//Topics service topics
func Topics(num int) []Topic {
	return []Topic{
		Topic{ID: 1, Title: "title", Content: "username", CreateAt: time.Now()},
		Topic{ID: 2, Title: "title_2", Content: "username_2", CreateAt: time.Now()},
	}
}

//FindByID service topics
func FindByID(num int64) (topic TO, err error) {
	return TO{
		Topic: Topic{
			ID:      int64(num),
			Content: "content" + strconv.FormatInt(int64(num), 10),
		},
		Username: "小明",
		Tags:     []string{"java", "golang", "javascript"},
		Replies: []reply.TO{
			reply.TO{
				Reply: reply.Reply{
					Content: "Reply Content 1",
				},
				Username: "Reply1",
			},
			reply.TO{
				Reply: reply.Reply{
					Content: "Reply Content 2",
				},
				Username: "Reply2",
			},
			reply.TO{
				Reply: reply.Reply{
					Content: "Reply Content 3",
				},
				Username: "Reply3",
			},
			reply.TO{
				Reply: reply.Reply{
					Content: "Reply Content 4",
				},
				Username: "Reply4",
			},
			reply.TO{
				Reply: reply.Reply{
					Content: "Reply Content 4",
				},
				Username: "Reply4",
			},
		},
	}, nil
}
