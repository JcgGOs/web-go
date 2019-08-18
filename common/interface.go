package common

import "database/sql"

//Validator interface
type Validator interface {

	//Validate the object
	Validate() []Err
}

//Mapper database Mapper
type Mapper func(*sql.Rows) []interface{}
