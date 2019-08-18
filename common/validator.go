package common

import (
	"fmt"
	"regexp"
)

//Err Error
type Err struct {
	Field   string
	Message string
}

//Validate validate the field
func Validate(field string, value string, rules ...string) []Err {
	errors := []Err{}

	for idx := range rules {
		switch rules[idx] {
		case "required":
			if len(value) < 1 {
				errors = append(
					errors,
					Err{
						Field:   field,
						Message: "field-required",
					},
				)
			}
		case "email":
			expression := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` // to match the email address
			pattern := regexp.MustCompile(expression)
			if !pattern.MatchString(value) {
				errors = append(
					errors,
					Err{
						Field:   field,
						Message: "invalid-email",
					},
				)
			}
		default:
			fmt.Println("Can not found rule: %V", rules[idx])
		}
	}
	return errors
}
