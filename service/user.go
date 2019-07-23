package service

import (
	"bloom.io/model"
)

//FindUserByID get user by Id
func FindUserByID(id int8) model.User {
	return model.User{Username: "username", Password: "password"}
}
