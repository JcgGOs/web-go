package service

import (
	"bloom.io/model"
)

//Login Action
func Login(username string, password string) (interface{}, bool) {
	if username == password {
		return model.User{Username: username, Password: password}, true
	} else {
		return nil, false
	}
}
