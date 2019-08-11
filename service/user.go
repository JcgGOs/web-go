package service

import (
	"bloom.io/model"
)

//UserByID get user by Id
func UserByID(id int8) model.User {
	return model.User{Username: "username", Password: "password"}
}
