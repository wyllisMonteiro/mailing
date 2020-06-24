package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var USER_DB 		= "root"
var PASS_DB 		= "root"
var IP_DB 			= "127.0.0.1:3306"
var NAME_DB 		= "mailing"

func ConnectToBDD() (*sql.DB, error){
	db, err := sql.Open("mysql", USER_DB + ":" + PASS_DB + "@tcp(" + IP_DB + ")/" + NAME_DB)

    if err != nil {
		return nil, err
	} else {
		return db, nil
	}
}