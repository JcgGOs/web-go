package login

import "bloom.io/user"

//Login Action
func Login(username string, password string) (interface{}, bool) {
	if username == password {
		return user.User{Username: username, Password: password}, true
	}

	return nil, false
}
