package model

import (
	"strconv"
	"time"
)

//OAuth information
//Type: Wechat/Alipay/Github
type OAuth struct {
	ID       int64
	UserID   int64
	Type     string
	Token    string
	Extra    string
	ExpireAt time.Time
	CreateAt time.Time
	UpdateAt time.Time
}

func (obj *OAuth) key() string {
	return "oauth_" + strconv.FormatInt(obj.ID, 10)
}
