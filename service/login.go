package service

//Login Action
func Login(username string, password string) (string, bool) {
	if username == password {
		return "login successfully", true
	} else {
		return "username password not match", false
	}
}
