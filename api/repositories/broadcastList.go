package repositories

import (
	"fmt"
	"github.com/wyllisMonteiro/mailing/api/config"
	"database/sql"
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

	broadcast_id, err := CreateBroadcast(db, postBody)

	if err != nil {
		return
	}

	for mailIndex := 0; mailIndex < len(postBody.Mails); mailIndex++ {
		subscriber, err := GetOneSubscriber(postBody.Mails[mailIndex])
		if err != nil {
			panic(err.Error())
			return
		}

		insert, err := db.Query("INSERT INTO `broadcast_subscriber` (`broadcast_id`, `subscriber_id`) VALUES (?, ?)", broadcast_id, subscriber.ID)

		if err != nil {
			panic(err.Error())
		}

		defer insert.Close()
	}
}

func CreateBroadcast(db *sql.DB, postBody PostBody) (int64, error) {
	res, err := db.Exec("INSERT `broadcast`(`name`, `description`) VALUES (?, ?)", postBody.Name, postBody.Description)
	if err != nil {
		return 0, err
	}

	broadcast_id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return broadcast_id, nil
}