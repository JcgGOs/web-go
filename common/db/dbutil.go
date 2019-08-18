package db

import (
	"database/sql"
	"log"

	"bloom.io/common"

	//import mysql
	_ "github.com/go-sql-driver/mysql"
)

//Query quer sql
func Query(query string, mapper common.Mapper) (bool, []interface{}) {
	conn, err := sql.Open("mysql", "root:mysql@tcp(127.0.0.1:3306)/test?parseTime=true")
	defer conn.Close()

	if err != nil {
		log.Fatalln(err)
		return false, nil
	}

	rows, err := conn.Query(query)
	defer rows.Close()

	if err != nil {
		log.Fatalln(err)
		return false, nil
	}
	return true, mapper(rows)
}
