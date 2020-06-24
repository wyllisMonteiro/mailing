package broadcast

import (
	"github.com/wyllisMonteiro/mailing/api/config"
	sub "github.com/wyllisMonteiro/mailing/api/repositories/subscriber"
)

type CreateBroadcastRequest struct {
	Name   string
    Description string
    Mails []string
}

type SubRequest struct {
	BroadcastName string
	SubscriberMail string
}

type BroadcastResponse struct {
	ID   int    `json:"id"`
    Name string `json:"name"`
    Description string `json:"description"`
}

func findBy(key string, val string) (BroadcastResponse, error) {
	var broadResponse BroadcastResponse

	db, err := config.ConnectToBDD()
	
	defer db.Close()

	if err != nil {
		return broadResponse, err
	}	

	err = db.QueryRow("SELECT id FROM broadcast WHERE " + key + " = ?", val).Scan(&broadResponse.ID)
	
	if err != nil {
		return broadResponse, err
	}

	return broadResponse, nil
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

func AddSubscriber(subRequest SubRequest) {
	subscriber, err := sub.FindBy("mail", subRequest.SubscriberMail)
	if err != nil {
		panic(err.Error())
		return
	}

	broad, err := findBy("name", subRequest.BroadcastName)
	if err != nil {
		panic(err.Error())
		return
	}

	db, err := config.ConnectToBDD()

	insert, err := db.Query("INSERT INTO `broadcast_subscriber` (`broadcast_id`, `subscriber_id`) VALUES (?, ?)", broad.ID, subscriber.ID)

	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()
}

func DeleteSubscriber(subRequest SubRequest) {
	subscriber, err := sub.FindBy("mail", subRequest.SubscriberMail)
	if err != nil {
		panic(err.Error())
		return
	}

	broad, err := findBy("name", subRequest.BroadcastName)
	if err != nil {
		panic(err.Error())
		return
	}

	db, err := config.ConnectToBDD()

	delete, err := db.Query("DELETE FROM `broadcast_subscriber` WHERE broadcast_id = ? AND subscriber_id = ?", broad.ID, subscriber.ID)

	if err != nil {
		panic(err.Error())
	}

	defer delete.Close()
}