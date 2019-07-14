package model

//User user
type User struct {
	Username string `form:"username" json:"username" xml:"username"`
	Password string `form:"password" json:"password"`
}
