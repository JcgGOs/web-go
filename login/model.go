package login

import (
	"bloom.io/common"
)

// Command User login request cmd
type Command struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

//Validate the command
func (command *Command) Validate() (errors []common.Err, ok bool) {

	//Username
	errors = append(errors, common.Validate(
		"username",
		command.Username, "required")...,
	)

	//Password
	errors = append(errors, common.Validate(
		"password",
		command.Password, "required")...,
	)

	return errors, len(errors) < 1
}
