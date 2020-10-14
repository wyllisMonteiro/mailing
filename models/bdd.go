package models

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	// usefull to import mysql
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// DB : database
var DB *sql.DB

// GoDotEnvVariable : Get environment variable in .env file
func GoDotEnvVariable(key string) (string, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return "", err
	}

	return os.Getenv(key), nil
}

// GetVarsDB : Get environment variable for setting up database in .env file
func GetVarsDB() (string, string, string, string, error) {
	userDB, err := GoDotEnvVariable("USERDB")
	if err != nil {

		return "", "", "", "", err
	}

	passDB, err := GoDotEnvVariable("PASSDB")
	if err != nil {
		return "", "", "", "", err
	}

	ipDB, err := GoDotEnvVariable("IPDB")
	if err != nil {
		return "", "", "", "", err
	}

	nameDB, err := GoDotEnvVariable("NAMEDB")
	if err != nil {
		return "", "", "", "", err
	}

	return userDB, passDB, ipDB, nameDB, nil
}

// ConnectToBDD : Make connexion with database
func ConnectToBDD() {
	userDB, passDB, ipDB, nameDB, err := GetVarsDB()
	if err != nil {
		fmt.Println(userDB)
		fmt.Println(passDB)
		fmt.Println(ipDB)
		fmt.Println(nameDB)
		return
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", userDB, passDB, ipDB, nameDB)
	DB, err = sql.Open("mysql", dsn)

	if err != nil {
		log.Println(err)
	}
}
