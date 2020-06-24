package repositories

import (
	"github.com/wyllisMonteiro/mailing/api/config"
)

type User struct {
    ID   int    `json:"id"`
    Login string `json:"login"`
    Password string `json:"password"`
    Token string `json:"token"`
}

var user User

func GetOneUser(login string) (User, error) {
	db, err := config.ConnectToBDD()
	
	defer db.Close()

	if err != nil {
		return user, err
	}

	err = db.QueryRow("SELECT id, login, password FROM user WHERE login = ?", login).Scan(&user.ID, &user.Login, &user.Password)
	
	if err != nil {
		return user, err
	}

	return user, nil
}

func InsertUserToken(token string, user_id int) {
	db, err := config.ConnectToBDD()

	defer db.Close()

	// perform a db.Query insert
	insert, err := db.Query("UPDATE `user` SET `token` = ? WHERE `user`.`id` = ?", token, user_id)

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}

	// be careful deferring Queries if you are using transactions
	defer insert.Close()
}
