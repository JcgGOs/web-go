package reply

import (
	"strconv"
	"time"
)

//Reply reply to topic
type Reply struct {
	ID       int64     `json:"id,omitempty"`
	TopicID  int64     `json:"topic_id,omitempty"`
	ReplyID  int64     `json:"reply_id,omitempty"`
	UserID   int64     `json:"user_id,omitempty"`
	Content  string    `json:"content,omitempty"`
	CreateAt time.Time `json:"create_at,omitempty"`
	UpdateAt time.Time `json:"update_at,omitempty"`
}

//TO reply to topic
type TO struct {
	Reply
	Username string
}

func (obj *Reply) key() string {
	return "reply_" + strconv.FormatInt(obj.ID, 10)
}
