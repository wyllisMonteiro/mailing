package repositories

import (
	"fmt"
	"github.com/wyllisMonteiro/mailing/api/config"
	//"github.com/wyllisMonteiro/mailing/api/repositories"
)

type PostBody struct {
	Name   string
    Description string
    Mails []string
}

func AddBroadcastList(postBody PostBody) {
	db, err := config.ConnectToBDD()
	
	defer db.Close()

	if err != nil {
		return
	}

	user, err := GetOneSubscriber("kevin@gmail.com")
	if err != nil {
		panic(err.Error())
		return
	}

	fmt.Println(user.Name)

	/*// perform a db.Query insert
	insert, err := db.Query("UPDATE `user` SET `token` = ? WHERE `user`.`id` = ?", token, user_id)

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}

	// be careful deferring Queries if you are using transactions
	defer insert.Close()*/
}