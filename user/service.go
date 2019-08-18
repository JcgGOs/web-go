package user

//FindByID get user by Id
func FindByID(id int8) User {
	return User{ID: int64(id), Username: "username", Password: "password"}
}
