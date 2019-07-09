package model

type User struct {
	Account  string `form:"name"`
	Password string `form:"password"`
}
