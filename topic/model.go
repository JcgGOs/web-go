package topic

import (
	"time"

	"bloom.io/reply"
)

//Topic model
type Topic struct {
	ID       int64     `json:"id,omitempty"`
	UserID   int64     `json:"user_id,omitempty"`
	Title    string    `json:"title,omitempty"`
	Content  string    `json:"content,omitempty"`
	Status   int       `json:"status,omitempty"`
	ReplyNum int8      `json:"reply_num,omitempty"`
	CreateAt time.Time `json:"create_at,omitempty"`
	UpdateAt time.Time `json:"update_at,omitempty"`
}

//TO transfer object
type TO struct {
	Topic
	Username string
	Tags     []string
	Replies  []reply.TO
}
