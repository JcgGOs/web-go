package model

import "time"

//Topic model
type Topic struct {
	ID       int8      `json:"id,omitempty"`
	Title    string    `json:"title,omitempty"`
	Username string    `json:"username,omitempty"`
	CreateAt time.Time `json:"create_at,omitempty"`
}
