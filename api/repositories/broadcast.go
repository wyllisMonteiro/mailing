package repositories

import (
	"strconv"
	"fmt"
	"net/http"
	"github.com/wyllisMonteiro/mailing/api/config"
	"github.com/wyllisMonteiro/mailing/api/service"
)

type Broadcasts struct {
	ID string `json:"id"`
	Name   string `json:"name"`
	Description string `json:"description"`
	Mails []string `json:"mails"`
}

type Broadcast struct {
	ID string `json:"id"`
	Name   string `json:"name"`
	Description string `json:"description"`
	Mail string `json:"mail"`
}

func BroadcastFindBy(key string, val string) (Broadcast, error) {
	var broadResponse Broadcast

	db, err := config.ConnectToBDD()
	
	defer db.Close()

	if err != nil {
		return broadResponse, err
	}

	err = db.QueryRow("SELECT id, name, description FROM broadcast WHERE " + key + " = ?", val).Scan(&broadResponse.ID, &broadResponse.Name, &broadResponse.Description)
	
	if err != nil {
		fmt.Println(val)
		return broadResponse, err
	}

	return broadResponse, nil
}

func BroadcastFindWithSubs(name string) (Broadcasts, error) {
	var broadsResponse Broadcasts
	var broadResponse Broadcast

	db, err := config.ConnectToBDD()
	
	defer db.Close()

	if err != nil {
		return broadsResponse, err
	}	
	
	selectFields 	:= "broadcast.id, broadcast.name, broadcast.description, subscriber.mail"
	fromTable 		:= "broadcast, broadcast_subscriber, subscriber"
	where 			:= "broadcast.name = ? AND broadcast.id = broadcast_subscriber.broadcast_id AND subscriber.id = broadcast_subscriber.subscriber_id"
	
	results, err := db.Query("SELECT " + selectFields + " FROM " + fromTable + " WHERE " + where, name)
	if err != nil {
		fmt.Println(err.Error())
		return broadsResponse, err
	}

	for results.Next() {
        err = results.Scan(&broadResponse.ID, &broadResponse.Name, &broadResponse.Description, &broadResponse.Mail)
        if err != nil {
            fmt.Println(err.Error())
		}
		
		broadsResponse.Mails = append(broadsResponse.Mails, broadResponse.Mail)
	}
	
	broadsResponse.ID = broadResponse.ID
	broadsResponse.Name = broadResponse.Name
	broadsResponse.Description = broadResponse.Description

	return broadsResponse, nil
}

func CreateBroadcast(createBroadcastRequest Broadcasts) (Broadcasts, error) {
	db, err := config.ConnectToBDD()
	
	defer db.Close()

	res, err := db.Exec("INSERT `broadcast`(`name`, `description`) VALUES (?, ?)", createBroadcastRequest.Name, createBroadcastRequest.Description)
	if err != nil {
		return createBroadcastRequest, err
	}

	broadcast_id, err := res.LastInsertId()
	if err != nil {
		return createBroadcastRequest, err
	}

	createBroadcastRequest.ID = strconv.Itoa(int(broadcast_id))

	for mailIndex := 0; mailIndex < len(createBroadcastRequest.Mails); mailIndex++ {
		subscriber, err := SubscriberFindBy("mail", createBroadcastRequest.Mails[mailIndex])
		if err != nil {
			return createBroadcastRequest, err
		}

		insert, err := db.Query("INSERT INTO `broadcast_subscriber` (`broadcast_id`, `subscriber_id`) VALUES (?, ?)", broadcast_id, subscriber.ID)

		if err != nil {
			return createBroadcastRequest, err
		}

		defer insert.Close()
	}

	return createBroadcastRequest, nil
}

type SubRequest struct {
	BroadcastName string `json:"broadcast_name"`
	SubscriberMail string `json:"subscriber_mail"`
}

func BroadcastAddSubscriber(subRequest SubRequest) (SubRequest, error) {
	subscriber, err := SubscriberFindBy("mail", subRequest.SubscriberMail)
	if err != nil {
		return subRequest, err
	}

	broad, err := BroadcastFindBy("name", subRequest.BroadcastName)
	if err != nil {
		return subRequest, err
	}

	db, err := config.ConnectToBDD()

	insert, err := db.Query("INSERT INTO `broadcast_subscriber` (`broadcast_id`, `subscriber_id`) VALUES (?, ?)", broad.ID, subscriber.ID)

	if err != nil {
		return subRequest, err
	}

	defer insert.Close()

	return subRequest, nil
}

func BroadcastDeleteSubscriber(subRequest SubRequest) (SubRequest, error) {
	subscriber, err := SubscriberFindBy("mail", subRequest.SubscriberMail)
	if err != nil {
		return subRequest, err
	}

	broad, err := BroadcastFindBy("name", subRequest.BroadcastName)
	if err != nil {
		return subRequest, err
	}

	db, err := config.ConnectToBDD()

	delete, err := db.Query("DELETE FROM `broadcast_subscriber` WHERE broadcast_id = ? AND subscriber_id = ?", broad.ID, subscriber.ID)

	if err != nil {
		return subRequest, err
	}

	defer delete.Close()

	return subRequest, nil
}