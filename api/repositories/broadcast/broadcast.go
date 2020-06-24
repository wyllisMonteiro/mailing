package broadcast

import (
	"fmt"
	"github.com/wyllisMonteiro/mailing/api/config"
	sub "github.com/wyllisMonteiro/mailing/api/repositories/subscriber"
)

type CreateBroadcastRequest struct {
	Name   string
    Description string
    Mails []string
}

type AddSubRequest struct {
	BroadcastName string
	SubscriberMail string
}

func CreateBroadcast(createBroadcastRequest CreateBroadcastRequest) {
	db, err := config.ConnectToBDD()
	
	defer db.Close()

	res, err := db.Exec("INSERT `broadcast`(`name`, `description`) VALUES (?, ?)", createBroadcastRequest.Name, createBroadcastRequest.Description)
	if err != nil {
		return
	}

	broadcast_id, err := res.LastInsertId()
	if err != nil {
		return
	}

	for mailIndex := 0; mailIndex < len(createBroadcastRequest.Mails); mailIndex++ {
		subscriber, err := sub.FindBy("mail", createBroadcastRequest.Mails[mailIndex])
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

func AddSubscriber(addSubRequest AddSubRequest) {
	subscriber, err := sub.FindBy("mail", addSubRequest.SubscriberMail)
	if err != nil {
		panic(err.Error())
		return
	}

	fmt.Println(subscriber.ID)
}