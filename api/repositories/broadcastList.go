package repositories

import (
	"github.com/wyllisMonteiro/mailing/api/config"
)

type PostBody struct {
	Name   string
    Description string
    Mails []string
}

func AddBroadcastList(test PostBody) {
	db, err := config.ConnectToBDD()
	
	defer db.Close()

	if err != nil {
		//return user, err
	}

	/*// perform a db.Query insert
	insert, err := db.Query("UPDATE `user` SET `token` = ? WHERE `user`.`id` = ?", token, user_id)

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}

	// be careful deferring Queries if you are using transactions
	defer insert.Close()*/
}