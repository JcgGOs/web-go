package model

//User user
type User struct {
	ID       int8   `json:"id,omitempty"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email,omitempty"`
}
